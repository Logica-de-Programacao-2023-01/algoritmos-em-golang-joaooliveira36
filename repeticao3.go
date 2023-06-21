package main

import "fmt"

func main() {
	for numeros_impares := 0; numeros_impares <= 19; numeros_impares++ {
		if numeros_impares%2 != 0 {
			fmt.Println(numeros_impares)
		}
	}
}
