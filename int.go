package abi

import (
	"io"
	"math/big"

	"github.com/renproject/surge"
)

type U8 uint8

func (u8 U8) Marshal(w io.Writer) error {
	return surge.Marshal(uint8(u8), w)
}

func (u8 *U8) Unmarshal(r io.Reader) error {
	return surge.Unmarshal((*uint8)(u8), r)
}

func (u8 U8) SizeHint() int {
	return 1
}

func (U8) Type() Type {
	return TypeU8
}

type U16 uint16

func (u16 U16) Marshal(w io.Writer) error {
	return surge.Marshal(uint16(u16), w)
}

func (u16 *U16) Unmarshal(r io.Reader) error {
	return surge.Unmarshal((*uint16)(u16), r)
}

func (u16 U16) SizeHint() int {
	return 2
}

func (U16) Type() Type {
	return TypeU16
}

type U32 uint32

func (u32 U32) Marshal(w io.Writer) error {
	return surge.Marshal(uint32(u32), w)
}

func (u32 *U32) Unmarshal(r io.Reader) error {
	return surge.Unmarshal((*uint32)(u32), r)
}

func (u32 U32) SizeHint() int {
	return 4
}

func (U32) Type() Type {
	return TypeU32
}

type U64 uint64

func (u64 U64) Marshal(w io.Writer) error {
	return surge.Marshal(uint64(u64), w)
}

func (u64 *U64) Unmarshal(r io.Reader) error {
	return surge.Unmarshal((*uint64)(u64), r)
}

func (u64 U64) SizeHint() int {
	return 8
}

func (U64) Type() Type {
	return TypeU64
}

type U128 struct {
	inner *big.Int
}

func (u128 U128) Marshal(w io.Writer) error {
	return surge.Marshal(paddedTo16(u128.inner), w)
}

func (u128 *U128) Unmarshal(r io.Reader) error {
	b := [16]byte{}
	if err := surge.Unmarshal(&b, r); err != nil {
		return err
	}
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	u128.inner.SetBytes(b[:])
	return nil
}

func (u128 U128) SizeHint() int {
	return 16
}

func (U128) Type() Type {
	return TypeU128
}

type U256 struct {
	inner *big.Int
}

func (u256 U256) Marshal(w io.Writer) error {
	return surge.Marshal(paddedTo32(u256.inner), w)
}

func (u256 *U256) Unmarshal(r io.Reader) error {
	b := [32]byte{}
	if err := surge.Unmarshal(&b, r); err != nil {
		return err
	}
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	u256.inner.SetBytes(b[:])
	return nil
}

func (u256 U256) SizeHint() int {
	return 32
}

func (U256) Type() Type {
	return TypeU256
}

// paddedTo16 encodes a big integer as a big-endian into a 16-byte array. It
// will panic if the big integer is more than 16 bytes.
// Modified from:
// https://github.com/ethereum/go-ethereum/blob/master/common/math/big.go
// 17381ecc6695ea9c2d8e5ee0aee5cf70d59a301a
func paddedTo16(bigint *big.Int) [16]byte {
	if bigint.BitLen()/8 >= 16 {
		panic("too big")
	}
	ret := [16]byte{}
	readBits(bigint, ret[:])
	return ret
}

// paddedTo32 encodes a big integer as a big-endian into a 32-byte array. It
// will panic if the big integer is more than 32 bytes.
// Modified from:
// https://github.com/ethereum/go-ethereum/blob/master/common/math/big.go
// 17381ecc6695ea9c2d8e5ee0aee5cf70d59a301a
func paddedTo32(bigint *big.Int) [32]byte {
	if bigint.BitLen()/8 >= 32 {
		panic("too big")
	}
	ret := [32]byte{}
	readBits(bigint, ret[:])
	return ret
}

// readBits encodes the absolute value of bigint as big-endian bytes. Callers
// must ensure that buf has enough space. If buf is too short the result will be
// incomplete.
// Modified from:
// https://github.com/ethereum/go-ethereum/blob/master/common/math/big.go
// 17381ecc6695ea9c2d8e5ee0aee5cf70d59a301a
func readBits(bigint *big.Int, buf []byte) {
	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}

const (
	// wordBits is the number of bits in a big word.
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// wordBytes is the number of bytes in a big word.
	wordBytes = wordBits / 8
)
