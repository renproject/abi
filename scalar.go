package abi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strconv"

	"github.com/renproject/surge"
)

// Bool represents a boolean value.
type Bool struct {
	inner bool
}

// NewBool returns a bool wrapped as a Bool.
func NewBool(x bool) Bool {
	return Bool{inner: x}
}

// Bool returns the inner bool.
func (b Bool) Bool() bool {
	return b.inner
}

// Equal compares one Bool to another. If they are equal, then it returns true.
// Otherwise, it returns false.
func (b Bool) Equal(other Bool) bool {
	return b.inner == other.inner
}

// Type returns the type identifier.
func (Bool) Type() Type {
	return TypeBool
}

// SizeHint returns the number of bytes required to represent a Bool in binary.
func (b Bool) SizeHint() int {
	return 1
}

// Marshal the Bool to binary. Marshaling will try to avoid allocating more than
// the specified maximum number of bytes. If it needs to allocate too many
// bytes, and error may be returned instead.
func (b Bool) Marshal(w io.Writer, m int) (int, error) {
	return surge.Marshal(w, b.inner, m)
}

// Unmarshal the Bool from binary. Unmarshaling will not allocate more than the
// specified maximum number of bytes. If it needs to allocate too many bytes,
// and error is returned instead.
func (b *Bool) Unmarshal(r io.Reader, m int) (int, error) {
	return surge.Unmarshal(r, &b.inner, m)
}

// MarshalJSON implements the JSON marshaler interface. Bools are marshaled as
// decimal strings (for consistency with larger integer types).
func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.inner)
}

// UnmarshalJSON implements the JSON unmarshaler interface. Bools are unmarshaled
// as decimal strings (for consistency with larger integer types).
func (b *Bool) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &b.inner)
}

func (b Bool) String() string {
	if b.inner {
		return "true"
	}
	return "false"
}

// U8 represents an 8-bit unsigned integer.
type U8 struct {
	inner uint8
}

// NewU8 returns a uint8 wrapped as a U8.
func NewU8(x uint8) U8 {
	return U8{inner: x}
}

// Uint8 returns the inner uint8.
func (u8 U8) Uint8() uint8 {
	return u8.inner
}

// Add one U8 to another and return the result.
func (u8 U8) Add(other U8) U8 {
	ret := U8{inner: u8.inner + other.inner}
	if ret.inner < u8.inner {
		panic("overflow")
	}
	return ret
}

// Sub one U8 from another and return the result.
func (u8 U8) Sub(other U8) U8 {
	ret := U8{inner: u8.inner - other.inner}
	if ret.inner > u8.inner {
		panic("underflow")
	}
	return ret
}

// AddAssign will add one U8 to another and assign the result to the left-hand
// side.
func (u8 *U8) AddAssign(other U8) {
	u8.inner = u8.inner + other.inner
	if u8.inner < other.inner {
		panic("overflow")
	}
}

// SubAssign will sub one U8 from another and assign the result to the left-hand
// side.
func (u8 *U8) SubAssign(other U8) {
	u8.inner = u8.inner - other.inner
	if u8.inner > other.inner {
		panic("underflow")
	}
}

// Equal compares one U8 to another. If they are equal, then it returns true.
// Otherwise, it returns false.
func (u8 U8) Equal(other U8) bool {
	return u8.inner == other.inner
}

// Type returns the type identifier.
func (U8) Type() Type {
	return TypeU8
}

// SizeHint returns the number of bytes required to represent a U8 in binary.
func (u8 U8) SizeHint() int {
	return 1
}

// Marshal the U8 to binary. Marshaling will try to avoid allocating more than
// the specified maximum number of bytes. If it needs to allocate too many
// bytes, and error may be returned instead.
func (u8 U8) Marshal(w io.Writer, m int) (int, error) {
	return surge.Marshal(w, u8.inner, m)
}

// Unmarshal the U8 from binary. Unmarshaling will not allocate more than the
// specified maximum number of bytes. If it needs to allocate too many bytes,
// and error is returned instead.
func (u8 *U8) Unmarshal(r io.Reader, m int) (int, error) {
	return surge.Unmarshal(r, &u8.inner, m)
}

// MarshalJSON implements the JSON marshaler interface. U8s are marshaled as
// decimal strings (for consistency with larger integer types).
func (u8 U8) MarshalJSON() ([]byte, error) {
	return json.Marshal(u8.String())
}

// UnmarshalJSON implements the JSON unmarshaler interface. U8s are unmarshaled
// as decimal strings (for consistency with larger integer types).
func (u8 *U8) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	x, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return err
	}
	u8.inner = uint8(x)
	return nil
}

func (u8 U8) String() string {
	return fmt.Sprintf("%v", u8.inner)
}

// U16 represents a 16-bit unsigned integer.
type U16 struct {
	inner uint16
}

