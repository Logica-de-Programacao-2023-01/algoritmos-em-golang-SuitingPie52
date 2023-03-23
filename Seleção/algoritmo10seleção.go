package main

import "fmt"

func main() {
	var idade uint

	fmt.Println("Informe a sua idade.")
	fmt.Scanln(&idade)
	fmt.Print("Sua classificação é: ")

	if idade <= 9 {
		fmt.Print("mirim.")
	} else if idade >= 10 && idade <= 13 {
		fmt.Println("infantil.")
	} else if idade >= 14 && idade <= 17 {
		fmt.Println("juvenil.")
	} else {
		fmt.Println("adulto.")
	}
}
