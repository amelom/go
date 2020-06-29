//
// Copyright (C) 2020 - -
//

package main

import (
	"log"
	"sort"
)

// Pair struct
type Pair struct {
	Key   string
	Value int
}

type ArrSort struct {
	sum     float64
	itemArr []string
}

type ArrSortList struct {
	items []ArrSort
}

// PairList : A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}
func main() {

	cupons := map[string]float64{
		"MLA1": 37990,
		"MLA2": 3799.12,
		"MLA3": 7289,
		"MLA4": 80,
		"MLA5": 90,
	}
	var amount float64 = 500
	arr := calculate(&cupons, &amount)
	log.Println(arr)
}

/*
* Funcion que recibe una lista y monto, la cual calcula el escenario optimó para la mayor cantidad de items de la lista por precio base
 */
// calculate func
func calculate(cupons *map[string]float64, amount *float64) []string {
	arrPromo := []string{}

	items := make(map[string]float64)
	// Evalua que los items no esten repetidos
	for j, i := range *cupons {
		if items[j] == 0 {
			items[j] = i
		}
	}
	// ordene
	p := orderMap(&items)
	tempP := p
	var arrsortList ArrSortList
	var temp float64
	// contar
	for j := range tempP {
		var sortList ArrSort
		arrPromo2, sumArray, count := countArr(p, arrPromo, amount)
		sortList.itemArr = arrPromo2
		sortList.sum = sumArray
		arrsortList.items = append(arrsortList.items, sortList)
		p[j].Value = 0
		if count >= len(p) {
			temp = arrsortList.items[0].sum
			arrPromo = arrsortList.items[0].itemArr
			break
		}
	}

	for _, k := range arrsortList.items {
		if k.sum > temp {
			arrPromo = k.itemArr
		}
	}

	return arrPromo
}
func countArr(p PairList, arrPromo []string, amount *float64) ([]string, float64, int) {
	sumArray := 0
	count := 0
	for _, k := range p {
		if sumArray+k.Value <= int(*amount) {
			count++
			sumArray = sumArray + k.Value
			if k.Value != 0 {
				arrPromo = append(arrPromo, k.Key)
			}
		}
	}
	return arrPromo, float64(sumArray), count
}

// orderMap func
func orderMap(noble *map[string]float64) PairList {
	p := make(PairList, len(*noble))
	i := 0
	for k, v := range *noble {
		p[i] = Pair{k, int(v)}
		i++
	}
	sort.Sort(p)
	return p
}
