package abi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/renproject/surge"
)

// A String is a slice of bytes.
type String string

// Type returns the type identifier.
func (str String) Type() Type {
	return TypeString
}

// SizeHint returns the number of bytes required to represent a string in
// binary.
func (str String) SizeHint() int {
	return 4 + len(string(str))
}

// Marshal the string to binary. Marshaling will try to avoid allocating more
// than the specified maximum number of bytes. If it needs to allocate too many
// bytes, and error may be returned instead.
func (str String) Marshal(w io.Writer, m int) (int, error) {
	return surge.Marshal(w, string(str), m)
}

// Unmarshal the string from binary. Unmarshaling will not allocate more than
// the specified maximum number of bytes. If it needs to allocate too many
// bytes, and error is returned instead.
func (str *String) Unmarshal(r io.Reader, m int) (int, error) {
	return surge.Unmarshal(r, (*string)(str), m)
}

// MarshalJSON implements the JSON marshaler interface.
func (str String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(str))
}

// UnmarshalJSON implements the JSON unmarshaler interface.
func (str *String) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*string)(str))
}

type Bytes []byte

func (b Bytes) Type() Type {
	return TypeBytes
}

func (b Bytes) SizeHint() int {
	return surge.SizeHint([]byte(b))
}

func (b Bytes) Marshal(w io.Writer, m int) (int, error) {
	return surge.Marshal(w, []byte(b), m)
}

func (b *Bytes) Unmarshal(r io.Reader, m int) (int, error) {
	return surge.Unmarshal(r, (*[]byte)(b), m)
}

func (b Bytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Bytes) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	data, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(str)
	if err != nil {
		return err
	}
	*b = data
	return nil
}

func (b Bytes) String() string {
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(b)
}

type Bytes32 [32]byte

func (b32 Bytes32) Type() Type {
	return TypeBytes
}

func (b32 Bytes32) SizeHint() int {
	return 32
}

func (b32 Bytes32) Marshal(w io.Writer, m int) (int, error) {
	n, err := w.Write(b32[:])
	return m - n, err
}

func (b32 *Bytes32) Unmarshal(r io.Reader, m int) (int, error) {
	_, err := io.ReadFull(r, (*b32)[:])
	return m, err
}

func (b32 Bytes32) MarshalJSON() ([]byte, error) {
	return json.Marshal(b32.String())
}

func (b32 *Bytes32) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	data, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(str)
	if err != nil {
		return err
	}
	if len(data) != 32 {
		return fmt.Errorf("expected len=32, got len=%v", len(data))
	}
	copy((*b32)[:], data)
	return nil
}

func (b32 Bytes32) String() string {
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(b32[:])
}

type Bytes65 [65]byte

func (b65 Bytes65) Type() Type {
	return TypeBytes
}

func (b65 Bytes65) SizeHint() int {
	return 65
}

func (b65 Bytes65) Marshal(w io.Writer, m int) (int, error) {
	n, err := w.Write(b65[:])
	return m - n, err
}

func (b65 *Bytes65) Unmarshal(r io.Reader, m int) (int, error) {
	_, err := io.ReadFull(r, (*b65)[:])
	return m, err
}

func (b65 Bytes65) MarshalJSON() ([]byte, error) {
	return json.Marshal(b65.String())
}

func (b65 *Bytes65) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	data, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(str)
	if err != nil {
		return err
	}
	if len(data) != 65 {
		return fmt.Errorf("expected len=65, got len=%v", len(data))
	}
	copy((*b65)[:], data)
	return nil
}

func (b65 Bytes65) String() string {
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(b65[:])
}
