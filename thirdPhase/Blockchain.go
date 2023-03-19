package thirdphase

const CUT_OFF_AGE = 12

type Blockchain struct {
	// BlockNode BlockNode
	BlockNode []BlockNode
}

type BlockNode struct {
	B        Block
	Parent   *BlockNode
	Children []*BlockNode
	Height   int
	UPool    UTXOPool
}

func (n *BlockNode) NewBlockNode(b Block, parent BlockNode, uPool UTXOPool) *BlockNode {
	n.B = b
	n.Parent = &parent
	n.Children = make([]*BlockNode, 0)
	n.UPool = uPool
	if &parent != nil {
		n.Height = parent.Height + 1
		parent.Children = append(parent.Children, n)
	} else {
		n.Height = 1
	}

	return n
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
	return &Blockchain{}
}

/** Získaj najvyšší (maximum height) blok */
func (b *Blockchain) GetBlockAtMaxHeight() Block {
	highestBlock := BlockNode{}
	for _, blockNode := range b.BlockNode {
		if blockNode.Height > highestBlock.Height {
			highestBlock = blockNode
		}
	}
	return highestBlock.B
}

/** Získaj UTXOPool na ťaženie nového bloku na vrchu najvyššieho (max height) bloku */
func (b *Blockchain) GetTransactionPool() *TransactionPool {
	return &TransactionPool{}
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
	return false
}

/** Pridaj transakciu do transakčného poolu */
func (b *Blockchain) TransactionAdd(tx Transaction) {

}
