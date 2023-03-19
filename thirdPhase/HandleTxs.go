package thirdphase

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

type HandleTxs struct {
	UTXOPool UTXOPool
}

/**
* Vytvorí verejný ledger (účtovnú knihu), ktorého aktuálny UTXOPool (zbierka nevyčerpaných
* transakčných výstupov) je {@code utxoPool}. Malo by to vytvoriť bezpečnú kópiu
* utxoPool pomocou konštruktora UTXOPool (UTXOPool uPool).
 */
// func (h *HandleTxs) HandleTxs(utxoPool UTXOPool) {
// 	h.UTXOPool = utxoPool
// }

// /**
// * @return aktuálny UTXO pool.
// * Ak nenájde žiadny aktuálny UTXO pool, tak vráti prázdny (nie nulový) objekt {@code UTXOPool}.
//  */
// func (h *HandleTxs) UTXOPoolGet() UTXOPool {
// 	if h.UTXOPool.H == nil {
// 		return UTXOPool{}
// 	}

// 	return h.UTXOPool
// }

/**
 * @return true, ak
 * (1) sú všetky nárokované výstupy {@code tx} v aktuálnom UTXO pool,
 * (2) podpisy na každom vstupe {@code tx} sú platné,
 * (3) žiadne UTXO nie je nárokované viackrát, equals
 * (4) všetky výstupné hodnoty {@code tx}s sú nezáporné a
 * (5) súčet vstupných hodnôt {@code tx}s je väčší alebo rovný súčtu jej
 *     výstupných hodnôt; a false inak.
 */

// i made a seperate function for each condition for clarity
func (h *HandleTxs) TxIsValid(tx Transaction) bool {
	if !h.firstCondition(tx) {
		fmt.Println("first condition is invalid")
		return false
	}
	if !h.secondCondition(tx) {
		fmt.Println("second condition is invalid")
		return false
	}
	if !h.thirdCondition() {
		fmt.Println("third condition is invalid")
		return false
	}
	if !h.fourthCondition(tx) {
		fmt.Println("fourth condition is invalid")
		return false
	}
	if !h.fifthCondition(tx) {
		fmt.Println("fifth condition is invalid")
		return false
	}

	return true
}

func (h *HandleTxs) firstCondition(tx Transaction) bool {
	isInPool := false
	for _, txOutput := range tx.Outputs {
		for _, output := range h.UTXOPool.H {
			if output.address == txOutput.address && output.value == txOutput.value {
				isInPool = true
			}
		}
		if !isInPool {
			return false
		}
		isInPool = false
	}
	return true
}

func (h *HandleTxs) secondCondition(tx Transaction) bool {
	for i, input := range tx.Inputs {
		address := h.getAddress(input.outputIndex, input.prevTxHash)
		if address == nil {
			fmt.Println("nil address")
			return false
		}
		hashed := sha256.Sum256(tx.getDataToSign(i))
		err := rsa.VerifyPKCS1v15(address, crypto.SHA256, hashed[:], input.signature)
		if err != nil {
			fmt.Printf("error while verifying signature: %v", err)
			return false
		}
	}
	return true
}

func (h *HandleTxs) getAddress(outputIndex int, prevTxHash []byte) *rsa.PublicKey {
	tmpUTXO := UTXO{
		TxHash: prevTxHash,
		Index:  outputIndex,
	}
	output := h.UTXOPool.GetTxOutput(tmpUTXO)
	if output != (Output{}) {
		return output.address
	}

	return nil
}

func (h *HandleTxs) thirdCondition() bool {
	utxos := h.UTXOPool.GetAllUTXO()
	for i := 0; i < len(utxos); i++ {
		for k, utxo := range utxos {
			if k == i {
				break
			}
			if utxo.equals(utxos[i]) {
				// if utxo.hashCode() == utxos[i].hashCode() {
				return false
			}
		}
	}
	return true
}

func (h *HandleTxs) fourthCondition(tx Transaction) bool {
	for _, output := range tx.Outputs {
		if output.value < 0 {
			return false
		}
	}
	return true
}

// súčet vstupných hodnôt {@code tx}s je väčší alebo rovný súčtu jej
//   - výstupných hodnôt; a false inak.
func (h *HandleTxs) fifthCondition(tx Transaction) bool {
	sumOutputs := 0.0
	sumInputs := 0.0
	for i := 0; i < len(tx.Outputs); i++ {
		sumOutputs += tx.Outputs[i].value
	}
	for i := 0; i < len(tx.Inputs); i++ {
		value := h.getValue(tx.Inputs[i].outputIndex, tx.Inputs[i].prevTxHash)
		if value < 0 {
			return false
		}
		sumInputs += value
	}
	if sumOutputs > sumInputs {
		return false
	}
	return true
}

func (h *HandleTxs) getValue(outputIndex int, prevTxHash []byte) float64 {
	tmpUTXO := UTXO{
		TxHash: prevTxHash,
		Index:  outputIndex,
	}
	output := h.UTXOPool.GetTxOutput(tmpUTXO)
	if output != (Output{}) {
		return output.value
	}

	return -1
}

/**
 * Spracováva každú epochu (iteráciu) prijímaním neusporiadaného radu navrhovaných
 * transakcií, kontroluje správnosť každej transakcie, vracia pole vzájomne
 * platných prijatých transakcií a aktualizuje aktuálny UTXO pool podľa potreby.
 */
func (h *HandleTxs) Handler(possibleTxs []Transaction) []Transaction {
	validTransactions := make([]Transaction, 0)
	for _, transaction := range possibleTxs {
		isValid := h.TxIsValid(transaction)
		if isValid {
			// if the transaction is valid, append to validTransactions
			validTransactions = append(validTransactions, transaction)
		}
		for _, input := range transaction.Inputs {
			utxoToRemove := UTXO{
				TxHash: input.prevTxHash,
				Index:  input.outputIndex,
			}
			// remove all old input from current utxopool
			h.UTXOPool.RemoveUTXO(&utxoToRemove)
		}

		for i, output := range transaction.Outputs {
			newUTXO := UTXO{
				TxHash: transaction.Hash,
				Index:  i,
			}
			// add new outputs
			h.UTXOPool.AddUTXO(newUTXO, output)
		}
	}

	return validTransactions
}
