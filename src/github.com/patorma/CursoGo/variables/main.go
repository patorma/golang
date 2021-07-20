package main

import "fmt"

func main() {
	// es opcionable poner el tipo de dato
	//aqui con :=  digo declaro la variable y le asigno ""
	dog, cat := "🐕", "🐱"
	// las variables siguen siendo estaticas
	// para reasignar un valor de una variable en go pongo signo igual
	// cat = "gatico"
	// si una de las variables aparece es nueva se debe usar :=
	cat, face := "gato", "😊"

	fmt.Println(dog, cat, face)
}
