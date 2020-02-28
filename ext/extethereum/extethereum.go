package extethereum

import (
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

func (addr Address) Marshal(w io.Writer) (uint32, error) {
	n, err := w.Write(addr[:])
	return uint32(n), err
}

func (addr *Address) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	if m < ethcommon.AddressLength {
		return 0, surge.ErrMaxBytesExceeded
	}
	b := [ethcommon.AddressLength]byte{}
	n, err := r.Read(b[:])
	if err != nil {
		return uint32(n), err
	}
	*addr = Address(ethcommon.Address(b))
	return uint32(n), nil
}

func (Address) SizeHint() uint32 {
	return uint32(ethcommon.AddressLength)
}

func (Address) Type() abi.Type {
	return ext.TypeEthereumAddress
}

type Tx struct {
	Hash abi.Bytes32 `json:"txhash"`
}

func NewTx(hash abi.Bytes32) Tx {
	return Tx{Hash: hash}
}

func (tx Tx) Marshal(w io.Writer) (uint32, error) {
	return tx.Hash.Marshal(w)
}

func (tx *Tx) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	return tx.Hash.Unmarshal(r, m)
}

func (tx Tx) SizeHint() uint32 {
	return tx.Hash.SizeHint()
}

func (Tx) Type() abi.Type {
	return ext.TypeEthereumTx
}

// EncodeArguments into an Ethereum ABI compatible byte slice.
func EncodeArguments(ls abi.List) []byte {
	ethargs := make(ethabi.Arguments, 0, ls.Len())
	ethvals := make([]interface{}, 0, ls.Len())

	ls.ForEach(func(_ int, elem abi.Value) {
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
	})

	packed, err := ethargs.Pack(ethvals...)
	if err != nil {
		panic(fmt.Errorf("error packing: %v", err))
	}
	return packed
}
