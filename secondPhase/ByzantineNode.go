package secondPhase

/**
 * Tento Byzantský uzol by sa dal považovať za vypnutý.
 * Nikdy nevysiela žiadne transakcie ani neodpovedá
 * na komunikáciu s inými uzlami.
 *
 * Toto je len jeden príklad (najjednoduchší) takéhoto
 * byzantského (škodlivého) uzla.
 */

type ByzantineNode struct {
}

func (b *ByzantineNode) Node(p_graph float64, p_byzantine float64, p_txDistribution float64, numRounds int) {
	return
}

func (b *ByzantineNode) followeesSet(followees []bool) {
	return
}

func (b *ByzantineNode) pendingTransactionSet(pendingTransactions []Transaction) {
	return
}

func (b *ByzantineNode) followersSend() []Transaction {
	return []Transaction{}
}

func (b *ByzantineNode) followeesReceive(candidates [][2]int) {
	return
}
