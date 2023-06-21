package main

import "fmt"

func main() {
	var peso, altura, imc float32

	fmt.Print("Qual seu peso? ")
	fmt.Scan(&peso)

	fmt.Print("Qual sua altura? ")
	fmt.Scan(&altura)

	imc = peso / (altura * altura)

	fmt.Println("O seu IMC é de: ", imc)

	if imc <= 18.5 {
		fmt.Println("O seu peso está abaixo o normal")
	} else if imc <= 24.9 {
		fmt.Println("O seu peso está normal")
	} else if imc >= 25 {
		fmt.Println("Você está acima do peso")
	}
}
