package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite outro número: ")
	fmt.Scan(&num2)

	if num1 < num2 {
		fmt.Printf("O numero %d é menor\n", num1)
	}
	if num2 < num1 {
		fmt.Printf("O número %d é menor\n", num2)
	}
}
