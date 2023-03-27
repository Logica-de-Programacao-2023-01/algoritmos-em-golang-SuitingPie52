package main

import "fmt"

func main() {
	var x, maior int

	fmt.Println("Informe uma sequência de números inteiros. (Digite 0 para sair)")

	for {
		fmt.Scanln(&x)

		if x == 0 {
			break
		}
		if x >= maior {
			maior = x
		}
	}
	fmt.Println(" O maior número é:", maior)
}
