package extrenvm

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	"github.com/renproject/abi"
	"github.com/renproject/abi/ext"
	"github.com/renproject/surge"
)

// Tx represents a RenVM transaction. It may, or may not, include transactions
// from other blockchains (if it does, then they are usually included as part of
// the RenVM transaction inputs).
type Tx struct {
	// Hash of the to address and the inputs arguments.
	Hash abi.Bytes32 `json:"hash"`

	// To address specifies the RenVM address that will be receiving the
	// transaction. For cross-chain transactions, this will be the unique
	// human-readable address that identifies the specific cross-chain gateway
	// being used.
	To abi.String `json:"to"`

	// In argumenst are provided by the transaction sender. These arguments may,
	// or may not, need to be validated by the receiver. Either way, it is
	// impossible for the receiver to generate these arguments.
	In Arguments `json:"in"`

	// Autogen arguments are can be generated from the input arguments. If these
	// arguments were provided as input, then they would need to be validated by
	// the receiver (which usually involves autogenerating the arguments from
	// scratch, and comparing them against the originally provided ones).
	Autogen Arguments `json:"autogen"`

	// Out arguments are produced by executing the transaction.
	Out Arguments `json:"out"`
}

// NewTxHash returns the Sum256 hash of the To address and In arguments for a
// RenVM transaction. Autogen arguments are ignored because they are
// deterministically generated from the In arguments, and the Out arguments are
// ignored because they do not exist until the transaction has been fully
// executed.
func NewTxHash(to abi.String, in Arguments) abi.Bytes32 {
	buf := new(bytes.Buffer)
	buf.Grow(to.SizeHint() + in.SizeHint())
	if _, err := to.Marshal(buf, abi.MaxBytes); err != nil {
		panic(fmt.Sprintf("error computing txhash: %v", err))
	}
	if _, err := in.Marshal(buf, abi.MaxBytes); err != nil {
		panic(fmt.Sprintf("error computing txhash: %v", err))
	}
	return abi.Bytes32(sha256.Sum256(buf.Bytes()))
}

// NewTx computes the transaction hash and returns a filled in RenVM
// transaction.
func NewTx(to abi.String, in, autogen, out Arguments) Tx {
	return Tx{
		Hash:    NewTxHash(to, in),
		To:      to,
		In:      in,
		Autogen: autogen,
		Out:     out,
	}
}

// SizeHint returns the number of bytes required to represent this transaction
// in binary.
func (tx Tx) SizeHint() int {
	return tx.Hash.SizeHint() +
		tx.To.SizeHint() +
		tx.In.SizeHint() +
		tx.Autogen.SizeHint() +
		tx.Out.SizeHint()
}

// Marshal the transaction into binary. A maximum number of bytes can be
// allocated while marshaling (this maximum is not strict).
func (tx Tx) Marshal(w io.Writer, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	m, err := tx.Hash.Marshal(w, m)
	if err != nil {
		return m, err
	}
	if m, err = tx.To.Marshal(w, m); err != nil {
		return m, err
	}
	if m, err = tx.In.Marshal(w, m); err != nil {
		return m, err
	}
	if m, err = tx.Autogen.Marshal(w, m); err != nil {
		return m, err
	}
	return tx.Out.Marshal(w, m)
}

// Unmarshal the transaction from binary. A maximum number of bytes can be
// allocated while unmarshaling (to protect against malicious input).
func (tx *Tx) Unmarshal(r io.Reader, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	m, err := tx.Hash.Unmarshal(r, m)
	if err != nil {
		return m, err
	}
	if m, err = tx.To.Unmarshal(r, m); err != nil {
		return m, err
	}
	if m, err = tx.In.Unmarshal(r, m); err != nil {
		return m, err
	}
	if m, err = tx.Autogen.Unmarshal(r, m); err != nil {
		return m, err
	}
	return tx.Out.Unmarshal(r, m)
}

func (Tx) Type() abi.Type {
	return ext.TypeRenVMTx
}

type Argument struct {
	Name  abi.String
	Value abi.Value
}

func (Argument) Type() abi.Type {
	return ext.TypeRenVMArgument
}

func (arg Argument) SizeHint() int {
	return arg.Name.SizeHint() +
		arg.Value.Type().SizeHint() +
		arg.Value.SizeHint()
}

func (arg Argument) Marshal(w io.Writer, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	m, err := arg.Name.Marshal(w, m)
	if err != nil {
		return m, err
	}
	return abi.Marshal(w, arg.Value, m)
}

func (arg *Argument) Unmarshal(r io.Reader, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	m, err := arg.Name.Unmarshal(r, m)
	if err != nil {
		return m, err
	}
	arg.Value, m, err = abi.Unmarshal(r, m)
	return m, err
}

func (arg Argument) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":  arg.Name,
		"value": []json.Marshaler{arg.Value.Type(), arg.Value},
	}
	return json.Marshal(m)
}

func (arg *Argument) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	// Unmarshal arg name.
	var name abi.String
	if err := name.UnmarshalJSON(m["name"]); err != nil {
		return err
	}

	// Unmarshal arg value.
	value, err := abi.UnmarshalJSON(m["value"])
	if err != nil {
		return err
	}

	arg.Name = name
	arg.Value = value
	return nil
}

type Arguments []Argument

func (Arguments) Type() abi.Type {
	return ext.TypeRenVMArguments
}

func (args Arguments) SizeHint() int {
	size := 4
	for _, arg := range args {
		size += arg.SizeHint()
	}
	return size
}

func (args Arguments) Marshal(w io.Writer, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	// Marshal length.
	m, err := surge.Marshal(w, uint32(len(args)), m)
	if err != nil {
		return m, err
	}
	// Marshal arguments.
	for _, arg := range args {
		m, err = abi.Marshal(w, arg, m)
		if err != nil {
			return m, err
		}
	}
	return m, nil
}

func (args *Arguments) Unmarshal(r io.Reader, m int) (int, error) {
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}

	// Unmarshal length.
	len := uint32(0)
	m, err := surge.Unmarshal(r, &len, m)
	if err != nil {
		return m, err
	}
	// Check length.
	if int(len) < 0 {
		return m, fmt.Errorf("unmarshal error: len>=0, got len=%v", int(len))
	}
	m -= int(len)
	if m <= 0 {
		return m, surge.ErrMaxBytesExceeded
	}
	// Unmarshal args.
	*args = make([]Argument, 0, len)
	for i := 0; i < int(len); i++ {
		var arg Argument
		m, err = arg.Unmarshal(r, m)
		if err != nil {
			return m, err
		}
		*args = append(*args, arg)
	}
	return m, nil
}
