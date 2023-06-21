package main

import "fmt"

func main() {

	for numero := 0; numero < 100; numero++ {

		if numero%3 == 0 && numero%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if numero%5 == 0 {
			fmt.Println("Buzz")
		} else if numero%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(numero)
		}
	}
}
