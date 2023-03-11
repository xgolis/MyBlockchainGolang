package faza

type UTXOPool struct {
	/**
	 * Aktuálna zbierka UTXO, pričom každé z nich je mapované na zodpovedajúci
	 * výstup transakcie
	 */
	H map[*UTXO]Output
}

func (u *UTXOPool) addUTXO(utxo UTXO, txOut Output) {
	u.H[&utxo] = txOut
}

func (u *UTXOPool) removeUTXO(utxo UTXO) {
	delete(u.H, &utxo)
}

func (u *UTXOPool) getTxOutput(ut UTXO) Output {
	return u.H[&ut]
}

func (u *UTXOPool) contains(utxo UTXO) bool {
	_, ok := u.H[&utxo]
	return ok
}

func (u *UTXOPool) getAllUTXO() []UTXO {
	var allUTXO []UTXO
	for u := range u.H {
		allUTXO = append(allUTXO, *u)
	}
	return allUTXO
}
