package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	nombre := "Betito"
	fmt.Println("Obten una letra de Betito", string(nombre[2]))

	for i := 0; i < 7; i++ {
		fmt.Println("Valor indice", string(nombre[i]))
	}
}
