package main

import "testing"

func TestLinear(t *testing.T) {
	resultado := Linear([5]int{1, 3, 6, 8, 4}, 8)
	esperado := 3

	if resultado != esperado {
		t.Errorf("esperado: %d, resultado: %d", esperado, resultado)
	}
}
