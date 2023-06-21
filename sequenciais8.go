package main

import "fmt"

func main() {
	var dias_trabalhados, valor_hora, horas_dia float32

	fmt.Print("Quantos dias você trabalha? ")
	fmt.Scan(&dias_trabalhados)
	fmt.Print("Qual o valor da sua hora? ")
	fmt.Scan(&valor_hora)
	fmt.Print("Quantas horas você trabalha por dia? ")
	fmt.Scan(&horas_dia)

	fmt.Println("O valor do seu salário é de:", (horas_dia*valor_hora)*dias_trabalhados)
}
