package main

import "fmt"

func main() {
	var number, sum, count float64

	fmt.Println("Digite uma sequência de números inteiros. (Digite 0 para sair)")

	for {
		fmt.Scanln(&number)
		if number == 0 {
			break
		}
		sum += number
		count++
	}
	if sum == 0 {
		fmt.Println("Você não informou nenhum número.")
	} else {
		fmt.Println("Sua média é:", float64(sum)/float64(count))
	}
}
