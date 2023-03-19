package secondPhase

type TrustedNode struct {
}

func (b *TrustedNode) Node(p_graph float64, p_byzantine float64, p_txDistribution float64, numRounds int) {
	return
}

func (b *TrustedNode) followeesSet(followees []bool) {
	return
}

func (b *TrustedNode) pendingTransactionSet(pendingTransactions []Transaction) {
	return
}

func (b *TrustedNode) followersSend() []Transaction {
	return []Transaction{}
}

func (b *TrustedNode) followeesReceive(candidates [][2]int) {
	return
}
