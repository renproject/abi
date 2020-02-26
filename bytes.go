package abi

import (
	"fmt"
	"io"

	"github.com/renproject/surge"
)

type Bytes struct {
	inner []byte
}

func (bytes Bytes) Marshal(w io.Writer) error {
	len := uint32(len(bytes.inner))
	if err := surge.Marshal(uint32(len), w); err != nil {
		return err
	}
	_, err := w.Write(bytes.inner)
	return err
}

func (bytes *Bytes) Unmarshal(r io.Reader) error {
	len := uint32(0)
	if err := surge.Unmarshal(&len, r); err != nil {
		return err
	}
	if len > MaxBytesLength {
		return fmt.Errorf("expected len<=%v, got len=%v", MaxBytesSize, len)
	}
	*v = make([]byte, len)
	if _, err := io.ReadFull(r, *v); err != nil {
		return err
	}
	return nil
}

func (bytes Bytes) SizeHint() int {
	return 4 + len(bytes.inner) // Length prefix + number of bytes in the slice
}

func (bytes Bytes) Type() Type {
	return TypeBytes
}

func (bytes Bytes) Equal(other Value) bool {
	if v, ok := other.(Bytes); ok {
		return bytes.Equal(bytes.inner, v.inner)
	}
	return false
}
