package extzcash

import (
	"github.com/renproject/abi"
	"github.com/renproject/abi/ext/extbitcoin"
)

// An Address represents a ZCash address.
type Address abi.String

// A UTXOIndex uniquely identifies an unspent transaction output, and can be
// used to find the complete UTXO information on the ZCash blockchain.
type UTXOIndex = extbitcoin.UTXOIndex

// A UTXO is the complete information of an unspent transaction output. It
// includes the UTXOIndex.
type UTXO = extbitcoin.UTXO
