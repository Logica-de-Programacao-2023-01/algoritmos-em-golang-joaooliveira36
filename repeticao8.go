package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	for divisor := 1; divisor <= numero; divisor++ {
		if numero%divisor == 0 {
			fmt.Println(divisor)
		}
	}
}
