package thirdphase

type TransactionPool struct {
	H map[*ByteArrayWrapper]Transaction
}

func (t *TransactionPool) AddTransaction(tx Transaction) {
	if t.H == nil {
		t.H = make(map[*ByteArrayWrapper]Transaction)
	}
	hash := ByteArrayWrapper{
		Contents: tx.Hash,
	}
	t.H[&hash] = tx
}

func (t *TransactionPool) RemoveTransaction(txHash []byte) {
	hash := ByteArrayWrapper{
		Contents: txHash,
	}
	delete(t.H, &hash)
}

func (t *TransactionPool) GetTransaction(txHash []byte) Transaction {
	hash := ByteArrayWrapper{
		Contents: txHash,
	}
	return t.H[&hash]
}

func (t *TransactionPool) GetTransactions() []Transaction {
	T := make([]Transaction, 0)
	for _, tx := range t.H {
		T = append(T, tx)
	}
	return T
}
