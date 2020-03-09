package abi

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/renproject/surge"
)

// MaxBytes of an ABI object. Defaults to 32 MB.
var MaxBytes = 32 * 1024 * 1024

// A Type identifier is used to identify types over storage/network boundaries.
// This is necessary when the receiver of a value does not know ahead-of-time
// what type of value to expect.
type Type uint16

// Type identifiers for all core ABI types. They are categorised into bytes,
// scalar types, and abstract data types.
const (
	// Nil
	TypeNil = Type(0)

	// Bytes
	TypeString  = Type(1)
	TypeBytes   = Type(2)
	TypeBytes32 = Type(3)
	TypeBytes65 = Type(4)

	// Scalar types
	TypeBool = Type(11)
	TypeU8   = Type(12)
	TypeU16  = Type(13)
	TypeU32  = Type(14)
	TypeU64  = Type(15)
	TypeU128 = Type(16)
	TypeU256 = Type(17)

	// Abstract data types
	TypeMaybe  = Type(101)
	TypeList   = Type(102)
	TypeRecord = Type(103)
)

// NewTypeFromUint16 converts a uint16 into a Type. Returns false if the
// conversion fails.
func NewTypeFromUint16(i uint16) (Type, bool) {
	switch Type(i) {
	case TypeString, TypeBytes, TypeBytes32, TypeBytes65:
		return Type(i), true
	case TypeBool, TypeU8, TypeU16, TypeU32, TypeU64, TypeU128, TypeU256:
		return Type(i), true
	case TypeMaybe, TypeList, TypeRecord:
		return Type(i), true
	default:
		return TypeNil, false
	}
}

// NewTypeFromString converts a string into a Type. Returns false if the
// conversion fails.
func NewTypeFromString(str string) (Type, bool) {
	switch str {
	case TypeString.String():
		return TypeString, true
	case TypeBytes.String():
		return TypeBytes, true
	case TypeBytes32.String():
		return TypeBytes32, true
	case TypeBytes65.String():
		return TypeBytes65, true

	case TypeBool.String():
		return TypeBool, true
	case TypeU8.String():
		return TypeU8, true
	case TypeU16.String():
		return TypeU16, true
	case TypeU32.String():
		return TypeU32, true
	case TypeU64.String():
		return TypeU64, true
	case TypeU128.String():
		return TypeU128, true
	case TypeU256.String():
		return TypeU256, true

	case TypeMaybe.String():
		return TypeMaybe, true
	case TypeList.String():
		return TypeList, true
	case TypeRecord.String():
		return TypeRecord, true

	default:
		return TypeNil, false
	}
}

// SizeHint returns the number of bytes required to represent a Type in binary.
func (Type) SizeHint() int {
	return 2
}

// Marshal the Type to binary. Marshaling will try to avoid allocating more than
// the specified maximum number of bytes. If it needs to allocate too many
// bytes, and error may be returned instead.
func (ty Type) Marshal(w io.Writer, m int) (int, error) {
	return surge.Marshal(w, uint16(ty), m)
}

// Unmarshal the Type from binary. Unmarshaling will not allocate more than the
// specified maximum number of bytes. If it needs to allocate too many bytes,
// and error is returned instead.
func (ty *Type) Unmarshal(r io.Reader, m int) (int, error) {
	var ok bool
	var i uint16
	m, err := surge.Unmarshal(r, &i, m)
	if err != nil {
		return m, err
	}
	*ty, ok = NewTypeFromUint16(i)
	if !ok {
		return m, fmt.Errorf("non-exhaustive pattern: Type(%v)", i)
	}
	return m, nil
}

// MarshalJSON implements the JSON marshaler interface by marshaling the Type
// into a string.
func (ty Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(ty.String())
}

// UnmarshalJSON implements the JSON unmarshaler interface by unmarshaling the
// Type from a string.
func (ty *Type) UnmarshalJSON(data []byte) error {
	var ok bool
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*ty, ok = NewTypeFromString(str)
	if !ok {
		return fmt.Errorf("non-exhaustive pattern: Type(%v)", str)
	}
	return nil
}

func (ty Type) String() string {
	switch ty {
	case TypeString:
		return "str"
	case TypeBytes:
		return "b"
	case TypeBytes32:
		return "b32"
	case TypeBytes65:
		return "b65"

	case TypeBool:
		return "bool"
	case TypeU8:
		return "u8"
	case TypeU16:
		return "u16"
	case TypeU32:
		return "u32"
	case TypeU64:
		return "u64"
	case TypeU128:
		return "u128"
	case TypeU256:
		return "u256"

	case TypeMaybe:
		return "maybe"
	case TypeList:
		return "list"
	case TypeRecord:
		return "record"
	}
	return "nil"
}

// A Value must be able to identify its ABI-compatible type identifier.
type Value interface {
	surge.Marshaler
	json.Marshaler
}

// SizeHint returns the number of bytes required to represent a value in binary.
func SizeHint(v Value) int {
	return 2 + v.SizeHint()
}
