package thirdphase

import (
	"bytes"
	"fmt"
)

const CUT_OFF_AGE = 12

type Blockchain struct {
	// BlockNode BlockNode
	TransactionPool TransactionPool
	BlockNode       []BlockNode
}

type BlockNode struct {
	B        Block
	Parent   *BlockNode
	Children []*BlockNode
	Height   int
	UPool    UTXOPool
}

func NewBlockNode(b Block, parent *BlockNode, uPool UTXOPool) *BlockNode {
	n := BlockNode{}
	n.B = b
	n.Parent = parent
	n.Children = make([]*BlockNode, 0)
	n.UPool = uPool
	if parent != nil {
		n.Height = parent.Height + 1
		parent.Children = append(parent.Children, &n)
	} else {
		n.Height = 1
	}

	return &n
}

func (b *Blockchain) GetUTXOPoolAtMaxHeight() UTXOPool {
	highestBlock := BlockNode{}
	for _, blockNode := range b.BlockNode {
		if blockNode.Height > highestBlock.Height {
			highestBlock = blockNode
		}
	}
	return highestBlock.UPool
}

/**
 * vytvor prázdny blockchain iba s prvým (Genesis) blokom. Predpokladajme, že
 * {@code genesisBlock} je platný blok
 */
func NewBlockchain(genesisBlock Block) *Blockchain {
	coinbaseTx := genesisBlock.Coinbase
	index := 0
	utxoPool := NewUTXOPool()
	utxo := UTXO{
		TxHash: genesisBlock.Coinbase.Hash,
		Index:  index,
	}
	utxoPool.AddUTXO(utxo, coinbaseTx.Outputs[index])
	blockNode := NewBlockNode(genesisBlock, nil, *utxoPool)
	return &Blockchain{
		BlockNode:       []BlockNode{*blockNode},
		TransactionPool: TransactionPool{},
	}
}

/** Získaj najvyšší (maximum height) blok */
func (b *Blockchain) GetBlockAtMaxHeight() (Block, int) {
	highestBlock := BlockNode{}
	maxHeight := 0
	for _, blockNode := range b.BlockNode {
		if blockNode.Height > highestBlock.Height {
			highestBlock = blockNode
			maxHeight = blockNode.Height
		}
	}
	return highestBlock.B, maxHeight
}

/**
* Pridaj {@code block} do blockchainu, ak je platný. Kvôli platnosti by mali
* byť všetky transakcie platné a blok by mal byť na
* {@code height > (maxHeight - CUT_OFF_AGE)}.
*
* Môžete napríklad vyskúšať vytvoriť nový blok nad blokom Genesis (výška bloku
* 2), ak height blockchainu je {@code <=
* CUT_OFF_AGE + 1}. Len čo {@code height > CUT_OFF_AGE + 1}, nemôžete vytvoriť
* nový blok vo výške 2.
*
* @return true, ak je blok úspešne pridaný
 */

func (b *Blockchain) BlockAdd(block Block) bool {
	// if prevBlockHash is nil or his hash length is 0 (genesis) returns false
	if block.PrevBlockHash == nil || len(block.PrevBlockHash) == 0 {
		fmt.Println("previousBlockHash is incorrect")
		return false
	}

	// gets the father Node
	parent, err := b.getParent(block)
	if err != nil {
		fmt.Println("did not get parent")
		return false
	}

	// this checks whether the height of potentional new block
	// is not heigher then maxHeight - CUT_OFF_AGE
	_, maxHeight := b.GetBlockAtMaxHeight()
	if parent.Height+1 <= maxHeight-CUT_OFF_AGE {
		fmt.Println("did not meet the condition parent.Height+1 <= maxHeight-CUT_OFF_AGE")
		return false
	}

	// this checks whether all of the transactions are valid
	handler := HandleTxs{
		UTXOPool: parent.UPool,
	}

	validTxs := handler.Handler(block.Txs)
	if len(validTxs) != len(block.Txs) {
		fmt.Println("not validated")
		return false
	}

	// adds a coinbase tx on index 0 to a handlers UTXOPool
	utxoPool := handler.UTXOPool
	utxo := UTXO{
		TxHash: block.Coinbase.Hash,
		Index:  0,
	}
	utxoPool.AddUTXO(utxo, block.Coinbase.Outputs[0])

	// removes Transactions from transaction pool
	for _, tx := range validTxs {
		b.TransactionPool.RemoveTransaction(tx.Hash)
	}

	node := BlockNode{
		B:      block,
		Parent: &parent,
		UPool:  utxoPool,
	}
	// pointer to children node is added to the parent node Children array
	parent.Children = append(parent.Children, &node)
	b.BlockNode = append(b.BlockNode, node)

	b.cleanOldBlocks(maxHeight)
	return true
}

/** Pridaj transakciu do transakčného poolu */
func (b *Blockchain) TransactionAdd(tx Transaction) {
	b.TransactionPool.AddTransaction(tx)
}

func (b *Blockchain) getParent(bl Block) (BlockNode, error) {
	for _, potentionalParent := range b.BlockNode {
		if bytes.Compare(potentionalParent.B.Hash, bl.PrevBlockHash) == 0 {
			return potentionalParent, nil
		}
	}
	return BlockNode{}, fmt.Errorf("did not find parent")
}

// cleans up old blocks
func (b *Blockchain) cleanOldBlocks(maxHieght int) {
	for i, node := range b.BlockNode {
		if node.Height+1 < maxHieght-CUT_OFF_AGE {
			remove(b.BlockNode, i)
		}
	}
}

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(slice []BlockNode, s int) []BlockNode {
	return append(slice[:s], slice[s+1:]...)
}
