package main

import "fmt"

func main() {
	var numero, maior_numero int

	fmt.Println("Digite seus números (Digite 0 para parar): ")

	for {
		fmt.Scan(&numero)
		if numero == 0 {
			break
		}
		if numero > maior_numero {
			maior_numero = numero
		}
	}
	if maior_numero == 0 {
		fmt.Println("Nenhum valor diferente de 0 foi digitada")
	} else {
		fmt.Println("O maior número é:", maior_numero)
	}
}
