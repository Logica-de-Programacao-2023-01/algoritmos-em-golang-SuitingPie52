package main

import "fmt"

func main() {
	var x int

	fmt.Println("Informe um número.")
	fmt.Scanln(&x)

	if x%3 == 0 && x%5 == 0 {
		fmt.Println("Este número é divisível por 3 e 5 ao mesmo tempo.")
	} else {
		fmt.Println("Este número não é divisível por 3 e 5 ao mesmo tempo.")
	}
}
