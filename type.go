package abi

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/renproject/surge"
)

// Type enumerations for all core ABI types. They are categorised into bytes,
// core data types, and abstract data types.
const (
	TypeUnrecognised = Type(0)

	// Bytes
	TypeString  = Type(1)
	TypeBytes   = Type(2)
	TypeBytes32 = Type(3)

	// Data types
	TypeU8   = Type(11)
	TypeU16  = Type(12)
	TypeU32  = Type(13)
	TypeU64  = Type(14)
	TypeU128 = Type(15)
	TypeU256 = Type(16)

	// Abstract data types
	TypeList   = Type(101)
	TypeRecord = Type(102)
)

type Type uint16

func NewTypeFromUint16(i uint16) (Type, bool) {
	switch Type(i) {
	case TypeString, TypeBytes, TypeBytes32, TypeU8, TypeU16, TypeU32, TypeU64, TypeU128, TypeU256, TypeList, TypeRecord:
		return Type(i), true
	default:
		return TypeUnrecognised, false
	}
}

func NewTypeFromString(str string) (Type, bool) {
	switch str {
	case TypeString.String():
		return TypeString, true
	case TypeBytes.String():
		return TypeBytes, true
	case TypeBytes32.String():
		return TypeBytes32, true

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

	case TypeList.String():
		return TypeList, true
	case TypeRecord.String():
		return TypeRecord, true
	default:
		return TypeUnrecognised, false
	}
}

func (ty Type) Marshal(w io.Writer) (uint32, error) {
	return surge.Marshal(uint16(ty), w)
}

func (ty *Type) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	var ok bool
	var i uint16
	n, err := surge.Unmarshal(&i, r, m)
	if err != nil {
		return n, err
	}
	*ty, ok = NewTypeFromUint16(i)
	if !ok {
		return n, fmt.Errorf("unexpected type=%v", i)
	}
	return n, nil
}

func (ty Type) SizeHint() uint32 {
	return 2
}

func (ty Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(ty.String())
}

func (ty *Type) UnmarshalJSON(data []byte) error {
	var ok bool
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*ty, ok = NewTypeFromString(str)
	if !ok {
		return fmt.Errorf("unexpected type=%v", str)
	}
	return nil
}

func (ty Type) String() string {
	switch ty {
	case TypeString:
		return "string"
	case TypeBytes:
		return "bytes"
	case TypeBytes32:
		return "bytes32"

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

	case TypeList:
		return "list"
	case TypeRecord:
		return "record"
	}

	return "unrecognised"
}
