package main

import "fmt"

func main() {
	for i := 0; i <= 19; {
		if i%2 != 0 {
			fmt.Println(i)
		}
		i++
	}
}
