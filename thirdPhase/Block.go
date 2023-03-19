package thirdphase

import (
	"crypto/rsa"
	"crypto/sha256"
)

type Block struct {
	Hash          []byte
	PrevBlockHash []byte
	Coinbase      Transaction
	Txs           []Transaction
}

const COINBASE = 6.25

/** {@code address} je adresa, na ktorú bude poslaná coinbase transakcia */
func NewBlock(prevHash []byte, address *rsa.PublicKey) Block {
	Coinbase := NewCoinbaseTx(COINBASE, address)
	Txs := []Transaction{}
	return Block{
		PrevBlockHash: prevHash,
		Coinbase:      Coinbase,
		Txs:           Txs,
	}
}

// func (b *Block) GetHash() []byte {
// 	return b.Hash
// }

// func (b *Block) GetPrevBlockHash() []byte {
// 	return b.PrevBlockHash
// }

// func (b *Block) GetTransactions() []Transaction {
// 	return b.Txs
// }

// func (b *Block) GetTransaction(index int) Transaction {
// 	return b.Txs[index]
// }

func (b *Block) TransactionAdd(tx Transaction) {
	b.Txs = append(b.Txs, tx)
}

func (b *Block) GetBlock() []byte {
	rawBlock := []byte{}
	if b.PrevBlockHash != nil {
		for i := 0; i < len(b.PrevBlockHash); i++ {
			rawBlock = append(rawBlock, b.PrevBlockHash[i])
		}
	}
	for i := 0; i < len(b.Txs); i++ {
		rawTx := b.Txs[i].GetTx()
		for j := 0; j < len(rawTx); j++ {
			rawBlock = append(rawBlock, rawTx[j])
		}
	}
	raw := make([]byte, len(rawBlock))
	for i := 0; i < len(raw); i++ {
		raw[i] = rawBlock[i]
	}
	return raw
}

func (b *Block) Finalize() {
	md := sha256.New()
	_, err := md.Write(b.GetBlock())
	if err != nil {
		panic(err)
	}
	b.Hash = md.Sum(nil)
}
