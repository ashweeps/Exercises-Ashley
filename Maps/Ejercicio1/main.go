package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Ingrese una frase:")

	lector := bufio.NewScanner(os.Stdin)
	lector.Scan()

	frase := lector.Text()

	palabras := strings.Fields(frase)

	conteo := make(map[string]int)

	for _, palabra := range palabras {
		conteo[palabra]++
	}

	fmt.Println("CONTEO DE PALABRAS")

	var palabraMayor string
	mayor := 0

	for palabra, cantidad := range conteo {

		fmt.Println(palabra, ":", cantidad)

		if cantidad > mayor {
			mayor = cantidad
			palabraMayor = palabra
		}
	}

	if mayor == 0 {
		fmt.Println("No se ingresaron palabras")
	} else {
		fmt.Println("Palabra mas repetida:", palabraMayor)
		fmt.Println("Cantidad de repeticiones:", mayor)
	}
}
