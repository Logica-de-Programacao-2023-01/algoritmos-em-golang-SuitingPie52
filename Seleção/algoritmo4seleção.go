package main

import "fmt"

func main() {
	var altura float64
	var peso float64
	var sexo string

	fmt.Println("Informe seu sexo. (Digite H para homem e M para mulher)")
	fmt.Scanln(&sexo)
	fmt.Println("Informe sua altura em metros.")
	fmt.Scanln(&altura)
	fmt.Println("Informe seu peso em quilos.")
	fmt.Scanln(&peso)
	i := peso / (altura * altura)

	if sexo == "H" && i > 24.9 {
		fmt.Println("Você está acima do peso.")
	} else if sexo == "H" && i < 20 {
		fmt.Println("Você está abaixo do peso.")
	} else if sexo == "H" && i >= 20 && i <= 24.9 {
		fmt.Println("Você está no seu peso ideal.")
	} else if sexo == "M" && i > 23.9 {
		fmt.Println("Você está acima do peso.")
	} else if sexo == "M" && i < 19 {
		fmt.Println("Você está abaixo do peso.")
	} else if sexo == "M" && i >= 19 && i <= 23.9 {
		fmt.Println("Você está no seu peso ideal.")
	} else {
		fmt.Println("error 404")
	}
}
