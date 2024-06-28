package main

import "fmt"

func particao(entrada [6]int, low, high int) ([6]int, int) {
	pivo := entrada[high]
	i := low

	for j := low; j < high; j++ {
		if entrada[j] < pivo {
			entrada[i], entrada[j] = entrada[j], entrada[i]
			i++
		}
	}

	entrada[i], entrada[high] = entrada[high], entrada[i]
	return entrada, i
}

func quickSort(entrada [6]int, low, high int) [6]int {
	if low < high {
		var p int
		entrada, p = particao(entrada, low, high)
		entrada = quickSort(entrada, low, p-1)
		entrada = quickSort(entrada, p+1, high)
	}

	return entrada
}

func quickSortInicio(entrada [6]int) [6]int {
	return quickSort(entrada, 0, len(entrada)-1)
}

func main() {
	naoOrdenada := [6]int{5, 3, 1, 4, 2, 0}
	ordenada := quickSortInicio(naoOrdenada)

	fmt.Println("Lista ordenada: ", ordenada)
}
