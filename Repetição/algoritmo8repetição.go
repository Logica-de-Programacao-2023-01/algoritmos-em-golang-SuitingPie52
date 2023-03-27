package main

import "fmt"

func main() {
	var x int

	fmt.Println("Informe um número.")
	fmt.Scanln(&x)
	fmt.Println("Seus divisores são:")

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Println(i)

		}
	}
}