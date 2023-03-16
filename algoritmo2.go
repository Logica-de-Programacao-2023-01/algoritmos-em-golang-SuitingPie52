package main

import "fmt"

func main() {

	var x float64

	fmt.Println("Informe um número.")
	fmt.Scanln(&x)
	fmt.Println("Seu dobro é:", x*2)
	fmt.Println("Seu triplo é:", x*3)
	fmt.Println("E seu quádruplo é:", x*4)
}
