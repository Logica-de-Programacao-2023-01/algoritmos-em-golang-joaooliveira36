package main

import "fmt"

func main() {
	var num1, num2, num3 float32

	fmt.Println("Esse algoritmo calcula a media ponderada de 3 numeros, digite-os.")

	fmt.Print("Qual o seu primeiro número? ")
	fmt.Scan(&num1)
	fmt.Print("Qual o seu segundo número? ")
	fmt.Scan(&num2)
	fmt.Print("Qual o seu terceiro número? ")
	fmt.Scan(&num3)

	fmt.Print("A media ponderada desses números é: ", ((num1*2)+(num2*3)+(num3*5))/10)
}
