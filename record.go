package abi

import (
	"encoding/binary"
	"io"
	"sort"

	"github.com/renproject/surge"
)

type Record struct {
	inner []field
}

func NewRecord() Record {
	return Record{inner: make([]field, 0)}
}

func NewRecordWithCapacity(cap int) Record {
	return Record{inner: make([]field, 0, cap)}
}

func (record Record) Marshal(w io.Writer) (uint32, error) {
	// Marshal length.
	b := [4]byte{}
	binary.LittleEndian.PutUint32(b[:], uint32(len(record.inner)))
	n, err := w.Write(b[:])
	if err != nil {
		return uint32(n), err
	}

	// Marshal elements.
	n1 := uint32(n)
	for _, field := range record.inner {
		// Marshal field name.
		n2, err := field.name.Marshal(w)
		n1 += n2
		if err != nil {
			return n1, err
		}
		// Marshal field type.
		n3, err := field.value.Type().Marshal(w)
		n1 += n3
		if err != nil {
			return n1, err
		}
		// Marshal field value.
		n4, err := field.value.Marshal(w)
		n1 += n4
		if err != nil {
			return n1, err
		}
	}
	return n1, nil
}

func (record *Record) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	// Unmarshal length.
	len := uint32(0)
	n1, err := surge.Unmarshal(&len, r, m)
	if err != nil {
		return n1, err
	}
	m -= n1
	if m < len {
		return n1, surge.ErrMaxBytesExceeded
	}

	// Unmarshal elements.
	record.inner = make([]field, 0, len)
	for i := uint32(0); i < len; i++ {
		// Unmarshal field name.
		var name String
		n2, err := name.Unmarshal(r, m)
		n1 += n2
		m -= n2
		if err != nil {
			return n1, err
		}
		// Unmarshal field type.
		var ty Type
		n3, err := ty.Unmarshal(r, m)
		n1 += n3
		m -= n3
		if err != nil {
			return n1, err
		}
		// Unmarshal field value.
		value, n4, err := Unmarshal(r, ty, m)
		n1 += n4
		m -= n4
		if err != nil {
			return n1, err
		}
		record.inner = append(record.inner, field{name: name, value: value})
	}
	return n1, nil
}

func (record Record) SizeHint() uint32 {
	size := uint32(4) // Size of length prefix
	for _, field := range record.inner {
		size += field.name.SizeHint()         // Size of field name
		size += field.value.Type().SizeHint() // Size of field type
		size += field.value.SizeHint()        // Size of field value
	}
	return size
}

func (record Record) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (record *Record) UnmarshalJSON(data []byte) error {
	return nil
}

func (record Record) Type() Type {
	return TypeRecord
}

func (record *Record) Set(k string, v Value) {
	i := sort.Search(len(record.inner), func(i int) bool {
		return string(record.inner[i].name) >= k
	})
	if i < 0 {
		record.inner = append(record.inner, field{name: String(k), value: v})
		return
	}
	if string(record.inner[i].name) == k {
		record.inner[i].value = v
		return
	}
	record.inner = append(record.inner[:i], append([]field{field{name: String(k), value: v}}, record.inner[i:]...)...)
}

func (record *Record) Remove(k string) {
	i := sort.Search(len(record.inner), func(i int) bool {
		return string(record.inner[i].name) >= k
	})
	if i < 0 {
		return
	}
	record.inner = append(record.inner[:i], record.inner[i+1:]...)
}

func (record Record) Get(k string) (Value, bool) {
	i := sort.Search(len(record.inner), func(i int) bool {
		return string(record.inner[i].name) >= k
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
		f(string(field.name), field.value)
	}
}

type field struct {
	name  String
	value Value
}
