package main

import "fmt"

func main() {
	numerito := 45

	if numerito > 10 {
		fmt.Println("Soy mayor a 10")
	} else if numerito < 10 {
		fmt.Println("Soy un numero menor a 10")
	} else {
		fmt.Println("No soy un numerito :c")
	}

	if numerito == 45 {
		fmt.Println("Soy el numerito 45")
	}
}
