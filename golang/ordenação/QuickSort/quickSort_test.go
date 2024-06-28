package main

import "testing"

func TestQuickSort(t *testing.T) {

	naoOrdenado := [6]int{5, 3, 1, 4, 2, 0}
	ordenado := quickSortInicio(naoOrdenado)
	esperado := [6]int{0, 1, 2, 3, 4, 5}

	if ordenado != esperado {
		t.Errorf("esperado: %d, resultado: %d", esperado, ordenado)
	}
}
