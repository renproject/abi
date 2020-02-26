package abi

import (
	"encoding/binary"
	"fmt"
	"io"
	"sort"

	"github.com/renproject/surge"
)

const (
	maxRecordFieldNameLength = 64
)

type field struct {
	name  string
	value Value
}

type Record struct {
	inner []field
}

func (record Record) Marshal(w io.Writer) error {
	b := [4]byte{}
	binary.LittleEndian.PutUint32(b[:], uint32(len(record.inner)))
	if err := w.Write(b[:]); err != nil {
		return err
	}
	for _, field := range record.inner {
		// Marshal the field name
		name := []byte(field.name)
		nameLen := uint32(len(name))
		if err := surge.Marshal(uint32(name), w); err != nil {
			return err
		}
		if _, err := w.Write(name); err != nil {
			return err
		}
		// Marshal the field value
		if err := Marshal(field.value, w); err != nil {
			return err
		}
	}
	return nil
}

func (record *Record) Unmarshal(r io.Reader) error {
	_, err := record.unmarshalAbstractDataType(r, MaxAbstractDataTypeSize)
	return err
}

func (record *Record) unmarshalAbstractDataType(r io.Reader, abstractDataTypeSizeLimit int) (int, error) {
	size := uint32(0)
	if err := Unmarshal(&size, r); err != nil {
		return size, err
	}

	// Restrict the size of the abstract data type to prevent forced OOM
	// allocations.
	if size > abstractDataTypeSizeLimit {
		return size, fmt.Errorf("expected size<=%v, got size=%v", MaxRecordLen, size)
	}
	abstractDataTypeSizeLimit -= size

	record.inner := make([]field, size)
	for i := 0; i < len; i++ {
		// Unmarshal field name
		nameLen := uint32(0)
		if err := surge.Unmarshal(&nameLen, r); err != nil {
			return err
		}
		if nameLen > maxRecordFieldNameLength {
			return fmt.Errorf("expected len<=%v, got len=%v", maxRecordFieldNameLength, nameLen)
		}
		nameData := make([]byte, nameLen)
		if _, err := io.ReadFull(r, nameData); err != nil {
			return err
		}
		name := string(nameData)
		size += nameLen
		abstractDataTypeSizeLimit -= nameLen
		if abstractDataTypeSizeLimit < 0 {
			return size, fmt.Errorf("exceeded maximum abstract data size")
		}

		// Unmarshal field value
		value, abstractDataTypeSize, err := unmarshal(r, abstractDataTypeSizeLimit)
		if err != nil {
			return size, err
		}

		// Handle the size of abstract data types within abstract data types.
		abstractDataTypeSizeLimit -= abstractDataTypeSize
		if abstractDataTypeSizeLimit < 0 {
			return size, fmt.Errorf("exceeded maximum abstract data size")
		}
		record.inner[i] = field{name: name, value: value}
	}

	return size, nil
}

func (ls Record) SizeHint() int {
	size := 4 // Size of length prefix
	for _, elem := range ls {
		size += elem.SizeHint()
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
		record.inner[i].name = v
		return
	}
	record.inner = append(record.inner[:i], append([]field{field{name: k, value: v}}, record.inner[i:]...)...)
}

func (record *Record) Remove(i int) {
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

func (record Record) Size() int {
	return len(record.inner)
}

func (record Record) ForEach(f func(k string, v Value)) {
	for _, field := range record.inner {
		f(field.name, field.value)
	}
}