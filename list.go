package abi

import (
	"encoding/binary"
	"io"

	"github.com/renproject/surge"
)

type List struct {
	inner []Value
}

func NewList() List {
	return List{inner: make([]Value, 0, 0)}
}

func NewListWithCapacity(cap int) List {
	return List{inner: make([]Value, 0, cap)}
}

func (ls List) Marshal(w io.Writer) (uint32, error) {
	// Marshal length.
	b := [4]byte{}
	binary.LittleEndian.PutUint32(b[:], uint32(len(ls.inner)))
	n, err := w.Write(b[:])
	if err != nil {
		return uint32(n), err
	}
	// Marshal elements,
	n1 := uint32(n)
	for _, elem := range ls.inner {
		// Marshal type.
		n2, err := elem.Type().Marshal(w)
		n1 += n2
		if err != nil {
			return n1, err
		}
		// Marshal value.
		n3, err := elem.Marshal(w)
		n1 += n3
		if err != nil {
			return n1, err
		}
	}
	return n1, nil
}
func (ls *List) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	// Unmarshal the length.
	len := uint32(0)
	n1, err := surge.Unmarshal(&len, r, m)
	if err != nil {
		return n1, err
	}
	m -= n1
	if m < n1 {
		return n1, surge.ErrMaxBytesExceeded
	}
	ls.inner = make([]Value, 0, len)

	// Unmarshal elements.
	for i := uint32(0); i < len; i++ {
		// Unmarshal type.
		var ty Type
		n2, err := ty.Unmarshal(r, m)
		n1 += n2
		m -= n2
		if err != nil {
			return n1, err
		}

		// Unmarshal value
		v, n3, err := Unmarshal(r, ty, m)
		n1 += n3
		m -= n3
		if err != nil {
			return n1, err
		}
		ls.inner = append(ls.inner, v)
	}
	return n1, nil
}

func (ls List) SizeHint() uint32 {
	size := uint32(4) // Size of length prefix
	for _, elem := range ls.inner {
		size += elem.Type().SizeHint() // Size of element type
		size += elem.SizeHint()        // Size of element value
	}
	return size
}

func (ls List) MarshalJSON() ([]byte, error) {
	panic("unimplemented")
}

func (ls *List) UnmarshalJSON(data []byte) error {
	panic("unimplemented")
}

func (List) Type() Type {
	return TypeList
}

func (ls *List) Append(v Value) {
	ls.inner = append(ls.inner, v)
}

func (ls *List) AppendAll(other List) {
	ls.inner = append(ls.inner, other.inner...)
}

func (ls *List) Insert(i int, v Value) {
	ls.inner = append(ls.inner[:i], append([]Value{v}, ls.inner[i:]...)...)
}

func (ls *List) Remove(i int) {
	ls.inner = append(ls.inner[:i], ls.inner[i+1:]...)
}

func (ls *List) RemoveAll() {
	ls.inner = ls.inner[0:0]
}

func (ls List) Get(i int) Value {
	return ls.inner[i]
}

func (ls List) Len() int {
	return len(ls.inner)
}

func (ls List) ForEach(f func(i int, v Value)) {
	for i, v := range ls.inner {
		f(i, v)
	}
}
