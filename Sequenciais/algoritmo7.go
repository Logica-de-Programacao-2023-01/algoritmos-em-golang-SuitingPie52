package main

import "fmt"

func main() {

	var x int

	fmt.Println("Informe um número inteiro.")
	fmt.Scanln(&x)
	fmt.Println("Seu sucessor é:", x+1)
	fmt.Println("Seu antecessor é", x-1)
}
