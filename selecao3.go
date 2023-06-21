package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número para ver se é par ou impar: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("O número %d é par\n", num)
	} else {
		fmt.Printf("O numer %d não é par\n", num)
	}
}
