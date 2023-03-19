package thirdphase

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func ThirdPhaseMain() {
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

	genesisBlock := NewBlock(nil, &publicKeyBob)
	genesisBlock.Finalize()

	blockchain := NewBlockchain(genesisBlock)
	blockHandler := HandleBlocks{
		BlockChain: blockchain,
	}
	block := NewBlock(genesisBlock.Hash, &publicKeyAlice)

	tx := Transaction{}
	tx.AddInput(genesisBlock.Coinbase.Hash, 0)
	tx.AddOutput(2, &publicKeyAlice)
	tx.AddOutput(2, &publicKeyAlice)
	tx.AddOutput(2, &publicKeyAlice)

	err = tx.SignTx(prGen_bob, 0)
	if err != nil {
		panic(err)
	}

	blockchain.TransactionPool.AddTransaction(tx)
	block.TransactionAdd(tx)
	block.Finalize()

	fmt.Printf("BlockProcess response: %v\n", blockHandler.BlockProcess(block))

}
