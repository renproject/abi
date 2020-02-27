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

func NewList() List {
	return List{inner: make([]Value, 0, 0)}
}

func NewListWithCapacity(cap int) List {
	return List{inner: make([]Value, 0, cap)}
}

func (ls List) Marshal(w io.Writer) error {
	b := [4]byte{}
	binary.LittleEndian.PutUint32(b[:], uint32(len(ls.inner)))
	if _, err := w.Write(b[:]); err != nil {
		return err
	}
	for _, elem := range ls.inner {
		// Use the top-level marshaling function so that type information is
		// written.
		if err := Marshal(elem, w); err != nil {
			return err
		}
	}
	return nil
}

func (ls *List) Unmarshal(r io.Reader) error {
	_, err := ls.unmarshalAndReturnSize(r, MaxSize)
	return err
}

func (ls *List) unmarshalAndReturnSize(r io.Reader, maxSize uint32) (uint32, error) {
	len := uint32(0)
	if err := surge.Unmarshal(&len, r); err != nil {
		return 0, err
	}

	// Restrict the size of the abstract data type to prevent malicious input
	// forcing massive memory allocations.
	if len > maxSize {
		return 0, fmt.Errorf("expected size<=%v, got size=%v", maxSize, len)
	}
	ls.inner = make([]Value, 0, len)
	maxSize -= len
	if maxSize < 0 {
		return len, fmt.Errorf("exceeded maximum size")
	}

	sizeInBytes := len
	for i := uint32(0); i < len; i++ {
		// Use top-level unmarshaling function so that type information is read.
		elem, elemSize, err := unmarshalAndReturnSize(r, maxSize)
		if err != nil {
			return sizeInBytes, err
		}

		// Restrict the size of the abstract data type to prevent malicious input
		// forcing massive memory allocations.
		sizeInBytes += (elemSize - 1) // Offset by -1, because we started with sizeInBytes := len
		maxSize -= (elemSize - 1)     // Offset by -1, because we have already done maxSize -= len
		if maxSize < 0 {
			return sizeInBytes, fmt.Errorf("exceeded maximum size")
		}
		ls.inner = append(ls.inner, elem)
	}
	return sizeInBytes, nil
}

func (ls List) SizeHint() int {
	size := 4 // Size of length prefix
	for _, elem := range ls.inner {
		size += 2               // Size of element type prefix
		size += elem.SizeHint() // Size of element
	}
	return size
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
