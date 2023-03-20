package main

import "fmt"

func main() {
	var x int
	var y int
	var z int

	fmt.Println("Informe um número inteiro.")
	fmt.Scanln(&x)
	fmt.Println("Informe outro número inteiro.")
	fmt.Scanln(&y)
	fmt.Println("Informe mais um inteiro.")
	fmt.Scanln(&z)
	fmt.Print("Seu menor número é: ")
	if x < y && x < z {
		fmt.Print(x)
	} else if y < z && y < x {
		fmt.Print(y)
	} else if z < x && z < y {
		fmt.Print(z)
	} else {
		fmt.Print("Há mais de um menor.")
	}

}
