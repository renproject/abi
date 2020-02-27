package abi

import (
	"encoding/binary"
	"fmt"
	"io"
	"sort"

	"github.com/renproject/surge"
)

type Record struct {
	inner []field
}

func (record Record) Marshal(w io.Writer) error {
	// Write length prefix
	b := [4]byte{}
	binary.LittleEndian.PutUint32(b[:], uint32(len(record.inner)))
	if _, err := w.Write(b[:]); err != nil {
		return err
	}

	for _, field := range record.inner {
		// Write length prefix for the field name
		name := []byte(field.name)
		b := [2]byte{}
		binary.LittleEndian.PutUint16(b[:], uint16(len(name)))
		if _, err := w.Write(b[:]); err != nil {
			return err
		}
		// Write field name
		if _, err := w.Write(name); err != nil {
			return err
		}
		// Write field value using top-level marshaling function so that type
		// information is written
		if err := Marshal(field.value, w); err != nil {
			return err
		}
	}
	return nil
}

func (record *Record) Unmarshal(r io.Reader) error {
	_, err := record.unmarshalAndReturnSize(r, MaxSize)
	return err
}

func (record *Record) unmarshalAndReturnSize(r io.Reader, maxSize uint32) (uint32, error) {
	len := uint32(0)
	if err := surge.Unmarshal(&len, r); err != nil {
		return 0, err
	}

	// Restrict the size of the abstract data type to prevent forced OOM
	// allocations.
	if len > maxSize {
		return 0, fmt.Errorf("expected size<=%v, got size>=%v", maxSize, len)
	}
	record.inner = make([]field, 0, len)
	maxSize -= len
	if maxSize < 0 {
		return len, fmt.Errorf("exceeded maximum size")
	}

	sizeInBytes := len
	for i := uint32(0); i < len; i++ {
		// Unmarshal the field name.
		nameLen := uint16(0)
		if err := surge.Unmarshal(&nameLen, r); err != nil {
			return sizeInBytes, err
		}
		if uint32(nameLen) > maxSize {
			return sizeInBytes, fmt.Errorf("expected len<=%v, got len=%v", maxSize, nameLen)
		}
		nameData := make([]byte, nameLen)
		if _, err := io.ReadFull(r, nameData); err != nil {
			return sizeInBytes, err
		}
		name := string(nameData)
		sizeInBytes += uint32(nameLen) // The offset of -1 is counted when unmarshaling the value.
		maxSize -= uint32(nameLen)
		if maxSize < 0 {
			return sizeInBytes, fmt.Errorf("exceeded maximum size")
		}

		// Unmarshal field value using the top-level unmarshaling function so
		// that type information is read.
		value, valueSize, err := unmarshalAndReturnSize(r, maxSize)
		if err != nil {
			return sizeInBytes, err
		}
		sizeInBytes += (valueSize - 1) // Offset by -1, because we started with sizeInBytes := len
		maxSize -= (valueSize - 1)     // Offset by -1, because we have already done maxSize -= len
		if maxSize < 0 {
			return sizeInBytes, fmt.Errorf("exceeded maximum size")
		}
		record.inner = append(record.inner, field{name: name, value: value})
	}
	return sizeInBytes, nil
}

func (record Record) SizeHint() int {
	size := 4 // Size of length prefix
	for _, field := range record.inner {
		size += 2                       // Size of field name length prefix (note: we use uint16 for the length of field names)
		size += len([]byte(field.name)) // Size of field name
		size += 2                       // Size of field value type prefix
		size += field.value.SizeHint()  // Size of field value
	}
	return size
}

func (record Record) Type() Type {
	return TypeRecord
}

func (record *Record) Set(k string, v Value) {
	i := sort.Search(len(record.inner), func(i int) bool {
		return record.inner[i].name >= k
	})
	if i < 0 {
		record.inner = append(record.inner, field{name: k, value: v})
		return
	}
	if record.inner[i].name == k {
		record.inner[i].value = v
		return
	}
	record.inner = append(record.inner[:i], append([]field{field{name: k, value: v}}, record.inner[i:]...)...)
}

func (record *Record) Remove(k string) {
	i := sort.Search(len(record.inner), func(i int) bool {
		return record.inner[i].name >= k
	})
	if i < 0 {
		return
	}
	record.inner = append(record.inner[:i], record.inner[i+1:]...)
}

func (record Record) Get(k string) (Value, bool) {
	i := sort.Search(len(record.inner), func(i int) bool {
		return record.inner[i].name >= k
	})
	if i < 0 {
		return nil, false
	}
	return record.inner[i].value, true
}

func (record Record) Len() int {
	return len(record.inner)
}

func (record Record) ForEach(f func(k string, v Value)) {
	for _, field := range record.inner {
		f(field.name, field.value)
	}
}

type field struct {
	name  string
	value Value
}
