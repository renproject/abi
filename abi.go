package abi

import (
	"fmt"
	"io"

	"github.com/renproject/surge"
)

// MaxSize in bytes that an ABI object can consume. Currently, this is set
// to 32MB. For abstract data types, the size of the object includes all
// inner objects.
var MaxSize uint32 = 32 * 1024 * 1024

type Type uint16

const (
	// Bytes
	TypeBytes   = Type(1)
	TypeBytes32 = Type(2)

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

type Value interface {
	surge.Marshaler
	Type() Type
}

func Marshal(v Value, w io.Writer) error {
	if err := surge.Marshal(uint16(v.Type()), w); err != nil {
		return fmt.Errorf("error marshaling type: %v", err)
	}
	if err := v.Marshal(w); err != nil {
		return fmt.Errorf("error marshaling interface: %v", err)
	}
	return nil
}

func Unmarshal(r io.Reader) (Value, error) {
	v, _, err := unmarshalAndReturnSize(r, MaxSize)
	return v, err
}

func unmarshalAndReturnSize(r io.Reader, maxSize uint32) (Value, uint32, error) {
	ty := uint16(0)
	if err := surge.Unmarshal(&ty, r); err != nil {
		return nil, 0, fmt.Errorf("error unmarshaling type: %v", err)
	}

	switch Type(ty) {
	// Bytes
	case TypeBytes:
		v := Bytes{}
		len, err := v.unmarshalAndReturnLength(r)
		if err != nil {
			return nil, len, fmt.Errorf("error unmarshaling bytes: %v", err)
		}
		return v, len, nil
	case TypeBytes32:
		v := Bytes32{}
		if err := v.Unmarshal(r); err != nil {
			return nil, 32, fmt.Errorf("error unmarshaling bytes: %v", err)
		}
		return v, 32, nil

	// Data types
	case TypeU8:
		v := U8{}
		if err := v.Unmarshal(r); err != nil {
			return nil, 1, fmt.Errorf("error unmarshaling u8: %v", err)
		}
		return v, 1, nil
	case TypeU16:
		v := U16{}
		if err := v.Unmarshal(r); err != nil {
			return nil, 2, fmt.Errorf("error unmarshaling u16: %v", err)
		}
		return v, 2, nil
	case TypeU32:
		v := U32{}
		if err := v.Unmarshal(r); err != nil {
			return nil, 4, fmt.Errorf("error unmarshaling u32: %v", err)
		}
		return v, 4, nil
	case TypeU64:
		v := U64{}
		if err := v.Unmarshal(r); err != nil {
			return nil, 8, fmt.Errorf("error unmarshaling u64: %v", err)
		}
		return v, 8, nil
	case TypeU128:
		v := U128{}
		if err := v.Unmarshal(r); err != nil {
			return nil, 16, fmt.Errorf("error unmarshaling u128: %v", err)
		}
		return v, 16, nil
	case TypeU256:
		v := U256{}
		if err := v.Unmarshal(r); err != nil {
			return nil, 32, fmt.Errorf("error unmarshaling u256: %v", err)
		}
		return v, 32, nil

	// Abstract data types
	case TypeList:
		v := List{}
		size, err := v.unmarshalAndReturnSize(r, maxSize)
		if err != nil {
			return nil, size, fmt.Errorf("error unmarshaling list: %v", err)
		}
		return v, size, nil
	case TypeRecord:
		v := Record{}
		size, err := v.unmarshalAndReturnSize(r, maxSize)
		if err != nil {
			return nil, size, fmt.Errorf("error unmarshaling record: %v", err)
		}
		return v, size, nil
	}

	return nil, 0, fmt.Errorf("error unmarshaling value: unexpected type=%v", ty)
}
