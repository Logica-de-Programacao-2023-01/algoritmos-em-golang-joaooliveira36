package main

import "fmt"

func main() {
	for multiplos_de_3 := 0; multiplos_de_3 <= 30; multiplos_de_3++ {
		if multiplos_de_3%3 == 0 {
			fmt.Println(multiplos_de_3)
		}
	}
}
