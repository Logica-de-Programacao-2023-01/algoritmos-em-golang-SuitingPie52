package main

import "fmt"

func main() {
	var x int

	fmt.Println("Informe um número inteiro.")
	fmt.Scanln(&x)
	fmt.Println("Seu dobro é: ", x*2)
	fmt.Println("Seu triplo é: ", x*3)
	fmt.Println("Seu quádruplo é: ", x*4)
}
