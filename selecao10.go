package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Qual sua idade? ")
	fmt.Scan(&idade)

	if idade <= 9 {
		fmt.Println("Voce é da faixa etária mirim")
	} else if idade <= 13 {
		fmt.Println("Voce é da faixa etária infantil")
	} else if idade <= 17 {
		fmt.Println("Voce é da faixa etária juvenil")
	} else if idade >= 18 {
		fmt.Println("Voce é da faixa etária adulto")
	}
}
