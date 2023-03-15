package faza

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/pounze/ByteBuffer_golang/src/ByteBuffer"
)

type Transaction struct {
	// input   Input
	// output  Output
	Hash    []byte   `json:"hash"`
	Inputs  []Input  `json:"inputs"`
	Outputs []Output `json:"outpus"`
}

type Output struct {
	value   float64
	address *rsa.PublicKey
}

type Input struct {
	prevTxHash  []byte
	outputIndex int
	signature   []byte
}

// func (o *Output) Output(v float64, addr crypto.PublicKey) {
// 	o.value = v
// 	o.address = addr
// }

// func (i *Input) Input(prevHash []byte, index int) {
// 	if prevHash == nil {
// 		i.prevTxHash = nil
// 	} else {
// 		i.prevTxHash = prevHash
// 	}
// 	i.outputIndex = index
// }

func (i *Input) addSignature(sig []byte) {
	if sig == nil {
		i.signature = nil
	} else {
		i.signature = sig
	}
}

// func newTransaction() *Transaction {
// 	return &Transaction{
// 		inputs:  []Input{},
// 		outputs: []Output{},
// 	}
// }

// func (t *Transaction) Transaction(tx Transaction) {
// 	t.Hash = tx.Hash
// 	t.inputs = tx.inputs
// 	t.outputs = tx.outputs
// }

func (t *Transaction) AddInput(prevTxHash []byte, outputIndex int) {
	in := Input{
		prevTxHash:  prevTxHash,
		outputIndex: outputIndex,
	}
	t.Inputs = append(t.Inputs, in)
}

func (t *Transaction) AddOutput(value float64, address *rsa.PublicKey) {
	op := Output{
		value:   value,
		address: address,
	}
	t.Outputs = append(t.Outputs, op)
}

func RemoveIndex(s []Input, index int) []Input {
	return append(s[:index], s[index+1:]...)
}

func (t *Transaction) RemoveInput(value int) {
	t.Inputs = RemoveIndex(t.Inputs, value)
}

func (t *Transaction) RemoveInputUTXO(ut UTXO) {
	for i := 0; i < len(t.Inputs); i++ {
		in := t.Inputs[i]
		u := &UTXO{
			TxHash: in.prevTxHash,
			Index:  in.outputIndex,
		}
		if u.equals(ut) {
			t.RemoveInput(i)
			return
		}
	}
}

// Source: chat gpt
func getExponent(key *rsa.PublicKey) *big.Int {
	return big.NewInt(int64(key.E))
}
func getModulus(key *rsa.PublicKey) *big.Int {
	return key.N
}

func (t *Transaction) getDataToSign(index int) []byte {
	sigData := []byte{}
	if index > len(t.Inputs) {
		return nil
	}
	in := t.Inputs[index]
	prevTxHash := in.prevTxHash

	b := ByteBuffer.Buffer{}
	b.PutInt(in.outputIndex)
	outputIndex := b.Array()

	if prevTxHash != nil {
		for i := 0; i < len(prevTxHash); i++ {
			sigData = append(sigData, prevTxHash[i])
		}
	}
	for i := 0; i < len(outputIndex); i++ {
		sigData = append(sigData, outputIndex[i])
	}

	for _, op := range t.Outputs {
		bo := ByteBuffer.Buffer{}
		bo.PutDouble(op.value)
		value := bo.Array()
		addressExponent := getExponent(op.address).Bytes()
		addressModulus := getModulus(op.address).Bytes()
		for i := 0; i < len(value); i++ {
			sigData = append(sigData, value[i])
		}
		for i := 0; i < len(addressExponent); i++ {
			sigData = append(sigData, addressExponent[i])
		}
		for i := 0; i < len(addressModulus); i++ {
			sigData = append(sigData, addressModulus[i])
		}
	}
	sigD := make([]byte, len(sigData))
	i := 0
	for _, sb := range sigData {
		sigD[i] = sb
		i++
	}
	return sigD
}

func (t *Transaction) addSignature(signature []byte, index int) {
	t.Inputs[index].addSignature(signature)
}

func (t *Transaction) GetTx() []byte {
	Tx := []byte{}
	for _, in := range t.Inputs {
		prevTxHash := in.prevTxHash
		b := ByteBuffer.Buffer{}
		b.PutInt(in.outputIndex)
		outputIndex := b.Array()
		signature := in.signature
		if prevTxHash != nil {
			for i := 0; i < len(prevTxHash); i++ {
				Tx = append(Tx, prevTxHash[i])
			}
		}
		for i := 0; i < len(outputIndex); i++ {
			Tx = append(Tx, outputIndex[i])
		}
		if signature != nil {
			for i := 0; i < len(signature); i++ {
				Tx = append(Tx, signature[i])
			}
		}
	}
	for _, op := range t.Outputs {
		b := ByteBuffer.Buffer{}
		b.PutDouble(op.value)
		value := b.Array()
		addressExponent := getExponent(op.address).Bytes()
		addressModulus := getModulus(op.address).Bytes()
		for i := 0; i < len(value); i++ {
			Tx = append(Tx, value[i])
		}
		for i := 0; i < len(addressExponent); i++ {
			Tx = append(Tx, addressExponent[i])
		}
		for i := 0; i < len(addressModulus); i++ {
			Tx = append(Tx, addressModulus[i])
		}
	}
	tx := make([]byte, len(Tx))
	i := 0
	for _, sb := range Tx {
		tx[i] = sb
		i++
	}
	return tx
}

func (t *Transaction) SignTx(pr_key *rsa.PrivateKey, input int) error {
	hashed := sha256.Sum256(t.getDataToSign(input))
	signature, err := rsa.SignPKCS1v15(rand.Reader, pr_key, crypto.SHA256, hashed[:])
	if err != nil {
		return fmt.Errorf("error while creating signature: %v", err)
	}
	t.addSignature(signature, input)
	t.finalize()
	return nil
}

func (t *Transaction) finalize() {
	md := sha256.New()
	_, err := md.Write(t.GetTx())
	if err != nil {
		panic(err)
	}
	t.Hash = md.Sum(nil)
}

func (t *Transaction) SetHash(h []byte) {
	t.Hash = h
}

func (t *Transaction) GetInput(index int) Input {
	if index < len(t.Inputs) {
		return t.Inputs[index]
	}
	return Input{}
}

func (t *Transaction) GetOutput(index int) Output {
	if index < len(t.Outputs) {
		return t.Outputs[index]
	}
	return Output{}
}
