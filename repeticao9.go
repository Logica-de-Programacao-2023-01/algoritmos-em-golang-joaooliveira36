package main

import "fmt"

func main() {
	var num, soma, count int
	fmt.Println("Digite vários números inteiros (digite 0 para parar):")

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		soma += num
		count++
	}

	if count == 0 {
		fmt.Println("Não foi digitado nenhum número diferente de zero.")
	} else {
		media := float64(soma) / float64(count)
		fmt.Print("A média aritmética é: ", media)
	}
}
