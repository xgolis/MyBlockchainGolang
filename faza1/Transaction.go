package faza1

import (
	"crypto/rsa"
	"crypto/sha256"
	"math/big"

	"github.com/pounze/ByteBuffer_golang/src/ByteBuffer"
)

type Transaction struct {
	// input   Input
	// output  Output
	hash    []byte
	inputs  []Input
	outputs []Output
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
// 	t.hash = tx.hash
// 	t.inputs = tx.inputs
// 	t.outputs = tx.outputs
// }

func (t *Transaction) addInput(prevTxHash []byte, outputIndex int) {
	in := Input{
		prevTxHash:  prevTxHash,
		outputIndex: outputIndex,
	}
	t.inputs = append(t.inputs, in)
}

func (t *Transaction) addOutput(value float64, address *rsa.PublicKey) {
	op := Output{
		value:   value,
		address: address,
	}
	t.outputs = append(t.outputs, op)
}

func RemoveIndex(s []Input, index int) []Input {
	return append(s[:index], s[index+1:]...)
}

func (t *Transaction) removeInput(value int) {
	t.inputs = RemoveIndex(t.inputs, value)
}

func (t *Transaction) removeInputUTXO(ut UTXO) {
	for i := 0; i < len(t.inputs); i++ {
		in := t.inputs[i]
		u := &UTXO{
			txHash: in.prevTxHash,
			index:  in.outputIndex,
		}
		if u.equals(ut) {
			t.removeInput(i)
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
	if index > len(t.inputs) {
		return nil
	}
	in := t.inputs[index]
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

	for _, op := range t.outputs {
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
	t.inputs[index].addSignature(signature)
}

func (t *Transaction) getTx() []byte {
	Tx := []byte{}
	for _, in := range t.inputs {
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
	for _, op := range t.outputs {
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

func (t *Transaction) finalize() {
	md := sha256.New()
	_, err := md.Write(t.getTx())
	if err != nil {
		panic(err)
	}
	t.hash = md.Sum(nil)
}

func (t *Transaction) setHash(h []byte) {
	t.hash = h
}

func (t *Transaction) getInput(index int) Input {
	if index < len(t.inputs) {
		return t.inputs[index]
	}
	return Input{}
}

func (t *Transaction) getOutput(index int) Output {
	if index < len(t.inputs) {
		return t.outputs[index]
	}
	return Output{}
}
