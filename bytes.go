package abi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/renproject/surge"
)

type String string

func (str String) Marshal(w io.Writer) (uint32, error) {
	len := uint32(len(str))
	n1, err := surge.Marshal(len, w)
	if err != nil {
		return n1, err
	}
	n2, err := w.Write([]byte(str))
	if err != nil {
		return n1 + uint32(n2), err
	}
	return n1 + uint32(n2), err
}

func (str *String) Unmarshal(r io.Reader, m uint32) (uint32, error) {
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
	// Unmarshal data.
	b := make([]byte, len)
	n2, err := io.ReadFull(r, b)
	if err != nil {
		return n1 + uint32(n2), err
	}
	*str = String(string(b))
	return n1 + uint32(n2), nil
}

func (str String) SizeHint() uint32 {
	return uint32(4 + len(str)) // Length prefix + number of bytes in the string
}

func (str String) MarshalJSON() ([]byte, error) {
	return json.Marshal(str)
}

func (str *String) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*string)(str))
}

func (str String) Type() Type {
	return TypeString
}

type Bytes []byte

func (b Bytes) Marshal(w io.Writer) (uint32, error) {
	len := uint32(len(b))
	n1, err := surge.Marshal(len, w)
	if err != nil {
		return n1, err
	}
	n2, err := w.Write(b)
	return n1 + uint32(n2), err
}

func (b *Bytes) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	len := uint32(0)
	n1, err := surge.Unmarshal(&len, r, m)
	if err != nil {
		return n1, err
	}
	m -= n1
	if m < len {
		return n1, surge.ErrMaxBytesExceeded
	}
	*b = make([]byte, len)
	n2, err := io.ReadFull(r, *b)
	if err != nil {
		return n1 + uint32(n2), err
	}
	return n1 + uint32(n2), nil
}

func (b Bytes) SizeHint() uint32 {
	return uint32(4 + len(b)) // Length prefix + number of bytes in the slice
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

func (b Bytes32) Marshal(w io.Writer) (uint32, error) {
	n, err := w.Write(b[:])
	return uint32(n), err
}

func (b *Bytes32) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	n, err := io.ReadFull(r, (*b)[:])
	return uint32(n), err
}

func (b Bytes32) SizeHint() uint32 {
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
