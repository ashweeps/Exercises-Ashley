package main

import (
	"bufio"
	"fmt"
	"os"
	"tareapaquetes/conversor"
	"tareapaquetes/vocales"
)

func main() {

	var dolares float64
	var moneda string

	fmt.Print("Ingrese el valor en dólares: ")
	fmt.Scan(&dolares)

	fmt.Print("Ingrese la moneda, tipo EUR, LB, WON o BTC: ")
	fmt.Scan(&moneda)

	resultado := conversor.Convertir(dolares, moneda)
	fmt.Println("El valor convertido es:", resultado)

	lector := bufio.NewReader(os.Stdin)

	lector.ReadString('\n')

	fmt.Print("Ingrese una palabra o frase: ")
	frase, _ := lector.ReadString('\n')

	a, e, i, o, u := vocales.Contar(frase)

	fmt.Println("Vocal A:", a)
	fmt.Println("Vocal E:", e)
	fmt.Println("Vocal I:", i)
	fmt.Println("Vocal O:", o)
	fmt.Println("Vocal U:", u)
}
