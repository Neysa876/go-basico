package main

import "fmt"

func main() {

	suma := 10 + 5
	resta := 10 - 5
	multiplicacion := 10 * 5
	division := 10 / 5
	resto := 10 % 5

	fmt.Println("Operadores matematicos", suma, resta, multiplicacion, division, resto)

	suma++
	fmt.Println("Aumento de var suma", suma)

	suma--
	fmt.Println("Decremento de var suma", suma)

	suma += 10
	fmt.Println("Aumento con parametro de var suma", suma)

	oper1 := 10 < 5
	oper2 := 10 > 5
	oper3 := 10 <= 5
	oper4 := 10 >= 5
	oper5 := 10 == 5
	oper6 := 10 != 5
	oper7 := 10 > 5 && 10 < 30
	oper8 := 10 < 7 || 10 > 2

	fmt.Println("Operadores lógicos", oper1, oper2, oper3, oper4, oper5, oper6, oper7, oper8)
}
