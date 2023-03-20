package main

import "fmt"

func main() {
	var peso float64
	var altura float64

	fmt.Println("Informe seu peso em quilos.")
	fmt.Scanln(&peso)
	fmt.Println("Informe sua altura em metros.")
	fmt.Scanln(&altura)
	fmt.Println("Seu IMC é: ", peso/(altura*altura))
}
