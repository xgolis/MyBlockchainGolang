package faza

type UTXOPool struct {
	/**
	 * Aktuálna zbierka UTXO, pričom každé z nich je mapované na zodpovedajúci
	 * výstup transakcie
	 */
	H map[*UTXO]Output
}

func NewUTXOPool() *UTXOPool {
	utxoMap := make(map[*UTXO]Output)
	return &UTXOPool{
		H: utxoMap,
	}
}

func (u *UTXOPool) AddUTXO(utxo UTXO, txOut Output) {
	u.H[&utxo] = txOut
}

func (u *UTXOPool) RemoveUTXO(utxo UTXO) {
	delete(u.H, &utxo)
}

func (u *UTXOPool) GetTxOutput(ut UTXO) Output {
	return u.H[&ut]
}

func (u *UTXOPool) Contains(utxo UTXO) bool {
	_, ok := u.H[&utxo]
	return ok
}

func (u *UTXOPool) GetAllUTXO() []UTXO {
	var allUTXO []UTXO
	for u := range u.H {
		allUTXO = append(allUTXO, *u)
	}
	return allUTXO
}
