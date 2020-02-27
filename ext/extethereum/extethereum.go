package extethereum

import (
	"io"

	"github.com/renproject/abi"
)

type Address abi.Bytes

type Tx struct {
	Hash abi.Bytes32 `json:"txhash"`
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
