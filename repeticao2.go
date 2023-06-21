package main

import "fmt"

func main() {
	for numeros_pares := 0; numeros_pares <= 20; numeros_pares++ {
		if numeros_pares%2 == 0 {
			fmt.Println(numeros_pares)
		}
	}
}
