package secondPhase

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func containsValue(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func Simulation(args []string) {
	// Sú štyri požadované argumenty v príkazovom riadku: p_graph (.1, .2, .3),
	// p_byzantine (.15, .30, .45), p_txDistribution (.01, .05, .10),
	// a numRounds (10, 20). Mali by ste sa pokúsiť otestovať svoj TrustedNode
	// kód pre všetky 3x3x3x2 = 54 kombinácií.

	numNodes := 100
	// pravdepodobnost existencie hrany
	p_graph, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatalf("error while parsingFloat: %v", err)
	}
	// % ze byzantsky
	p_byzantine, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		log.Fatalf("error while parsingFloat: %v", err)
	}
	// pravdepodobnost priradenia pociatocnej transakcie ku
	// kazdemu uzlu
	p_txDistribution, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		log.Fatalf("error while parsingFloat: %v", err)
	}
	// pocet kol
	numRounds, err := strconv.Atoi(args[3])
	if err != nil {
		log.Fatalf("error while parsing int: %v", err)
	}

	// vyberte, ktoré uzly sú byzantské a ktorým dôverujete
	nodes := make([]Node, numNodes)
	for i := 0; i < numNodes; i++ {
		if rand.Float64() < p_byzantine {
			byzan := ByzantineNode{}
			nodes[i] = &byzan
			// node = ByzantineNode{p_graph, p_byzantine, p_txDistribution, numRounds}
		} else {
			trusted := TrustedNode{}
			nodes[i] = &trusted
		}
		nodes[i].Node(p_graph, p_byzantine, p_txDistribution, numRounds)
	}

	// inicializovať náhodné sledovanie grafu
	var followees [][]bool
	for i := 0; i < numNodes; i++ {
		for k := 0; k < numNodes; k++ {
			if i == k {
				continue
			}
			if rand.Float64() < p_graph {
				followees[i][k] = true
			}
		}
	}

	// upozorni všetky uzly o ich nasledovníkoch
	for i := 0; i < numNodes; i++ {
		nodes[i].followeesSet(followees[i])
	}

	numTx := 500
	validTxIds := make([]int, numTx)
	for i := 0; i < numTx; i++ {
		r := rand.Int()
		if !containsValue(validTxIds, r) {
			validTxIds = append(validTxIds, r)
		}
	}

	// distribuuje 500 transakcií do všetkých uzlov a inicializuje ich
	// počiatočný stav transakcií, ktoré každý uzol počul. Distribúcia
	// je náhodná s pravdepodobnosťou p_txDistribution pre každý pár
	// Transkacia-Uzol
	for i := 0; i < numNodes; i++ {
		pendingTransactions := []Transaction{}
		for txID := range validTxIds {
			if rand.Float64() < p_txDistribution {
				pendingTransactions = append(pendingTransactions, Transaction{
					id: txID,
				})
			}
		}
		nodes[i].pendingTransactionSet(pendingTransactions)
	}

	// Simuluj numRounds-krat
	for round := 0; round < numRounds; round++ {

		// zhromaždiť všetky návrhy do mapy. Kľúčom je index uzla prijímajúceho
		// návrhy. Hodnota je ArrayList obsahujúci polia celých čísel 1x2. Prvým
		// prvkom každého poľa je ID navrhovanej transakcie a druhý
		// element je indexové číslo uzla navrhujúceho transakciu.

		allProposals := make(map[int][][2]int)
		for i := 0; i < numNodes; i++ {
			proposals := nodes[i].followersSend()
			for _, tx := range proposals {
				if !containsValue(validTxIds, tx.id) {
					continue
				}

				for j := 0; j < numNodes; j++ {
					if !followees[i][j] {
						continue
					}

					canditate := [2]int{}
					canditate[0] = tx.id
					canditate[1] = i
					allProposals[j] = append(allProposals[j], [2]int(canditate))

				}
			}
		}
		// Distribuuje návrhy k ich zamýšľaným príjemcom ako kandidátom
		for i := 0; i < numNodes; i++ {
			if _, ok := allProposals[i]; ok {
				nodes[i].followeesReceive(allProposals[i])
			}
		}
	}

	// vypíš výsledky

	for i := 0; i < numNodes; i++ {
		transactions := nodes[i].followersSend()
		fmt.Printf("Transaction ids that Node %d believes consensus on:", i)
		for _, tx := range transactions {
			fmt.Println(tx.id)
		}
		fmt.Println()
		fmt.Println()
	}
}