// NewU16 returns a uint16 wrapped as a U16.
func NewU16(x uint16) U16 {
	return U16{inner: x}
}

// NewU16FromU8 returns a uint16 wrapped as a U16.
func NewU16FromU8(x U8) U16 {
	return U16{inner: uint16(x.Uint8())}
}

// Uint16 returns the inner uint16.
func (u16 U16) Uint16() uint16 {
	return u16.inner
}

// Add one U16 to another and return the result.
func (u16 U16) Add(other U16) U16 {
	ret := U16{inner: u16.inner + other.inner}
	if ret.inner < u16.inner {
		panic("overflow")
	}
	return ret
}

// Sub one U16 from another and return the result.
func (u16 U16) Sub(other U16) U16 {
	ret := U16{inner: u16.inner - other.inner}
	if ret.inner > u16.inner {
		panic("underflow")
	}
	return ret
}

// AddAssign will add one U16 to another and assign the result to the left-hand
// side.
func (u16 *U16) AddAssign(other U16) {
	u16.inner = u16.inner + other.inner
	if u16.inner < other.inner {
		panic("overflow")
	}
}

// SubAssign will sub one U16 from another and assign the result to the left-hand
// side.
func (u16 *U16) SubAssign(other U16) {
	u16.inner = u16.inner - other.inner
	if u16.inner > other.inner {
		panic("underflow")
	}
}

// Equal compares one U16 to another. If they are equal, then it returns true.
// Otherwise, it returns false.
func (u16 U16) Equal(other U16) bool {
	return u16.inner == other.inner
}

// Type returns the type identifier.
func (U16) Type() Type {
	return TypeU16
}

// SizeHint returns the number of bytes required to represent a U16 in binary.
func (u16 U16) SizeHint() int {
	return 2
}

// Marshal the U16 to binary. Marshaling will try to avoid allocating more than
// the specified maximum number of bytes. If it needs to allocate too many
// bytes, and error may be returned instead.
func (u16 U16) Marshal(w io.Writer, m int) (int, error) {
	return surge.Marshal(w, u16.inner, m)
}

// Unmarshal the U16 from binary. Unmarshaling will not allocate more than the
// specified maximum number of bytes. If it needs to allocate too many bytes,
// and error is returned instead.
func (u16 *U16) Unmarshal(r io.Reader, m int) (int, error) {
	return surge.Unmarshal(r, &u16.inner, m)
}

// MarshalJSON implements the JSON marshaler interface. U16s are marshaled as
// decimal strings (for consistency with larger integer types).
func (u16 U16) MarshalJSON() ([]byte, error) {
	return json.Marshal(u16.String())
}

// UnmarshalJSON implements the JSON unmarshaler interface. U16s are unmarshaled
// as decimal strings (for consistency with larger integer types).
func (u16 *U16) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	x, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return err
	}
	u16.inner = uint16(x)
	return nil
}

func (u16 U16) String() string {
	return fmt.Sprintf("%v", u16.inner)
}

// U32 represents a 32-bit unsigned integer.
type U32 struct {
	inner uint32
}

// NewU32 returns a uint32 wrapped as a U32.
func NewU32(x uint32) U32 {
	return U32{inner: x}
}

// NewU32FromU8 returns a uint32 wrapped as a U32.
func NewU32FromU8(x U8) U32 {
	return U32{inner: uint32(x.Uint8())}
}

// NewU32FromU16 returns a uint32 wrapped as a U32.
func NewU32FromU16(x U16) U32 {
	return U32{inner: uint32(x.Uint16())}
}

// Uint32 returns the inner uint32.
func (u32 U32) Uint32() uint32 {
	return u32.inner
}

// Add one U32 to another and return the result.
func (u32 U32) Add(other U32) U32 {
	ret := U32{inner: u32.inner + other.inner}
	if ret.inner < u32.inner {
		panic("overflow")
	}
	return ret
}

// Sub one U32 from another and return the result.
func (u32 U32) Sub(other U32) U32 {
	ret := U32{inner: u32.inner - other.inner}
	if ret.inner > u32.inner {
		panic("underflow")
	}
	return ret
}

// AddAssign will add one U32 to another and assign the result to the left-hand
// side.
func (u32 *U32) AddAssign(other U32) {
	u32.inner = u32.inner + other.inner
	if u32.inner < other.inner {
		panic("overflow")
	}
}

// SubAssign will sub one U32 from another and assign the result to the left-hand
// side.
func (u32 *U32) SubAssign(other U32) {
	u32.inner = u32.inner - other.inner
	if u32.inner > other.inner {
		panic("underflow")
	}
}

// Equal compares one U32 to another. If they are equal, then it returns true.
// Otherwise, it returns false.
func (u32 U32) Equal(other U32) bool {
	return u32.inner == other.inner
}

