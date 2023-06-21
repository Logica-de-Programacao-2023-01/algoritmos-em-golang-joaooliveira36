package main

import (
	"fmt"
	"sort"
)

func main() {
	var numeros int

	fmt.Print("Digite a quantidade de números que você deseja inserir: ")
	fmt.Scan(&numeros)

	lista := make([]int, numeros)

	for numero := 0; numero < numeros; numero++ {
		fmt.Printf("Digite o número %d: ", numero+1)
		fmt.Scan(&lista[numero])
	}

	fmt.Println("A lista de números inseridos é:", lista)

	sort.Ints(lista)
	fmt.Print("A lista em ordem cresente fica: ", lista)
}
