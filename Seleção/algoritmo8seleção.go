package main

import "fmt"

func main() {
	var dia int

	fmt.Println("Informe um número inteiro de 1 a 7.")
	fmt.Scanln(&dia)

	if dia == 1 {
		fmt.Println("Seu dia é domingo.")
	} else if dia == 2 {
		fmt.Println("Seu dia é segunda-feira.")
	} else if dia == 3 {
		fmt.Println("Seu dia é terça-feira.")
	} else if dia == 4 {
		fmt.Println("Seu dia é quarta-feira.")
	} else if dia == 5 {
		fmt.Println("Seu dia é quinta-feira.")
	} else if dia == 6 {
		fmt.Println("Seu dia é sexta-feira.")
	} else if dia == 7 {
		fmt.Println("Seu dia é domingo.")
	}
}
