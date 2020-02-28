package extbitcoin

import (
	"io"

	"github.com/renproject/abi"
	"github.com/renproject/abi/ext"
	"github.com/renproject/surge"
)

// An Address represents a Bitcoin address (with SegWit support).
type Address abi.String

func (Address) Type() abi.Type {
	return ext.TypeBitcoinAddress
}

// A UTXOIndex uniquely identifies an unspent transaction output, and can be
// used to find the complete UTXO information on the Bitcoin blockchain.
type UTXOIndex struct {
	TxHash abi.Bytes32 `json:"txHash"`
	VOut   abi.U32     `json:"vOut"`
}

func (utxoi UTXOIndex) Marshal(w io.Writer) (uint32, error) {
	n1, err := utxoi.TxHash.Marshal(w)
	if err != nil {
		return n1, err
	}
	n2, err := utxoi.VOut.Marshal(w)
	if err != nil {
		return n1 + n2, err
	}
	return n1 + n2, nil
}

func (utxoi *UTXOIndex) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	n1, err := utxoi.TxHash.Unmarshal(r, m)
	m -= n1
	if err != nil {
		return n1, err
	}
	if m < 0 {
		return n1, surge.ErrMaxBytesExceeded
	}

	n2, err := utxoi.VOut.Unmarshal(r, m)
	m -= n2
	if err != nil {
		return n2, err
	}
	if m < 0 {
		return n1 + n2, surge.ErrMaxBytesExceeded
	}

	return n1 + n2, nil
}

func (utxoi UTXOIndex) SizeHint() uint32 {
	return utxoi.TxHash.SizeHint() + utxoi.VOut.SizeHint()
}

func (UTXOIndex) Type() abi.Type {
	return ext.TypeBitcoinUTXOIndex
}

// A UTXO is the complete information of an unspent transaction output. It
// includes the UTXOIndex.
type UTXO struct {
	UTXOIndex
	Amount       abi.U64   `json:"amount"`
	ScriptPubKey abi.Bytes `json:"scriptPubKey"`
}

func (utxo UTXO) Marshal(w io.Writer) (uint32, error) {
	n1, err := utxo.UTXOIndex.Marshal(w)
	if err != nil {
		return n1, err
	}
	n2, err := utxo.Amount.Marshal(w)
	if err != nil {
		return n1 + n2, err
	}
	n3, err := utxo.ScriptPubKey.Marshal(w)
	if err != nil {
		return n1 + n2 + n3, err
	}
	return n1 + n2 + n3, nil
}

func (utxo *UTXO) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	n1, err := utxo.UTXOIndex.Unmarshal(r, m)
	m -= n1
	if err != nil {
		return n1, err
	}
	if m < 0 {
		return n1, surge.ErrMaxBytesExceeded
	}

	n2, err := utxo.Amount.Unmarshal(r, m)
	m -= n2
	if err != nil {
		return n1 + n2, err
	}
	if m < 0 {
		return n1 + n2, surge.ErrMaxBytesExceeded
	}

	n3, err := utxo.ScriptPubKey.Unmarshal(r, m)
	m -= n2
	if err != nil {
		return n1 + n2 + n3, err
	}
	if m < 0 {
		return n1 + n2 + n3, surge.ErrMaxBytesExceeded
	}

	return n1 + n2 + n3, nil
}

func (utxo UTXO) SizeHint() uint32 {
	return utxo.TxHash.SizeHint() + utxo.VOut.SizeHint()
}

func (UTXO) Type() abi.Type {
	return ext.TypeBitcoinUTXO
}
