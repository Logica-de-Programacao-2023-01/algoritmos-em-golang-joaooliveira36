package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número")
	fmt.Scan(&numero)
	fmt.Println("A tabuada do seu número é: ")

	for numeros := 0; numeros <= 10; numeros++ {
		fmt.Println(numero * numeros)
	}
}
