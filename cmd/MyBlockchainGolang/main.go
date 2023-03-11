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

	prGen_bob, err := rsa.GenerateKey(rand.Reader, 32)
	if err != nil {
		panic("unable to generate private key")
	}
	prGen_alice, err := rsa.GenerateKey(rand.Reader, 32)
	if err != nil {
		panic("unable to generate private key")
	}

	pk_bob := prGen_bob.PublicKey
	// pk_alice := prGen_alice.PublicKey

	fmt.Printf("bob:%d, alice: %d\n", prGen_alice, prGen_bob)
	// prGen_bob = *rsa.PrivateKey
	// crypto.PrivateKey.Public()

	tx := faza.Transaction{}
	tx.AddOutput(10, &pk_bob)

	initialHash := []byte{0}
	tx.AddInput(initialHash, 0)

	tx.SignTx(prGen_bob, 0)
	// prGen_alice.Sign(rand.Reader, )
}
