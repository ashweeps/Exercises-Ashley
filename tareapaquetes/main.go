package main

import (
	"fmt"
	"tareapaquetes/conversor"
	"tareapaquetes/vocales"
)

func main() {

	var dolares float64
	var moneda string

	fmt.Print("Ingrese el valor en dólares: ")
	fmt.Scan(&dolares)

	fmt.Print("Ingrese la moneda (EUR, LB, WON, BTC): ")
	fmt.Scan(&moneda)

	resultado := conversor.Convertir(dolares, moneda)
	fmt.Println("El valor convertido es:", resultado)

	var frase string

	fmt.Print("Ingrese una palabra: ")
	fmt.Scan(&frase)

	a, e, i, o, u := vocales.Contar(frase)

	fmt.Println("Vocal A:", a)
	fmt.Println("Vocal E:", e)
	fmt.Println("Vocal I:", i)
	fmt.Println("Vocal O:", o)
	fmt.Println("Vocal U:", u)
}
