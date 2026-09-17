package main

import "fmt"

func saludar() {
	fmt.Println("Hola esta es mi primera función")
}
func Bienvenida(nombre string) {
	fmt.Println("Bienvenid@", nombre)
}

func main() {
	var usr string
	fmt.Println("Ingresa tu nombre:")
	fmt.Scan(&usr)
	saludar()
	Bienvenida(usr)

}
