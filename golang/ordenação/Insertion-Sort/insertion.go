package main

import "fmt"

func insertion(entrada [10]int) [10]int {

	for i := 1; i < len(entrada); i++ {
		for j := i; j > 0 && entrada[j-1] > entrada[j]; j-- {
			entrada[j], entrada[j-1] = entrada[j-1], entrada[j]
		}
	}
	return entrada
}

func main() {

	naoOrdenada := [10]int{5, 3, 2, 1, 4, 6, 8, 7, 10, 9}
	ordenada := insertion(naoOrdenada)

	fmt.Println("Lista ordenada: ", ordenada)
}
