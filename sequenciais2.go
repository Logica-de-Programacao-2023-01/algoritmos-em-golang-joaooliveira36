package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número para mostrar seu dobro, triplo e quadruplo: ")
	fmt.Scan(&num)

	fmt.Println("O dobro do seu número é", num*2)
	fmt.Println("O triplo do seu número é", num*3)
	fmt.Println("O quadruplo do seu número é", num*4)
}
