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

type Value interface {
	surge.Marshaler
	Type() Type
}

func Unmarshal(r io.Reader, ty Type, m uint32) (Value, uint32, error) {
	switch ty {
	// Bytes
	case TypeString:
		v := String("")
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling string: %v", err)
		}
		return v, n, nil
	case TypeBytes:
		v := Bytes{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling bytes: %v", err)
		}
		return v, n, nil
	case TypeBytes32:
		v := Bytes32{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling bytes32: %v", err)
		}
		return v, n, nil

	// Data types
	case TypeU8:
		v := U8{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling u8: %v", err)
		}
		return v, n, nil
	case TypeU16:
		v := U16{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling u16: %v", err)
		}
		return v, n, nil
	case TypeU32:
		v := U32{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling u32: %v", err)
		}
		return v, n, nil
	case TypeU64:
		v := U64{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling u64: %v", err)
		}
		return v, n, nil
	case TypeU128:
		v := U128{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling u128: %v", err)
		}
		return v, n, nil
	case TypeU256:
		v := U256{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling u256: %v", err)
		}
		return v, n, nil

	// Abstract data types
	case TypeList:
		v := List{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling list: %v", err)
		}
		return v, n, nil
	case TypeRecord:
		v := Record{}
		n, err := v.Unmarshal(r, m)
		if err != nil {
			return nil, n, fmt.Errorf("error unmarshaling record: %v", err)
		}
		return v, n, nil
	}

	return nil, 0, fmt.Errorf("error unmarshaling value: unexpected type=%v", ty)
}

func UnmarshalJSON(data []byte, ty Type) (Value, error) {
	switch ty {
	// Bytes
	case TypeString:
		v := String("")
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling string: %v", err)
		}
		return v, nil
	case TypeBytes:
		v := Bytes{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling bytes: %v", err)
		}
		return v, nil
	case TypeBytes32:
		v := Bytes32{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling bytes32: %v", err)
		}
		return v, nil

	// Data types
	case TypeU8:
		v := U8{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling u8: %v", err)
		}
		return v, nil
	case TypeU16:
		v := U16{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling u16: %v", err)
		}
		return v, nil
	case TypeU32:
		v := U32{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling u32: %v", err)
		}
		return v, nil
	case TypeU64:
		v := U64{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling u64: %v", err)
		}
		return v, nil
	case TypeU128:
		v := U128{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling u128: %v", err)
		}
		return v, nil
	case TypeU256:
		v := U256{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling u256: %v", err)
		}
		return v, nil

	// Abstract data types
	case TypeList:
		v := List{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling list: %v", err)
		}
		return v, nil
	case TypeRecord:
		v := Record{}
		if err := v.UnmarshalJSON(data); err != nil {
			return nil, fmt.Errorf("error unmarshaling record: %v", err)
		}
		return v, nil
	}

	return nil, fmt.Errorf("error unmarshaling value: unexpected type=%v", ty)
}