// Type returns the type identifier.
func (U32) Type() Type {
	return TypeU32
}

// SizeHint returns the number of bytes required to represent a U32 in binary.
func (u32 U32) SizeHint() int {
	return 4
}

// Marshal the U32 to binary. Marshaling will try to avoid allocating more than
// the specified maximum number of bytes. If it needs to allocate too many
// bytes, and error may be returned instead.
func (u32 U32) Marshal(w io.Writer, m int) (int, error) {
	return surge.Marshal(w, u32.inner, m)
}

// Unmarshal the U32 from binary. Unmarshaling will not allocate more than the
// specified maximum number of bytes. If it needs to allocate too many bytes,
// and error is returned instead.
func (u32 *U32) Unmarshal(r io.Reader, m int) (int, error) {
	return surge.Unmarshal(r, &u32.inner, m)
}

// MarshalJSON implements the JSON marshaler interface. U32s are marshaled as
// decimal strings (for consistency with larger integer types).
func (u32 U32) MarshalJSON() ([]byte, error) {
	return json.Marshal(u32.String())
}

// UnmarshalJSON implements the JSON unmarshaler interface. U32s are unmarshaled
// as decimal strings (for consistency with larger integer types).
func (u32 *U32) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	x, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return err
	}
	u32.inner = uint32(x)
	return nil
}

func (u32 U32) String() string {
	return fmt.Sprintf("%v", u32.inner)
}

// U64 represents a 64-bit unsigned integer.
type U64 struct {
	inner uint64
}

// NewU64 returns a uint64 wrapped as a U64.
func NewU64(x uint64) U64 {
	return U64{inner: x}
}

// NewU64FromU8 returns a uint64 wrapped as a U64.
func NewU64FromU8(x U8) U64 {
	return U64{inner: uint64(x.Uint8())}
}

// NewU64FromU16 returns a uint64 wrapped as a U64.
func NewU64FromU16(x U16) U64 {
	return U64{inner: uint64(x.Uint16())}
}

// NewU64FromU32 returns a uint32 wrapped as a U64.
func NewU64FromU32(x U32) U64 {
	return U64{inner: uint64(x.Uint32())}
}

// Uint64 returns the inner uint64.
func (u64 U64) Uint64() uint64 {
	return u64.inner
}

// Add one U64 to another and return the result.
func (u64 U64) Add(other U64) U64 {
	ret := U64{inner: u64.inner + other.inner}
	if ret.inner < u64.inner {
		panic("overflow")
	}
	return ret
}

// Sub one U64 from another and return the result.
func (u64 U64) Sub(other U64) U64 {
	ret := U64{inner: u64.inner - other.inner}
	if ret.inner > u64.inner {
		panic("underflow")
	}
	return ret
}

// AddAssign will add one U64 to another and assign the result to the left-hand
// side.
func (u64 *U64) AddAssign(other U64) {
	u64.inner = u64.inner + other.inner
	if u64.inner < other.inner {
		panic("overflow")
	}
}

// SubAssign will sub one U64 from another and assign the result to the left-hand
// side.
func (u64 *U64) SubAssign(other U64) {
	u64.inner = u64.inner - other.inner
	if u64.inner > other.inner {
		panic("underflow")
	}
}

// Equal compares one U64 to another. If they are equal, then it returns true.
// Otherwise, it returns false.
func (u64 U64) Equal(other U64) bool {
	return u64.inner == other.inner
}

// Type returns the type identifier.
func (U64) Type() Type {
	return TypeU64
}

// SizeHint returns the number of bytes required to represent a U64 in binary.
func (u64 U64) SizeHint() int {
	return 8
}

// Marshal the U64 to binary. Marshaling will try to avoid allocating more than
// the specified maximum number of bytes. If it needs to allocate too many
// bytes, and error may be returned instead.
func (u64 U64) Marshal(w io.Writer, m int) (int, error) {
	return surge.Marshal(w, u64.inner, m)
}

// Unmarshal the U64 from binary. Unmarshaling will not allocate more than the
// specified maximum number of bytes. If it needs to allocate too many bytes,
// and error is returned instead.
func (u64 *U64) Unmarshal(r io.Reader, m int) (int, error) {
	return surge.Unmarshal(r, &u64.inner, m)
}

// MarshalJSON implements the JSON marshaler interface. U64s are marshaled as
// decimal strings (for consistency with larger integer types).
func (u64 U64) MarshalJSON() ([]byte, error) {
	return json.Marshal(u64.String())
}

// UnmarshalJSON implements the JSON unmarshaler interface. U64s are unmarshaled
// as decimal strings (for consistency with larger integer types).
func (u64 *U64) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	x, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return err
	}
	u64.inner = uint64(x)
	return nil
}

func (u64 U64) String() string {
	return fmt.Sprintf("%v", u64.inner)
}

