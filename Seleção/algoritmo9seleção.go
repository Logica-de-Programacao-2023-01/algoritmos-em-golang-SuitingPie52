package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64

	fmt.Println("Informe 3 números diferentes.")
	fmt.Scanln(&x, &y, &z)
	fmt.Println("Sua ordem crescente é: ")

	if x > y && y > z {
		fmt.Println(z, ",", y, ",", x)
	} else if x > z && z > y {
		fmt.Println(y, ",", z, ",", x)
	} else if y > x && x > z {
		fmt.Println(z, ",", x, ",", y)
	} else if y > z && z > x {
		fmt.Println(x, ",", z, ",", y)
	} else if z > x && x > y {
		fmt.Println(y, ",", x, ",", z)
	} else if z > y && y > x {
		fmt.Println(x, ",", y, ",", z)
	} else {
		fmt.Println("error 404")
	}
}
