package main

import "fmt"

func main() {
	var dia int

	fmt.Print("Digite um número de 1 a 7: ")
	fmt.Scan(&dia)

	dia_da_semana := ""

	if dia == 1 {
		dia_da_semana = "domingo"

	} else if dia == 2 {
		dia_da_semana = "segunda"

	} else if dia == 3 {
		dia_da_semana = "terça"

	} else if dia == 4 {
		dia_da_semana = "quarta"

	} else if dia == 5 {
		dia_da_semana = "quinta"

	} else if dia == 6 {
		dia_da_semana = "sexta"

	} else if dia == 7 {
		dia_da_semana = "sabado"

	}
	fmt.Printf("O dia da semana é %s", dia_da_semana)
}
