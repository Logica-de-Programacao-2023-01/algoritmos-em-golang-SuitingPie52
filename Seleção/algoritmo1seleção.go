package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Println("Informe um número inteiro.")
	fmt.Scanln(&x)
	fmt.Println("Informe outro número inteiro")
	fmt.Scanln(&y)
	fmt.Print("O maior número é: ")
	if x > y {
		fmt.Println(x)
	}
	if y > x {
		fmt.Println(y)
	}
}
