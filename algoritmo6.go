package main

import "fmt"

func main() {

	var salario float64

	fmt.Println("Informe seu salário.")
	fmt.Scanln(&salario)
	fmt.Println("Seu salário com 15% de aumento é:", salario+(salario/100)*15)
}
