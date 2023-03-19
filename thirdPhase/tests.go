package thirdphase

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func TestProcessBlock() error {

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
		return err
	}

	blockchain.TransactionPool.AddTransaction(tx)
	block.TransactionAdd(tx)
	block.Finalize()

	got := blockHandler.BlockProcess(block)

	if got != true {
		return fmt.Errorf("result of block process: %v", got)
	}
	return nil
}

func TestCreateBlock() error {

	// this block generates a random RSA keypairs
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

	// this block creates a genesisBlock and initializes blockchain
	genesisBlock := NewBlock(nil, &publicKeyBob)
	genesisBlock.Finalize()

	blockchain := NewBlockchain(genesisBlock)
	blockHandler := HandleBlocks{
		BlockChain: blockchain,
	}
	block := NewBlock(genesisBlock.Hash, &publicKeyAlice)

	// creates a Transaction
	tx := Transaction{}
	tx.AddInput(genesisBlock.Coinbase.Hash, 0)
	tx.AddOutput(2, &publicKeyAlice)
	tx.AddOutput(2, &publicKeyAlice)
	tx.AddOutput(2, &publicKeyAlice)

	// signs
	err = tx.SignTx(prGen_bob, 0)
	if err != nil {
		return err
	}

	// adds Txs to TxPool and block
	blockchain.TransactionPool.AddTransaction(tx)
	block.TransactionAdd(tx)
	block.Finalize()

	got := blockHandler.BlockCreate(&publicKeyAlice)

	if got.Hash == nil {
		return fmt.Errorf("resulting block hash is null")
	}
	return nil
}

func TestWrongValues() error {

	// this block generates a random RSA keypairs
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

	// this block creates a genesisBlock and initializes blockchain
	genesisBlock := NewBlock(nil, &publicKeyBob)
	genesisBlock.Finalize()

	blockchain := NewBlockchain(genesisBlock)
	blockHandler := HandleBlocks{
		BlockChain: blockchain,
	}
	block := NewBlock(genesisBlock.Hash, &publicKeyBob)

	// creates a Transaction
	tx := Transaction{}
	tx.AddInput(genesisBlock.Coinbase.Hash, 0)
	tx.AddOutput(2, &publicKeyAlice)
	tx.AddOutput(20, &publicKeyAlice)
	tx.AddOutput(2, &publicKeyAlice)

	// signs
	err = tx.SignTx(prGen_bob, 0)
	if err != nil {
		return err
	}

	// adds Txs to TxPool and block
	blockchain.TransactionPool.AddTransaction(tx)
	block.TransactionAdd(tx)
	block.Finalize()

	got := blockHandler.BlockCreate(&publicKeyBob)

	if got.Hash == nil {
		return nil
	}
	return fmt.Errorf("Block created")
}
