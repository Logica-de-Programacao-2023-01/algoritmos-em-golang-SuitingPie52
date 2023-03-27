package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Informe um número.")
	fmt.Scanln(&x)

	tabuada := 1
	for tabuada <= 10 {
		fmt.Println("Sua tabuada é:", x*float64(tabuada))
		tabuada++
	}
}
