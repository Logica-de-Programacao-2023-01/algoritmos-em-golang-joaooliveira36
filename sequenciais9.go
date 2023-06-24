package main

import "fmt"

func main() {
	var preço float32

	fmt.Print("")
	fmt.Scan(&preço)
	fmt.Println("O produto com 15% de desconto ficaria:", preço*0.85)
}
