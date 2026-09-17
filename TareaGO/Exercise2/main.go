package main

import "fmt"

func main() {

	var n int

	fmt.Println("Ingresa un numero entero positivo:")
	fmt.Scan(&n)

	fmt.Println("*****Tabla de multiplicar*****")

	for i := 1; i <= 10; i++ {
		fmt.Println(n, "x", i, "=", n*i)
	}

	fmt.Println("*****Es Par o Impar*****")

	if n%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

	switch {
	case n >= 1 && n <= 5:
		fmt.Println("Numero pequeño")

	case n >= 6 && n <= 10:
		fmt.Println("Numero mediano")

	case n > 10:
		fmt.Println("Numero grande")
	}
}
