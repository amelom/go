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

// PairList : A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}
func main() {

	cupons := map[string]float64{
		"MLA1": 400,
		"MLA2": 400,
		"MLA3": 260,
		"MLA4": 80,
		"MLA5": 90,
	}
	var amount float64 = 8000
	arr := calculate(&cupons, &amount)
	log.Println(arr)
}

/*
* Funcion que recibe una lista y monto, la cual calcula el escenario optimó para la mayor cantidad de items de la lista por precio base
 */
// calculate func
func calculate(cupons *map[string]float64, amount *float64) []string {
	arrPromo := []string{}
	count := 0
	items := make(map[string]float64)
	// Evalua que los items no esten repetidos
	for j, i := range *cupons {
		if items[j] == 0 {
			items[j] = i
		}
	}

	p := orderMap(&items)

	for _, k := range p {
		if count+k.Value <= int(*amount) {
			count = count + k.Value
			arrPromo = append(arrPromo, k.Key)
		}
	}
	return arrPromo
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
