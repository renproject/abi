package abi

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/renproject/surge"
)

type List struct {
	inner []Value
}

func (ls List) Marshal(w io.Writer) error {
	b := [4]byte{}
	binary.LittleEndian.PutUint32(b[:], uint32(len(ls.inner)))
	if err := w.Write(b[:]); err != nil {
		return err
	}
	for _, elem := range ls.inner {
		if err := Marshal(elem, w); err != nil {
			return err
		}
	}
	return nil
}

func (ls *List) Unmarshal(r io.Reader) error {
	_, err := ls.unmarshalAbstractDataType(r, MaxAbstractDataTypeSize)
	return err
}

func (ls *List) unmarshalAbstractDataType(r io.Reader, abstractDataTypeSizeLimit int) (int, error) {
	size := uint32(0)
	if err := Unmarshal(&size, r); err != nil {
		return size, err
	}

	// Restrict the size of the abstract data type to prevent forced OOM
	// allocations.
	if size > abstractDataTypeSizeLimit {
		return size, fmt.Errorf("expected size<=%v, got size=%v", MaxListLen, size)
	}
	abstractDataTypeSizeLimit -= size

	ls.inner := make([]Value, size)
	for i := 0; i < len; i++ {
		elem, abstractDataTypeSize, err := unmarshal(r, abstractDataTypeSizeLimit)
		if err != nil {
			return size, err
		}
		// Handle the size of abstract data types within abstract data types.
		abstractDataTypeSizeLimit -= abstractDataTypeSize
		if abstractDataTypeSizeLimit < 0 {
			return size, fmt.Errorf("exceeded maximum abstract data size")
		}
		ls.inner[i] = elem
	}

	return size, nil
}

func (ls List) SizeHint() int {
	size := 4 // Size of length prefix
	for _, elem := range ls {
		size += elem.SizeHint()
	}
	return size
}

func (List) Type() Type {
	return TypeList
}

func (ls *List) Append(v Value) {
	ls.inner = append(ls.inner, v)
}

func (ls *List) Extend(other List) {
	ls.inner = append(ls.inner, other.inner...)
}

func (ls *List) Insert(i int, v Value) {
	ls.inner = append(ls.inner[:i], append([]Value{Value}, ls.inner[i:]...)...)
}

func (ls *List) Remove(i int) {
	ls.inner = append(ls.inner[:i], ls.inner[i+1:]...)
}

func (ls List) Get(i int) Value {
	return ls.inner[i]
}

func (ls List) Size() int {
	return len(ls.inner)
}

func (ls List) ForEach(f func(i int, v Value)) {
	for i, v := range ls.inner {
		f(i, v)
	}
}