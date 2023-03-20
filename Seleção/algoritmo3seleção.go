package main

import "fmt"

func main() {
	var x int

	fmt.Println("Escreva um número inteiro.")
	fmt.Scanln(&x)
	fmt.Print("Seu número é: ")
	if x%2 == 0 {
		fmt.Print("par.")
	} else {
		fmt.Print("ímpar.")
	}
}
