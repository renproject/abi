package abi

import (
	"fmt"
	"io"

	"github.com/renproject/surge"
)

const (
	MaxBytesLength          = 1000000 // 1MB is the maximum number of bytes.
	MaxAbstractDataTypeSize = 10000   // 10K elements is the most elements that an abstract data type can have.
)

type Type uint16

const (
	// Bytes
	TypeBytes = Type(1)

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
	surge.SizeHinter
	Type() Type
}

func Marshal(v Value, w io.Writer) error {
	if err := surge.Marshal(uint16(v.Type()), w); err != nil {
		return fmt.Errorf("error marshaling type: %v", err)
	}
	if err := v.Interface.Marshal(w); err != nil {
		return fmt.Errorf("error marshaling interface: %v", err)
	}
	return nil
}

func Unmarshal(r io.Reader) (Value, error) {
	v, _, err := unmarshal(r, true)
	return v, err
}

func unmarshal(r io.Reader, abstractDataTypeSizeLimit int) (Value, int, error) {
	ty := uint16(0)
	if err := surge.Unmarshal(&ty, r); err != nil {
		return fmt.Errorf("error unmarshaling type: %v", err)
	}

	switch Type(ty) {
	// Bytes
	case Bytes:
		v := Bytes{}
		if err := v.Unmarshal(r); err != nil {
			return nil, fmt.Errorf("error unmarshaling bytes: %v", err)
		}
		return v, 0, nil

	// Data types
	case U8:
		v := U8{}
		if err := v.Unmarshal(r); err != nil {
			return nil, fmt.Errorf("error unmarshaling u8: %v", err)
		}
		return v, 0, nil
	case U16:
		v := U16{}
		if err := v.Unmarshal(r); err != nil {
			return nil, fmt.Errorf("error unmarshaling u16: %v", err)
		}
		return v, 0, nil
	case U32:
		v := U32{}
		if err := v.Unmarshal(r); err != nil {
			return nil, fmt.Errorf("error unmarshaling u32: %v", err)
		}
		return v, 0, nil
	case U64:
		v := U64{}
		if err := v.Unmarshal(r); err != nil {
			return nil, fmt.Errorf("error unmarshaling u64: %v", err)
		}
		return v, 0, nil
	case U128:
		v := U128{}
		if err := v.Unmarshal(r); err != nil {
			return nil, fmt.Errorf("error unmarshaling u128: %v", err)
		}
		return v, 0, nil
	case U256:
		v := U256{}
		if err := v.Unmarshal(r); err != nil {
			return nil, fmt.Errorf("error unmarshaling u256: %v", err)
		}
		return v, 0, nil

	// Abstract data types
	case TypeList:
		v := List{}
		size, err := v.unmarshalAbstractDataType(r, abstractDataTypeSizeLimit)
		if err != nil {
			return nil, 0, fmt.Errorf("error unmarshaling list: %v", err)
		}
		return v, size, nil
	case TypeRecord:
		v := Record{}
		size, err := v.unmarshalAbstractDataType(r, abstractDataTypeSizeLimit)
		if err != nil {
			return nil, 0, fmt.Errorf("error unmarshaling list: %v", err)
		}
		return v, size, nil
	}

	return nil, fmt.Errorf("error unmarshaling value: unexpected type=%v", ty)
}
