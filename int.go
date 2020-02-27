package abi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strconv"

	"github.com/renproject/surge"
)

// U8 represents an 8-bit unsigned integer.
type U8 struct {
	inner uint8
}

// NewU8 returns a uint8 wrapped as a U8.
func NewU8(x uint8) U8 {
	return U8{inner: x}
}

// Marshal to binary. Returns the number of bytes written to the io.Writer.
func (u8 U8) Marshal(w io.Writer) (uint32, error) {
	return surge.Marshal(u8.inner, w)
}

// Unmarshal from binary, with a restriction on the maximum memory allocation.
// Returns the number of bytes read from the io.Reader.
func (u8 *U8) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	return surge.Unmarshal(&u8.inner, r, m)
}

// SizeHint returns the number of bytes that a U8 requires in its binary
// representation.
func (u8 U8) SizeHint() uint32 {
	return 1
}

// MarshalJSON implements the json.Marshaler interface. U8s are marshaled as
// decimal strings (for consistency with larger integer types).
func (u8 U8) MarshalJSON() ([]byte, error) {
	return json.Marshal(u8.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u8 *U8) UnmarshalJSON(data []byte) error {
	var xString string
	if err := json.Unmarshal(data, &xString); err != nil {
		return err
	}
	x, err := strconv.ParseUint(xString, 10, 8)
	if err != nil {
		return err
	}
	u8.inner = uint8(x)
	return nil
}

// Type returns the ABI object type of a U8.
func (U8) Type() Type {
	return TypeU8
}

// Uint8 returns the inner uint8 value.
func (u8 U8) Uint8() uint8 {
	return u8.inner
}

// String implements the fmt.Stringer interface. U8s are printed as decimal
// strings.
func (u8 U8) String() string {
	return fmt.Sprintf("%v", u8.inner)
}

func (u8 U8) Add(other U8) U8 {
	ret := U8{inner: u8.inner + other.inner}
	if ret.inner < u8.inner {
		panic("overflow")
	}
	return ret
}

func (u8 U8) Sub(other U8) U8 {
	ret := U8{inner: u8.inner - other.inner}
	if ret.inner > u8.inner {
		panic("underflow")
	}
	return ret
}

func (u8 *U8) AddAssign(other U8) {
	u8.inner = u8.inner + other.inner
	if u8.inner < other.inner {
		panic("overflow")
	}
}

func (u8 *U8) SubAssign(other U8) {
	u8.inner = u8.inner - other.inner
	if u8.inner > other.inner {
		panic("underflow")
	}
}

func (u8 U8) Equal(other U8) bool {
	return u8.inner == other.inner
}

type U16 struct {
	inner uint16
}

func NewU16(x uint16) U16 {
	return U16{inner: x}
}

func (u16 U16) Marshal(w io.Writer) (uint32, error) {
	return surge.Marshal(u16.inner, w)
}

func (u16 *U16) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	return surge.Unmarshal(&u16.inner, r, m)
}

func (u16 U16) SizeHint() uint32 {
	return 2
}

func (u16 U16) MarshalJSON() ([]byte, error) {
	return json.Marshal(u16.String())
}

func (u16 *U16) UnmarshalJSON(data []byte) error {
	var xString string
	if err := json.Unmarshal(data, &xString); err != nil {
		return err
	}
	x, err := strconv.ParseUint(xString, 10, 16)
	if err != nil {
		return err
	}
	u16.inner = uint16(x)
	return nil
}

func (U16) Type() Type {
	return TypeU16
}

func (u16 U16) Uint16() uint16 {
	return u16.inner
}

func (u16 U16) String() string {
	return fmt.Sprintf("%v", u16.inner)
}

func (u16 U16) Add(other U16) U16 {
	ret := U16{inner: u16.inner + other.inner}
	if ret.inner < u16.inner {
		panic("overflow")
	}
	return ret
}

func (u16 U16) Sub(other U16) U16 {
	ret := U16{inner: u16.inner - other.inner}
	if ret.inner > u16.inner {
		panic("underflow")
	}
	return ret
}

func (u16 *U16) AddAssign(other U16) {
	u16.inner = u16.inner + other.inner
	if u16.inner < other.inner {
		panic("overflow")
	}
}

func (u16 *U16) SubAssign(other U16) {
	u16.inner = u16.inner - other.inner
	if u16.inner > other.inner {
		panic("underflow")
	}
}

func (u16 U16) Equal(other U16) bool {
	return u16.inner == other.inner
}

type U32 struct {
	inner uint32
}

func NewU32(x uint32) U32 {
	return U32{inner: x}
}

func (u32 U32) Marshal(w io.Writer) (uint32, error) {
	return surge.Marshal(u32.inner, w)
}

func (u32 *U32) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	return surge.Unmarshal(&u32.inner, r, m)
}

func (u32 U32) SizeHint() uint32 {
	return 4
}

func (u32 U32) MarshalJSON() ([]byte, error) {
	return json.Marshal(u32.String())
}

