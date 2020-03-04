package extethereum

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"

	"github.com/renproject/abi"
	"github.com/renproject/abi/ext"
	"github.com/renproject/surge"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

type Address ethcommon.Address

func (Address) Type() abi.Type {
	return ext.TypeEthereumAddress
}

func (Address) SizeHint() int {
	return ethcommon.AddressLength
}

func (addr Address) Marshal(w io.Writer, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	n, err := w.Write(addr[:])
	return m - n, err
}

func (addr *Address) Unmarshal(r io.Reader, m int) (int, error) {
	if m <= ethcommon.AddressLength {
		return m, surge.ErrMaxBytesExceeded
	}

	b := [ethcommon.AddressLength]byte{}
	n, err := r.Read(b[:])
	if err != nil {
		return m - n, err
	}
	*addr = Address(ethcommon.Address(b))
	return m - n, nil
}

func (addr Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(ethcommon.Address(addr).Hex())
}

func (addr *Address) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*addr = Address(ethcommon.HexToAddress(str))
	return nil
}

type Tx struct {
	Hash abi.Bytes32 `json:"txhash"`
}

func NewTx(hash abi.Bytes32) Tx {
	return Tx{Hash: hash}
}

func (Tx) Type() abi.Type {
	return ext.TypeEthereumTx
}

func (tx Tx) SizeHint() int {
	return tx.Hash.SizeHint()
}

func (tx Tx) Marshal(w io.Writer, m int) (int, error) {
	return tx.Hash.Marshal(w, m)
}

func (tx *Tx) Unmarshal(r io.Reader, m int) (int, error) {
	return tx.Hash.Unmarshal(r, m)
}

// EncodeArguments into an Ethereum ABI compatible byte slice.
func EncodeArguments(ls []abi.Value) []byte {
	ethargs := make(ethabi.Arguments, 0, len(ls))
	ethvals := make([]interface{}, 0, len(ls))

	for _, elem := range ls {
		var val interface{}
		var ty ethabi.Type
		var err error

		switch elem.Type() {
		case abi.TypeString:
			val = elem.(abi.String)
			ty, err = ethabi.NewType("string", nil)
		case abi.TypeBytes:
			val = elem.(abi.Bytes)
			ty, err = ethabi.NewType("bytes", nil)
		case abi.TypeBytes32:
			val = elem.(abi.Bytes32)
			ty, err = ethabi.NewType("bytes32", nil)

		case abi.TypeU8:
			val = big.NewInt(0).SetUint64(uint64(elem.(abi.U8).Uint8()))
			ty, err = ethabi.NewType("uint256", nil)
		case abi.TypeU16:
			val = big.NewInt(0).SetUint64(uint64(elem.(abi.U16).Uint16()))
			ty, err = ethabi.NewType("uint256", nil)
		case abi.TypeU32:
			val = big.NewInt(0).SetUint64(uint64(elem.(abi.U32).Uint32()))
			ty, err = ethabi.NewType("uint256", nil)
		case abi.TypeU64:
			val = big.NewInt(0).SetUint64(uint64(elem.(abi.U64).Uint64()))
			ty, err = ethabi.NewType("uint256", nil)
		case abi.TypeU128:
			val = elem.(abi.U128).Int()
			ty, err = ethabi.NewType("uint256", nil)
		case abi.TypeU256:
			val = elem.(abi.U256).Int()
			ty, err = ethabi.NewType("uint256", nil)

		case ext.TypeEthereumAddress:
			val = elem.(Address)
			ty, err = ethabi.NewType("address", nil)

		default:
			panic(fmt.Errorf("unexpected type=%v", elem.Type()))
		}

		if err != nil {
			panic(fmt.Errorf("error encoding type: %v", err))
		}
		ethargs = append(ethargs, ethabi.Argument{
			Type: ty,
		})
		ethvals = append(ethvals, val)
	}

	packed, err := ethargs.Pack(ethvals...)
	if err != nil {
		panic(fmt.Errorf("error packing: %v", err))
	}
	return packed
}
