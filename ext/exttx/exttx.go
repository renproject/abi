package exttx

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"

	"github.com/renproject/abi"
	"github.com/renproject/surge"
)

func ComputeTxHash(to abi.String, in Arguments) abi.Bytes32 {
	buf := new(bytes.Buffer)
	buf.Grow(int(to.SizeHint() + in.SizeHint()))
	if _, err := to.Marshal(buf); err != nil {
		panic(fmt.Sprintf("error computing txhash: %v", err))
	}
	if _, err := in.Marshal(buf); err != nil {
		panic(fmt.Sprintf("error computing txhash: %v", err))
	}
	return abi.Bytes32(sha256.Sum256(buf.Bytes()))
}

type Tx struct {
	Hash    abi.Bytes32 `json:"hash"`
	To      abi.String  `json:"to"`
	In      Arguments   `json:"in"`
	Autogen Arguments   `json:"autogen"`
	Out     Arguments   `json:"out"`
}

func NewTx(to abi.String, in, autogen, out Arguments) Tx {
	tx := Tx{
		Hash:    abi.Bytes32{},
		To:      to,
		In:      in,
		Autogen: autogen,
		Out:     out,
	}
	tx.Hash = ComputeTxHash(to, in)
	return tx
}

func (tx Tx) Marshal(w io.Writer) (uint32, error) {
	n1, err := tx.Hash.Marshal(w)
	if err != nil {
		return n1, err
	}
	n2, err := tx.To.Marshal(w)
	n1 += n2
	if err != nil {
		return n1, err
	}
	n3, err := tx.In.Marshal(w)
	n1 += n3
	if err != nil {
		return n1, err
	}
	n4, err := tx.Autogen.Marshal(w)
	n1 += n4
	if err != nil {
		return n1, err
	}
	n5, err := tx.Out.Marshal(w)
	n1 += n5
	if err != nil {
		return n1, err
	}
	return n1, nil
}

func (tx *Tx) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	n1, err := tx.Hash.Unmarshal(r, m)
	if err != nil {
		return n1, err
	}
	n2, err := tx.To.Unmarshal(r, m)
	n1 += n2
	m -= n2
	if err != nil {
		return n1, err
	}
	n3, err := tx.In.Unmarshal(r, m)
	n1 += n3
	m -= n3
	if err != nil {
		return n1, err
	}
	n4, err := tx.Autogen.Unmarshal(r, m)
	n1 += n4
	m -= n4
	if err != nil {
		return n1, err
	}
	n5, err := tx.Out.Unmarshal(r, m)
	n1 += n5
	m -= n5
	if err != nil {
		return n1, err
	}
	return n1, nil
}

func (tx Tx) SizeHint() uint32 {
	return tx.Hash.SizeHint() +
		tx.To.SizeHint() +
		tx.In.SizeHint() +
		tx.Autogen.SizeHint() +
		tx.Out.SizeHint()
}

type Arguments []Argument

func (args Arguments) Marshal(w io.Writer) (uint32, error) {
	// Marshal length.
	n1, err := surge.Marshal(uint32(len(args)), w)
	if err != nil {
		return n1, err
	}
	// Marshal arguments.
	for _, arg := range args {
		n2, err := surge.Marshal(arg, w)
		n1 += n2
		if err != nil {
			return n1, err
		}
	}
	return n1, nil
}

func (args *Arguments) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	len := uint32(0)
	n1, err := surge.Unmarshal(&len, r, m)
	if err != nil {
		return n1, err
	}
	m -= n1
	if m < len {
		return n1, surge.ErrMaxBytesExceeded
	}
	*args = make([]Argument, 0, len)
	for i := uint32(0); i < len; i++ {
		var arg Argument
		n2, err := surge.Unmarshal(&arg, r, m)
		n1 += n2
		m -= n2
		if err != nil {
			return n1, err
		}
		*args = append(*args, arg)
	}
	return n1, nil
}

func (args Arguments) SizeHint() uint32 {
	size := uint32(4)
	for _, arg := range args {
		size += arg.SizeHint()
	}
	return size
}

type Argument struct {
	Name  abi.String
	Value abi.Value
}

func (arg Argument) Marshal(w io.Writer) (uint32, error) {
	n1, err := arg.Name.Marshal(w)
	if err != nil {
		return n1, err
	}
	n2, err := arg.Value.Type().Marshal(w)
	if err != nil {
		return n1 + n2, err
	}
	n3, err := arg.Value.Marshal(w)
	if err != nil {
		return n1 + n2 + n3, err
	}
	return n1 + n2 + n3, nil
}

func (arg *Argument) Unmarshal(r io.Reader, m uint32) (uint32, error) {
	// Unmarshal arg name.
	var name abi.String
	n1, err := name.Unmarshal(r, m)
	m -= n1
	if err != nil {
		return n1, err
	}

	// Unmarshal arg type.
	var ty abi.Type
	n2, err := ty.Unmarshal(r, m)
	n1 += n2
	m -= n2
	if err != nil {
		return n1, err
	}

	// Unmarshal arg value.
	value, n3, err := abi.Unmarshal(r, ty, m)
	n1 += n3
	if err != nil {
		return n1, err
	}

	arg.Name = name
	arg.Value = value
	return n1, err
}

func (arg Argument) SizeHint() uint32 {
	return arg.Name.SizeHint() +
		arg.Value.Type().SizeHint() +
		arg.Value.SizeHint()
}

func (arg Argument) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"name":  arg.Name,
		"type":  arg.Value.Type(),
		"value": arg.Value,
	}
	return json.Marshal(m)
}

func (arg *Argument) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if len(m) > 3 {
		return fmt.Errorf("too many fields")
	}

	// Unmarshal arg name.
	rawName, ok := m["name"]
	if !ok {
		return fmt.Errorf("field not found: 'name'")
	}
	var name abi.String
	if err := name.UnmarshalJSON(rawName); err != nil {
		return err
	}

	// Unmarshal arg type.
	rawType, ok := m["type"]
	if !ok {
		return fmt.Errorf("field not found: 'type'")
	}
	var ty abi.Type
	if err := ty.UnmarshalJSON(rawType); err != nil {
		return err
	}

	// Unmarshal arg value.
	rawValue, ok := m["value"]
	if !ok {
		return fmt.Errorf("field not found: 'value'")
	}
	value, err := abi.UnmarshalJSON(rawValue, ty)
	if err != nil {
		return err
	}

	arg.Name = name
	arg.Value = value
	return nil
}
