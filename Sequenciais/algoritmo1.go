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
	fmt.Println("Informe mais um número inteiro.")
	fmt.Scanln(&z)
	fmt.Print("A soma entre eles é: ", x+y+z, ".")
}