func (u32 *U32) UnmarshalJSON(data []byte) error {
	var xString string
	if err := json.Unmarshal(data, &xString); err != nil {
		return err
	}
	x, err := strconv.ParseUint(xString, 10, 32)
	if err != nil {
		return err
	}
	u32.inner = uint32(x)
	return nil
}

func (U32) Type() Type {
	return TypeU32
}

func (u32 U32) Uint32() uint32 {
	return u32.inner
}

func (u32 U32) String() string {
	return fmt.Sprintf("%v", u32.inner)
}

func (u32 U32) Add(other U32) U32 {
	ret := U32{inner: u32.inner + other.inner}
	if ret.inner < u32.inner {
		panic("overflow")
	}
	return ret
}

func (u32 U32) Sub(other U32) U32 {
	ret := U32{inner: u32.inner - other.inner}
	if ret.inner > u32.inner {
		panic("underflow")
	}
	return ret
}

func (u32 *U32) AddAssign(other U32) {
	u32.inner = u32.inner + other.inner
	if u32.inner < other.inner {
		panic("overflow")
	}
}

func (u32 *U32) SubAssign(other U32) {
	u32.inner = u32.inner - other.inner
	if u32.inner > other.inner {
		panic("underflow")
	}
}

func (u32 U32) Equal(other U32) bool {
	return u32.inner == other.inner
}

type U64 struct {
	inner uint64
}

func NewU64(x uint64) U64 {
	return U64{inner: x}
}

func (u64 U64) Marshal(w io.Writer) (uint32, error) {
	return surge.Marshal(u64.inner, w)
}

func (u64 *U64) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	return surge.Unmarshal(&u64.inner, r, m)
}

func (u64 U64) SizeHint() uint32 {
	return 8
}

func (u64 U64) MarshalJSON() ([]byte, error) {
	return json.Marshal(u64.String())
}

func (u64 *U64) UnmarshalJSON(data []byte) error {
	var xString string
	if err := json.Unmarshal(data, &xString); err != nil {
		return err
	}
	x, err := strconv.ParseUint(xString, 10, 64)
	if err != nil {
		return err
	}
	u64.inner = x
	return nil
}

func (U64) Type() Type {
	return TypeU64
}

func (u64 U64) Uint64() uint64 {
	return u64.inner
}

func (u64 U64) String() string {
	return fmt.Sprintf("%v", u64.inner)
}

func (u64 U64) Add(other U64) U64 {
	ret := U64{inner: u64.inner + other.inner}
	if ret.inner < u64.inner {
		panic("overflow")
	}
	return ret
}

func (u64 U64) Sub(other U64) U64 {
	ret := U64{inner: u64.inner - other.inner}
	if ret.inner > u64.inner {
		panic("underflow")
	}
	return ret
}

func (u64 *U64) AddAssign(other U64) {
	u64.inner = u64.inner + other.inner
	if u64.inner < other.inner {
		panic("overflow")
	}
}

func (u64 *U64) SubAssign(other U64) {
	u64.inner = u64.inner - other.inner
	if u64.inner > other.inner {
		panic("underflow")
	}
}

func (u64 U64) Equal(other U64) bool {
	return u64.inner == other.inner
}

type U128 struct {
	inner *big.Int
}

func NewU128(x [16]byte) U128 {
	return U128{inner: new(big.Int).SetBytes(x[:])}
}

func (u128 U128) Marshal(w io.Writer) (uint32, error) {
	return surge.Marshal(paddedTo16(u128.inner), w)
}

func (u128 *U128) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	if m < 16 {
		return 0, surge.ErrMaxBytesExceeded
	}
	b := [16]byte{}
	n, err := surge.Unmarshal(&b, r, m)
	if err != nil {
		return n, err
	}
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	u128.inner.SetBytes(b[:])
	return n, nil
}

func (u128 U128) SizeHint() uint32 {
	return 16
}

func (u128 U128) MarshalJSON() ([]byte, error) {
	return json.Marshal(u128.String())
}

func (u128 *U128) UnmarshalJSON(data []byte) error {
	var xString string
	if err := json.Unmarshal(data, &xString); err != nil {
		return err
	}
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	_, ok := u128.inner.SetString(xString, 10)
	if !ok {
		return fmt.Errorf("bad 128-bit unsigned integer=%v", xString)
	}
	if u128.inner.Cmp(MaxU128.inner) > 0 {
		return fmt.Errorf("overflow 256-bit unsigned integer=%v", xString)
	}
	return nil
}

func (U128) Type() Type {
	return TypeU128
}

func (u128 U128) Int() *big.Int {
	return new(big.Int).Set(u128.inner)
}

func (u128 U128) String() string {
	return u128.inner.Text(10)
}

func (u128 U128) Add(other U128) U128 {
	ret := U128{}
	ret.inner.Add(u128.inner, other.inner)
	if ret.inner.Cmp(MaxU128.inner) >= 0 {
		panic("overflow")
	}
	return ret
}

