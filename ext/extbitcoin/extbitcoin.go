package extbitcoin

import (
	"github.com/renproject/abi"
)

// An Address represents a Bitcoin address (with SegWit support).
type Address abi.String

// A UTXOIndex uniquely identifies an unspent transaction output, and can be
// used to find the complete UTXO information on the Bitcoin blockchain.
type UTXOIndex struct {
	TxHash abi.Bytes32 `json:"tx"`
	VOut   abi.U32     `json:"vout"`
}

// A UTXO is the complete information of an unspent transaction output. It
// includes the UTXOIndex.
type UTXO struct {
	UTXOIndex
	Amount       abi.U64   `json:"amount"`
	ScriptPubKey abi.Bytes `json:"scriptPubKey"`
}
