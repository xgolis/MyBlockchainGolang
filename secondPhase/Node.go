package secondPhase

type Node interface {
	Node(p_graph float64, p_byzantine float64, p_txDistribution float64, numRounds int)
	// ByzantineNode(p_graph float64, p_byzantine float64, p_txDistribution float64, numRounds int)
	followeesSet(followees []bool)
	pendingTransactionSet(pendingTransactions []Transaction)
	followersSend() []Transaction
	followeesReceive(candidates [][2]int)
}
