package main

import "fmt"

func main() {
	var peso float32

	fmt.Print("Digite um peso em kg: ")
	fmt.Scan(&peso)
	fmt.Println("Esse peso em libras ficaria:", peso*2.205)
}
