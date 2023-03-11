package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/xgolis/MyBlockchainGolang/faza1"
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

	fmt.Printf("bob:%d, alice: %d", prGen_alice, prGen_bob)
	// prGen_bob = *rsa.PrivateKey
	// crypto.PrivateKey.Public()
	_ = faza1.UTXOPool{}
}
