package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/xgolis/MyBlockchainGolang/faza"
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

	tx := faza.Transaction{}
	tx.AddOutput(10, &publicKeyBob)

	initialHash := []byte{0}
	tx.AddInput(initialHash, 0)

	tx.SignTx(prGen_bob, 0)

	utxoPool := faza.NewUTXOPool()
	utxo := faza.UTXO{
		TxHash: tx.Hash,
		Index:  0,
	}
	utxoPool.AddUTXO(utxo, tx.GetOutput(0))

	tx2 := faza.Transaction{}
	tx2.AddInput(tx.Hash, 0)

	tx2.AddOutput(5, &publicKeyAlice)
	tx2.AddOutput(3, &publicKeyAlice)
	tx2.AddOutput(2, &publicKeyAlice)

	err = tx2.SignTx(prGen_bob, 0)
	if err != nil {
		fmt.Println(err)
	}
}
