package main

import "fmt"

type persona struct {
	nombre    string
	apellidos string
	edad      int
}

func main() {
	persona1 := persona{"Bill", "Gates", 65}
	persona2 := persona{"Neysa", "Alarcon", 25}

	fmt.Println("Persona1: ", persona1)
	fmt.Println("Persona2: ", persona2)

	persona2.edad += 1

	fmt.Println("Persona2 edad cambiada", persona2)

	persona1.saludar("Hola")
	persona2.saludar("Hi")

	nuevoAno := persona1.cumpleanos()
	fmt.Println("¿Cuántos años tienes Bill?", nuevoAno)

	fmt.Println("¿Cuántos años tienes Ney?", persona2.cumpleanos())
}

func (p persona) saludar(saludo string) {
	fmt.Println(saludo + ", " + p.nombre)
}

func (pers persona) cumpleanos() int {
	return pers.edad + 1
}
