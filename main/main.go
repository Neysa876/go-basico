package main

import (
	"fmt"
)

func main() {
	frutas := []string{"pera", "manzana", "sandia"}
	//					0		1			2
	fmt.Println("Lista de frutas: ", frutas[2])

	frutas = append(frutas, "naranja", "arándano", "durazno")
	frutas[0] = "limon"
	for i := 0; i < 6; i++ {
		fmt.Println("Fruta ", frutas[i])
	}

	for i := 0; i < len(frutas); i++ {
		if frutas[i] == "pera" {
			fmt.Println("Hay coincidencias: ", frutas[i])
		} else {
			fmt.Println("No hay coincidencias.")
			break
		}
	}
}
