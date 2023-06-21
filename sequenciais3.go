package main

import "fmt"

func main() {
	var peso, altura float32

	fmt.Print("Qual seu peso? ")
	fmt.Scan(&peso)

	fmt.Print("Qual sua altura? ")
	fmt.Scan(&altura)

	fmt.Println("O seu IMC é de: ", peso/(altura*altura))
}
