package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/xgolis/MyBlockchainGolang/firstPhase"
)

func main() {
	// páry klučov
	// key_bob := [32]byte{}
	// key_alice := [32]byte{}

	// for i := 0; i < 32; i++ {
	// 	key_bob[i] = byte(0)
	// 	key_alice[i] = byte(1)
	// }

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

	fmt.Printf("bob:%d, alice: %d\n", publicKeyBob, publicKeyAlice)
	// prGen_bob = *rsa.PrivateKey
	// crypto.PrivateKey.Public()

	tx := firstPhase.Transaction{}
	tx.AddOutput(10, &publicKeyBob)

	initialHash := []byte{0}
	tx.AddInput(initialHash, 0)

	tx.SignTx(prGen_bob, 0)

	utxoPool := firstPhase.NewUTXOPool()
	utxo := firstPhase.UTXO{
		TxHash: tx.Hash,
		Index:  0,
	}
	utxoPool.AddUTXO(utxo, tx.GetOutput(0))

	handleTxs := firstPhase.HandleTxs{
		UTXOPool: *utxoPool,
	}
	// handleTxs.HandleTxs(*utxoPool)
	// fmt.Printf("handleTxs.txIsValid(tx) returns: %t\n", handleTxs.TxIsValid(tx))
	// fmt.Println(handleTxs.UTXOPool)
	// handleTxs.Handler([]firstPhase.Transaction{tx})
	// fmt.Println(handleTxs.UTXOPool)
	// fmt.Printf("handleTxs.Handler([]Transaction{tx}) returns: %v", handleTxs.Handler([]firstPhase.Transaction{tx}))

	tx2 := firstPhase.Transaction{}
	tx2.AddInput(tx.Hash, 0)

	tx2.AddOutput(5, &publicKeyAlice)
	tx2.AddOutput(3, &publicKeyAlice)
	tx2.AddOutput(2, &publicKeyAlice)

	err = tx2.SignTx(prGen_bob, 0)
	if err != nil {
		fmt.Println(err)
	}

	// utxoPool = firstPhase.NewUTXOPool()
	utxo = firstPhase.UTXO{
		TxHash: tx2.Hash,
		Index:  0,
	}
	utxoPool.AddUTXO(utxo, tx2.GetOutput(0))
	utxo = firstPhase.UTXO{
		TxHash: tx2.Hash,
		Index:  1,
	}
	utxoPool.AddUTXO(utxo, tx2.GetOutput(1))
	utxo = firstPhase.UTXO{
		TxHash: tx2.Hash,
		Index:  2,
	}
	utxoPool.AddUTXO(utxo, tx2.GetOutput(2))

	handleTxs = firstPhase.HandleTxs{
		UTXOPool: *utxoPool,
	}
	// handleTxs.HandleTxs(*utxoPool)
	// fmt.Printf("handleTxs.txIsValid(tx) returns: %t", handleTxs.TxIsValid(tx2))
	handleTxs.Handler([]firstPhase.Transaction{tx2})
	// fmt.Printf("handleTxs.Handler([]Transaction{tx,tx2}) returns: %v", handleTxs.Handler([]firstPhase.Transaction{tx, tx2}))

}
