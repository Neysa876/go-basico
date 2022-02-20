package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		par := i % 2
		if par == 0 {
			fmt.Println("Numero par: ", i)
		}
	}
}
