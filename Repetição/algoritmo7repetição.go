package main

import "fmt"

func main() {
	numeros := 1

	for numeros <= 100 {
		if numeros%3 == 0 && numeros%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if numeros%3 == 0 {
			fmt.Println("Fizz")
		} else if numeros%5 == 0 {
			fmt.Println("Buzz")

		} else {
			fmt.Println(numeros)
		}
		numeros++
	}
}