type U128 struct {
	inner *big.Int
}

func NewU128(x [16]byte) U128 {
	return U128{inner: new(big.Int).SetBytes(x[:])}
}

func NewU128FromU8(x U8) U128 {
	return U128{inner: new(big.Int).SetUint64(uint64(x.Uint8()))}
}

func NewU128FromU16(x U16) U128 {
	return U128{inner: new(big.Int).SetUint64(uint64(x.Uint16()))}
}

func NewU128FromU32(x U32) U128 {
	return U128{inner: new(big.Int).SetUint64(uint64(x.Uint32()))}
}

func NewU128FromU64(x U64) U128 {
	return U128{inner: new(big.Int).SetUint64(x.Uint64())}
}

func NewU128FromInt(x *big.Int) U128 {
	if x.Sign() == -1 {
		panic("underflow")
	}
	if x.Cmp(MaxU128.inner) > 0 {
		panic("overflow")
	}
	return U128{inner: new(big.Int).Set(x)}
}

func (u128 U128) Int() *big.Int {
	return new(big.Int).Set(u128.inner)
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

func (U128) Type() Type {
	return TypeU128
}

func (u128 U128) SizeHint() int {
	return 16
}

func (u128 U128) Marshal(w io.Writer, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	b16 := paddedTo16(u128.inner)
	n, err := w.Write(b16[:])
	return m - n, err
}

func (u128 *U128) Unmarshal(r io.Reader, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	b16 := [16]byte{}
	n, err := io.ReadFull(r, b16[:])
	if err != nil {
		return m, err
	}
	m -= n
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	u128.inner.SetBytes(b16[:])
	return m, err
}

func (u128 U128) MarshalJSON() ([]byte, error) {
	return json.Marshal(u128.String())
}

func (u128 *U128) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if u128.inner == nil {
		u128.inner = new(big.Int)
	}
	_, ok := u128.inner.SetString(str, 10)
	if !ok {
		return fmt.Errorf("malformed: U128(%v)", str)
	}
	if u128.inner.Sign() == -1 {
		return fmt.Errorf("underflow: U128(%v)", str)
	}
	if u128.inner.Cmp(MaxU128.inner) > 0 {
		return fmt.Errorf("overflow: U128(%v)", str)
	}
	return nil
}

func (u128 U128) String() string {
	return u128.inner.Text(10)
}

type U256 struct {
	inner *big.Int
}

func NewU256(x [32]byte) U256 {
	return U256{inner: new(big.Int).SetBytes(x[:])}
}

func NewU256FromU8(x U8) U256 {
	return U256{inner: new(big.Int).SetUint64(uint64(x.Uint8()))}
}

func NewU256FromU16(x U16) U256 {
	return U256{inner: new(big.Int).SetUint64(uint64(x.Uint16()))}
}

func NewU256FromU32(x U32) U256 {
	return U256{inner: new(big.Int).SetUint64(uint64(x.Uint32()))}
}

func NewU256FromU64(x U64) U256 {
	return U256{inner: new(big.Int).SetUint64(x.Uint64())}
}

func NewU256FromU128(x U128) U256 {
	return NewU256FromInt(x.Int())
}

func NewU256FromInt(x *big.Int) U256 {
	if x.Sign() == -1 {
		panic("underflow")
	}
	if x.Cmp(MaxU256.inner) > 0 {
		panic("overflow")
	}
	return U256{inner: new(big.Int).Set(x)}
}

func (u256 U256) Int() *big.Int {
	return new(big.Int).Set(u256.inner)
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

func (U256) Type() Type {
	return TypeU256
}

func (u256 U256) SizeHint() int {
	return 32
}

func (u256 U256) Marshal(w io.Writer, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	b32 := paddedTo32(u256.inner)
	n, err := w.Write(b32[:])
	return m - n, err
}

func (u256 *U256) Unmarshal(r io.Reader, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	b32 := [32]byte{}
	n, err := io.ReadFull(r, b32[:])
	if err != nil {
		return m, err
	}
	m -= n
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	u256.inner.SetBytes(b32[:])
	return m, err
}

func (u256 U256) MarshalJSON() ([]byte, error) {
	return json.Marshal(u256.String())
}

func (u256 *U256) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if u256.inner == nil {
		u256.inner = new(big.Int)
	}
	_, ok := u256.inner.SetString(str, 10)
	if !ok {
		return fmt.Errorf("malformed: U256(%v)", str)
	}
	if u256.inner.Sign() == -1 {
		return fmt.Errorf("underflow: U256(%v)", str)
	}
	if u256.inner.Cmp(MaxU256.inner) > 0 {
		return fmt.Errorf("overflow: U256(%v)", str)
	}
	return nil
}

func (u256 U256) String() string {
	return u256.inner.Text(10)
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
