package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Informe seu salário. ")
	fmt.Scanln(&salario)

	if salario < 1000 {
		fmt.Println("Seu salário com 10% de aumento é:", salario+(salario/10))
	} else {
		fmt.Println("Seu salário com 5% de aumento é:", salario+(salario*0.05))
	}
}
