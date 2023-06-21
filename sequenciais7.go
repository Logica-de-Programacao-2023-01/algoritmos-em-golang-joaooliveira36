package main

import "fmt"

func main() {
	var num int

	fmt.Println("Digite um numero para mostrar seu sucessor e antecessor: ")
	fmt.Scan(&num)

	fmt.Println("O numero sucessor é", num+1)
	fmt.Println("O numero antecessor é", num-1)
}
