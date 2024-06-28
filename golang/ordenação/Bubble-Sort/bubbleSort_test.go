package main

import "testing"

func TestBubble(t *testing.T) {

	naoOrdenado := [10]int{9, 4, 7, 3, 1, 10, 8, 6, 5, 2}
	ordenado := Bubble(naoOrdenado)
	esperado := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if ordenado != esperado {
		t.Errorf("esperado '%d', resultado '%d'", esperado, ordenado)
	}
}
