package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("Seu número e divisivel por 3 e por 5")
	} else {
		fmt.Println("Seu número não é divisel por 3 e 5 ao mesmo tempo")
	}
}
