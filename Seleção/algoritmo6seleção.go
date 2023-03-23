package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Println("Informe um número inteiro.")
	fmt.Scanln(&x)
	fmt.Println("Informe outro número inteiro.")
	fmt.Scanln(&y)

	if x < 0 || y < 0 {
		fmt.Println("A soma entre esses números é:", x+y)
	} else {
		fmt.Println("O produto entre esses números é:", x*y)
	}
}
