package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Qual o seu primeiro número? ")
	fmt.Scan(&num1)

	fmt.Print("A soma dos seus dois números e de: ", num1+num2)
}
