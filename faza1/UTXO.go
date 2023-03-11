package main

import "hash/fnv"

type UTXO struct {
	txHash []byte
	index  int
}

/**
* Porovná toto UTXO s tým, ktoré bolo zadané v {@code other}, považuje ich za
* rovnocenné ak majú pole {@code txHash} s rovnakým obsahom a rovnaké hodnoty
* {@code index}
 */
func (u *UTXO) equals(utxo UTXO) bool {
	hash := utxo.txHash
	in := utxo.index

	if len(hash) != len(u.txHash) || u.index != in {
		return false
	}

	for i := 0; i < len(hash); i++ {
		if hash[i] != u.txHash[i] {
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
	hash = hash*17 + u.index
	hash = hash*31 + hashCode(u.txHash)
	return hash
}

/** Porovná toto UTXO so špecifikovaným v {@code utxo} */
func (u *UTXO) compareTo(utxo UTXO) int {
	hash := utxo.txHash
	in := utxo.index
	if in > u.index {
		return -1
	} else if in < u.index {
		return 1
	} else {
		len1 := len(u.txHash)
		len2 := len(hash)
		if len2 > len1 {
			return -1
		} else if len2 < len1 {
			return 1
		} else {
			for i := 0; i < len1; i++ {
				if hash[i] > u.txHash[i] {
					return -1
				} else if hash[i] < u.txHash[i] {
					return 1
				}
			}
			return 0
		}
	}
}
