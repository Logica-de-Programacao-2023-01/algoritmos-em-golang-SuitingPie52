package main

import "fmt"

func main() {

	var idade int

	fmt.Println("Informe sua idade (em anos).")
	fmt.Scanln(&idade)
	fmt.Println("Sua idade em dias (aproximadamente) é:", idade*365, "dias.")
}
