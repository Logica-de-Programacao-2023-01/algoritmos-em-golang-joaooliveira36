package main

import "fmt"

func main() {
	var salario_atual float32
	fmt.Println("Esse algoritmo calcula seu salario com aumeto de 15%")

	fmt.Print("Qual o seu salario? ")
	fmt.Scan(&salario_atual)

	fmt.Println("Se seu salario aumentasse 15% ficaria:", salario_atual*1.15)
}
