package ext

import (
	"io"

	"github.com/renproject/abi"
	"github.com/renproject/surge"
)

// An Address represents a Bitcoin address (with SegWit support).
type Address abi.String

func (Address) Type() abi.Type {
	return TypeBitcoinAddress
}

// A UTXOIndex uniquely identifies an unspent transaction output, and can be
// used to find the complete UTXO information on the Bitcoin blockchain.
type UTXOIndex struct {
	TxHash abi.Bytes32 `json:"txHash"`
	VOut   abi.U32     `json:"vOut"`
}

func (UTXOIndex) Type() abi.Type {
	return TypeBitcoinUTXOIndex
}

func (utxoi UTXOIndex) SizeHint() int {
	return utxoi.TxHash.SizeHint() + utxoi.VOut.SizeHint()
}

func (utxoi UTXOIndex) Marshal(w io.Writer, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	m, err := utxoi.TxHash.Marshal(w, m)
	if err != nil {
		return m, err
	}
	return utxoi.VOut.Marshal(w, m)
}

func (utxoi *UTXOIndex) Unmarshal(r io.Reader, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

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

// A UTXO is the complete information of an unspent transaction output. It
// includes the UTXOIndex.
type UTXO struct {
	UTXOIndex
	Amount       abi.U64   `json:"amount"`
	ScriptPubKey abi.Bytes `json:"scriptPubKey"`
}

func (UTXO) Type() abi.Type {
	return TypeBitcoinUTXO
}

func (utxo UTXO) SizeHint() int {
	return utxo.TxHash.SizeHint() + utxo.VOut.SizeHint()
}

func (utxo UTXO) Marshal(w io.Writer, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	m, err := utxo.UTXOIndex.Marshal(w, m)
	if err != nil {
		return m, err
	}
	if m, err = utxo.Amount.Marshal(w, m); err != nil {
		return m, err
	}
	return utxo.ScriptPubKey.Marshal(w, m)
}

func (utxo *UTXO) Unmarshal(r io.Reader, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	m, err := utxo.UTXOIndex.Unmarshal(r, m)
	if err != nil {
		return m, err
	}
	if m, err = utxo.Amount.Unmarshal(r, m); err != nil {
		return m, err
	}
	return utxo.ScriptPubKey.Unmarshal(r, m)
}
