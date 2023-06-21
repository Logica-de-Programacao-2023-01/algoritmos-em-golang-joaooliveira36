package main

import "fmt"

func main() {
	var salario float32

	fmt.Print("Qual é o seu salário? ")
	fmt.Scan(&salario)

	if salario < 1000 {
		salario = 1.1
	}
	if salario >= 1000 {
		salario = 1.05
	}

	fmt.Printf("Se seu seu salário sofresse um aumento ficaria %v\n", salario)
}
