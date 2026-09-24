package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main() {

	fmt.Println("🐒Bienvenid@ a la clase de paquetes🐒")
	mensaje := saludo.Saludar("Juan")
	fmt.Println(mensaje)

	resultado := operaciones.Suma(5, 3)
	fmt.Println("La suma es:", resultado)
}
