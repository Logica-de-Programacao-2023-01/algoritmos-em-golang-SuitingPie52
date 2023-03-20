package main

import "fmt"

func main() {

	var preço float64

	fmt.Println("Informe o preço do seu produto.")
	fmt.Scanln(&preço)
	fmt.Println("O seu produto com 10% de desconto é: R$", preço-(preço/10))
}
