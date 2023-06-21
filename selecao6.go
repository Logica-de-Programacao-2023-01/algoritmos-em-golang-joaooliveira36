package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Me fale o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Println("Me fale o segundo número: ")
	fmt.Scan(&num2)

	if num1 >= 0 && num2 >= 0 {
		multiplicação := num1 * num2
		fmt.Print("A multiplicalçao dos seus números é igual a: ", multiplicação)
	} else {
		soma := num1 + num2
		fmt.Print("A soma dos seus números é igual a: ", soma)
	}
}
