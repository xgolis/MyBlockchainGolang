package firstPhase

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

// Add tests for the first phase here
func TestHandler() bool {
	prGen_bob, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic("unable to generate private key")
	}
	prGen_alice, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic("unable to generate private key")
	}

	publicKeyBob := prGen_bob.PublicKey
	publicKeyAlice := prGen_alice.PublicKey

	// fmt.Printf("bob:%d, alice: %d\n", publicKeyBob, publicKeyAlice)

	tx := Transaction{}
	tx.AddOutput(10, &publicKeyBob)

	initialHash := []byte{0}
	tx.AddInput(initialHash, 0)

	tx.SignTx(prGen_bob, 0)

	utxoPool := NewUTXOPool()
	utxo := UTXO{
		TxHash: tx.Hash,
		Index:  0,
	}
	utxoPool.AddUTXO(utxo, tx.GetOutput(0))

	handleTxs := HandleTxs{
		UTXOPool: *utxoPool,
	}

	tx2 := Transaction{}
	tx2.AddInput(tx.Hash, 0)

	tx2.AddOutput(5, &publicKeyAlice)
	tx2.AddOutput(3, &publicKeyAlice)
	tx2.AddOutput(2, &publicKeyAlice)

	err = tx2.SignTx(prGen_bob, 0)
	if err != nil {
		fmt.Println(err)
	}

	utxo = UTXO{
		TxHash: tx2.Hash,
		Index:  0,
	}
	utxoPool.AddUTXO(utxo, tx2.GetOutput(0))
	utxo = UTXO{
		TxHash: tx2.Hash,
		Index:  1,
	}
	utxoPool.AddUTXO(utxo, tx2.GetOutput(1))
	utxo = UTXO{
		TxHash: tx2.Hash,
		Index:  2,
	}
	utxoPool.AddUTXO(utxo, tx2.GetOutput(2))

	handleTxs = HandleTxs{
		UTXOPool: *utxoPool,
	}

	txs := []Transaction{tx2}
	validTxs := handleTxs.Handler(txs)
	if len(validTxs) == len(txs) {
		return true
	}
	return false
}
