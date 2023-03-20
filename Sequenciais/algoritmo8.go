package main

import "fmt"

func main() {

	var dias float64
	var diaria float64

	fmt.Println("Informe a quantidade de dias trabalhados.")
	fmt.Scanln(&dias)
	fmt.Println("Informe o valor da sua diária.")
	fmt.Scanln(&diaria)
	fmt.Println("Seu salário é: R$", dias*diaria)
}
