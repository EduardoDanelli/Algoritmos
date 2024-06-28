package main

import "fmt"

func Bubble(entrada [10]int) [10]int {
	troca := true

	for troca {
		troca = false

		for i := 1; i < len(entrada); i++ {
			if entrada[i-1] > entrada[i] {
				entrada[i], entrada[i-1] = entrada[i-1], entrada[i]
				troca = true
			}
		}
	}

	return entrada
}

func main() {
	entradaDesordenada := [10]int{9, 4, 7, 3, 1, 10, 8, 6, 5, 2}

	ordenado := Bubble(entradaDesordenada)

	fmt.Println("Lista desordenada: ", entradaDesordenada)
	fmt.Println("Lista ordenada: ", ordenado)
}
