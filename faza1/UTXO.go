package faza1

type UTXO struct {
	txHash []byte
	index  int
}

func (u *UTXO) equals(other UTXO) bool {
	return true
}
