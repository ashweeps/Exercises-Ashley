package main

import "fmt"

func averageGrade(suma float64, cantidad int) float64 {
	return suma / float64(cantidad)
}

func opcion1() {
	var cantidad int
	var nota, suma float64

	fmt.Print("Cantidad de estudiantes: ")
	fmt.Scan(&cantidad)

	for i := 1; i <= cantidad; i++ {
		fmt.Print("Nota del estudiante ", i, ": ")
		fmt.Scan(&nota)
		suma += nota
	}

	promedio := averageGrade(suma, cantidad)
	fmt.Println("Promedio:", promedio)

	if promedio >= 70 {
		fmt.Println("Aprobado")
	} else {
		fmt.Println("Reprobado")
	}

	switch {
	case promedio >= 90:
		fmt.Println("Excellent performance")
	case promedio >= 80:
		fmt.Println("Good performance")
	case promedio >= 70:
		fmt.Println("Satisfactory performance")
	default:
		fmt.Println("Needs improvement")
	}
}

func opcion2() {
	var n, suma int

	fmt.Print("Ingrese n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		suma += i
	}

	fmt.Println("La suma es:", suma)
}

func opcion3() {
	var c float64

	fmt.Print("Celsius: ")
	fmt.Scan(&c)

	fmt.Println("Fahrenheit:", (c*9/5)+32)
}

func opcion4() {
	var f float64

	fmt.Print("Fahrenheit: ")
	fmt.Scan(&f)

	fmt.Println("Celsius:", (f-32)*5/9)
}

func main() {
	var opcion int

	for {
		fmt.Println("\nMENU")
		fmt.Println("1. Promedio de estudiantes")
		fmt.Println("2. Suma del 1 al n")
		fmt.Println("3. Celsius a Fahrenheit")
		fmt.Println("4. Fahrenheit a Celsius")
		fmt.Println("0. Salir")

		fmt.Print("Opcion: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			opcion1()
		case 2:
			opcion2()
		case 3:
			opcion3()
		case 4:
			opcion4()
		case 0:
			fmt.Println("Haz salido de este programa")
			return
		default:
			fmt.Println("Debes seleccionar uno de los numeros")
		}
	}
}
