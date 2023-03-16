package main

import "fmt"

func main() {

	var x float64
	var y float64
	var z float64

	fmt.Println("Informe um número (Peso 2).")
	fmt.Scanln(&x)
	fmt.Println("Informe outro número (Peso 3).")
	fmt.Scanln(&y)
	fmt.Println("Agora, o último (Peso 5).")
	fmt.Scanln(&z)
	fmt.Println("Sua média ponderada é:", (x*2+y*3+z*5)/10)
}
