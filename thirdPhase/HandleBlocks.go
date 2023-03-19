package thirdphase

import "crypto/rsa"

type HandleBlocks struct {
	BlockChain Blockchain
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
	parent := h.BlockChain.GetBlockAtMaxHeight()
	parentHash := parent.GetHash()
	current := Block{}
	current.Block(parentHash, myAddress)
	uPool := h.BlockChain.GetUTXOPoolAtMaxHeight()
	txPool := h.BlockChain.GetTransactionPool()
	handler := HandleTxs{
		UTXOPool: uPool,
	}
	txs := txPool.GetTransactions()
	rTxs := handler.Handler(txs)
	for i := 0; i < len(rTxs); i++ {
		current.TransactionAdd(rTxs[i])
	}
	current.finalize()
	if h.BlockChain.BlockAdd(current) {
		return current
	} else {
		return Block{}
	}
}

func (h *HandleBlocks) TxProcess(tx Transaction) {
	h.BlockChain.TransactionAdd(tx)
}
