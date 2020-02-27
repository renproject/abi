package abi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/renproject/surge"
)

type Bytes []byte

func (b Bytes) Marshal(w io.Writer) error {
	len := uint32(len(b))
	if err := surge.Marshal(uint32(len), w); err != nil {
		return err
	}
	_, err := w.Write(b)
	return err
}

func (b *Bytes) Unmarshal(r io.Reader) error {
	_, err := b.unmarshalAndReturnLength(r)
	return err
}

func (b *Bytes) unmarshalAndReturnLength(r io.Reader) (uint32, error) {
	len := uint32(0)
	if err := surge.Unmarshal(&len, r); err != nil {
		return len, err
	}
	if len > MaxSize {
		return len, fmt.Errorf("expected len<=%v, got len=%v", MaxSize, len)
	}
	*b = make([]byte, len)
	if _, err := io.ReadFull(r, *b); err != nil {
		return len, err
	}
	return len, nil
}

func (b Bytes) SizeHint() int {
	return 4 + len(b) // Length prefix + number of bytes in the slice
}

func (b Bytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Bytes) UnmarshalJSON(data []byte) error {
	var bString string
	if err := json.Unmarshal(data, &bString); err != nil {
		return err
	}
	data, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(bString)
	if err != nil {
		return err
	}
	*b = data
	return nil
}

func (b Bytes) Type() Type {
	return TypeBytes
}

func (b Bytes) Len() int {
	return len(b)
}

func (b Bytes) String() string {
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}

type Bytes32 [32]byte

func (b Bytes32) Marshal(w io.Writer) error {
	_, err := w.Write(b[:])
	return err
}

func (b *Bytes32) Unmarshal(r io.Reader) error {
	if _, err := io.ReadFull(r, (*b)[:]); err != nil {
		return err
	}
	return nil
}

func (b Bytes32) SizeHint() int {
	return 32
}

func (b Bytes32) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Bytes32) UnmarshalJSON(data []byte) error {
	var bString string
	if err := json.Unmarshal(data, &bString); err != nil {
		return err
	}
	data, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(bString)
	if err != nil {
		return err
	}
	if len(data) != 32 {
		return fmt.Errorf("expected len=32, got len=%v", len(data))
	}
	copy(b[:], data)
	return nil
}

func (b Bytes32) Type() Type {
	return TypeBytes32
}

func (b Bytes32) String() string {
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(b[:])
}