func (u128 U128) Sub(other U128) U128 {
	ret := U128{}
	ret.inner.Sub(u128.inner, other.inner)
	if ret.inner.Sign() == -1 {
		panic("underflow")
	}
	return ret
}

func (u128 *U128) AddAssign(other U128) {
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	u128.inner.Add(u128.inner, other.inner)
	if u128.inner.Cmp(MaxU128.inner) >= 0 {
		panic("overflow")
	}
}

func (u128 *U128) SubAssign(other U128) {
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	u128.inner.Sub(u128.inner, other.inner)
	if u128.inner.Sign() == -1 {
		panic("underflow")
	}
}

func (u128 U128) Equal(other U128) bool {
	if u128.inner == nil {
		return other.inner == nil || other.inner.Sign() == 0
	}
	if other.inner == nil {
		return u128.inner == nil || u128.inner.Sign() == 0
	}
	return u128.inner.Cmp(other.inner) == 0
}

type U256 struct {
	inner *big.Int
}

func NewU256(x [32]byte) U256 {
	return U256{inner: new(big.Int).SetBytes(x[:])}
}

func (u256 U256) Marshal(w io.Writer) (uint32, error) {
	return surge.Marshal(paddedTo32(u256.inner), w)
}

func (u256 *U256) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	if m < 32 {
		return 0, surge.ErrMaxBytesExceeded
	}
	b := [32]byte{}
	n, err := surge.Unmarshal(&b, r, m)
	if err != nil {
		return n, err
	}
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	u256.inner.SetBytes(b[:])
	return n, nil
}

func (u256 U256) SizeHint() uint32 {
	return 32
}

func (u256 U256) MarshalJSON() ([]byte, error) {
	return json.Marshal(u256.String())
}

func (u256 *U256) UnmarshalJSON(data []byte) error {
	var xString string
	if err := json.Unmarshal(data, &xString); err != nil {
		return err
	}
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	_, ok := u256.inner.SetString(xString, 10)
	if !ok {
		return fmt.Errorf("bad 256-bit unsigned integer=%v", xString)
	}
	if u256.inner.Cmp(MaxU256.inner) > 0 {
		return fmt.Errorf("overflow 256-bit unsigned integer=%v", xString)
	}
	return nil
}

func (U256) Type() Type {
	return TypeU256
}

func (u256 U256) Int() *big.Int {
	return new(big.Int).Set(u256.inner)
}

func (u256 U256) String() string {
	return u256.inner.Text(10)
}

func (u256 U256) Add(other U256) U256 {
	ret := U256{}
	ret.inner.Add(u256.inner, other.inner)
	if ret.inner.Cmp(MaxU256.inner) >= 0 {
		panic("overflow")
	}
	return ret
}

func (u256 U256) Sub(other U256) U256 {
	ret := U256{}
	ret.inner.Sub(u256.inner, other.inner)
	if ret.inner.Sign() == -1 {
		panic("underflow")
	}
	return ret
}

func (u256 *U256) AddAssign(other U256) {
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	u256.inner.Add(u256.inner, other.inner)
	if u256.inner.Cmp(MaxU256.inner) >= 0 {
		panic("overflow")
	}
}

func (u256 *U256) SubAssign(other U256) {
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	u256.inner.Sub(u256.inner, other.inner)
	if u256.inner.Sign() == -1 {
		panic("underflow")
	}
}

func (u256 U256) Equal(other U256) bool {
	if u256.inner == nil {
		return other.inner == nil || other.inner.Sign() == 0
	}
	if other.inner == nil {
		return u256.inner == nil || u256.inner.Sign() == 0
	}
	return u256.inner.Cmp(other.inner) == 0
}

// paddedTo16 encodes a big integer as a big-endian into a 16-byte array. It
// will panic if the big integer is more than 16 bytes.
// Modified from:
// https://github.com/ethereum/go-ethereum/blob/master/common/math/big.go
// 17381ecc6695ea9c2d8e5ee0aee5cf70d59a301a
func paddedTo16(bigint *big.Int) [16]byte {
	if bigint.BitLen()/8 > 16 {
		panic(fmt.Sprintf("too big: expected n<16, got n=%v", bigint.BitLen()/8))
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
	if bigint.BitLen()/8 > 32 {
		panic(fmt.Sprintf("too big: expected n<32, got n=%v", bigint.BitLen()/8))
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

var (
	MaxU8 = func() U8 {
		return U8{inner: 255}
	}
	MaxU16 = func() U16 {
		return U16{inner: 65535}
	}
	MaxU32 = func() U32 {
		return U32{inner: 4294967295}
	}
	MaxU64 = func() U64 {
		return U64{inner: 18446744073709551615}
	}
	MaxU128 = func() U128 {
		x, _ := new(big.Int).SetString("340282366920938463463374607431768211455", 10)
		return U128{inner: x}
	}()
	MaxU256 = func() U256 {
		x, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
		return U256{inner: x}
	}()
)
