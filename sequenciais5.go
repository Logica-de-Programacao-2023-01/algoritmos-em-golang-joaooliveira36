package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Esse algoritmo calcula a sua idade em dias. Qual sua idade?")
	fmt.Scan(&idade)
	fmt.Print("Sua idade em dias é de: ", idade*365)
}
