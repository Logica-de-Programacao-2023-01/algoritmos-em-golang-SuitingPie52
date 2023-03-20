package main

import "fmt"

func main() {

	var peso float64

	fmt.Println("Informe seu peso em quilos.")
	fmt.Scanln(&peso)
	fmt.Println("Seu peso em libras é:", peso*2.20462, "lbs.")
}
