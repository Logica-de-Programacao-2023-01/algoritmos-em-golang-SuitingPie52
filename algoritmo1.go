package main

import "fmt"

func main() {

	var x int
	var y int
	var z int

	fmt.Println("Informe um número inteiro.")
	fmt.Scanln(&y)
	fmt.Println("Informe mais um, para somar com o anterior.")
	fmt.Scanln(&x)
	fmt.Println("Informe um terceiro valor para a soma.")
	fmt.Scanln(&z)
	fmt.Println("Sua soma é:", x+y+z)
}
