package main

func Linear(entrada [5]int, chave int) int {
	for i, valor := range entrada {
		if valor == chave {
			return i
		}
	}

	return -1
}
