package thirdphase

import (
	"crypto/rsa"
	"fmt"
)

type HandleBlocks struct {
	BlockChain *Blockchain
}

/**
 * pridaj {@code block} do blockchainu ak je platný.
 *
 * @return true ak je blok platný a bol pridaný, inak false
 */
func (h *HandleBlocks) BlockProcess(block Block) bool {
	if block.Hash == nil {
		return false
	}
	return h.BlockChain.BlockAdd(block)
}

/** vytvor nový {@code block} nad max height {@code block} */
func (h *HandleBlocks) BlockCreate(myAddress *rsa.PublicKey) Block {
	parent, _ := h.BlockChain.GetBlockAtMaxHeight()
	parentHash := parent.Hash
	current := NewBlock(parentHash, myAddress)
	uPool := h.BlockChain.GetUTXOPoolAtMaxHeight()
	txPool := h.BlockChain.TransactionPool
	handler := HandleTxs{
		UTXOPool: uPool,
	}
	txs := txPool.GetTransactions()
	rTxs := handler.Handler(txs)
	if len(txs) != len(rTxs) {
		fmt.Println("not validated")
		return Block{
			Hash: nil,
		}
	}
	for i := 0; i < len(rTxs); i++ {
		current.TransactionAdd(rTxs[i])
	}
	current.Finalize()
	if h.BlockChain.BlockAdd(current) {
		return current
	} else {
		return Block{
			Hash: nil,
		}
	}
}

func (h *HandleBlocks) TxProcess(tx Transaction) {
	h.BlockChain.TransactionAdd(tx)
}
