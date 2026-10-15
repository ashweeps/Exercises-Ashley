package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var frase string

	fmt.Println("Ingrese una frase:")
	frase, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	palabras := strings.Fields(frase)

	conteo := map[string]int{}

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

	fmt.Println("Palabra mas repetida:", palabraMayor)
	fmt.Println("Cantidad de repeticiones:", mayor)
}
