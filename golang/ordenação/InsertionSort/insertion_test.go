package main

import "testing"

func TestInsertion(t *testing.T) {

	naoOrdenada := [10]int{5, 3, 2, 1, 4, 6, 8, 7, 10, 9}
	ordenada := insertion(naoOrdenada)
	esperado := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if ordenada != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, ordenada)
	}
}
