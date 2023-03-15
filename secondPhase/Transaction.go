package secondPhase

type Transaction struct {
	id int
}

func (t *Transaction) equals(tx Transaction) bool {
	if t.id != tx.id {
		return false
	}
	return true
}

func (t *Transaction) hashCode() int {
	return t.id
}
