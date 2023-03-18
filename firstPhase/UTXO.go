package firstPhase

import "hash/fnv"

type UTXO struct {
	TxHash []byte
	Index  int
}

/**
* Porovná toto UTXO s tým, ktoré bolo zadané v {@code other}, považuje ich za
* rovnocenné ak majú pole {@code TxHash} s rovnakým obsahom a rovnaké hodnoty
* {@code Index}
 */
func (u *UTXO) equals(utxo UTXO) bool {
	hash := utxo.TxHash
	in := utxo.Index

	if len(hash) != len(u.TxHash) || u.Index != in {
		return false
	}

	for i := 0; i < len(hash); i++ {
		if hash[i] != u.TxHash[i] {
			return false
		}
	}

	return true
}

func hashCode(arr []byte) int {
	h := fnv.New32a()
	h.Write(arr)
	return int(h.Sum32())
}

/**
* Jednoduchá implementácia UTXO hashCode, ktorá rešpektuje rovnosť UTXOs //
* (t.j. utxo1.equals (utxo2) => utxo1.hashCode () == utxo2.hashCode ())
 */
func (u *UTXO) hashCode() int {
	hash := 1
	hash = hash*17 + u.Index
	hash = hash*31 + hashCode(u.TxHash)
	return hash
}

/** Porovná toto UTXO so špecifikovaným v {@code utxo} */
func (u *UTXO) compareTo(utxo UTXO) int {
	hash := utxo.TxHash
	in := utxo.Index
	if in > u.Index {
		return -1
	} else if in < u.Index {
		return 1
	} else {
		len1 := len(u.TxHash)
		len2 := len(hash)
		if len2 > len1 {
			return -1
		} else if len2 < len1 {
			return 1
		} else {
			for i := 0; i < len1; i++ {
				if hash[i] > u.TxHash[i] {
					return -1
				} else if hash[i] < u.TxHash[i] {
					return 1
				}
			}
			return 0
		}
	}
}
