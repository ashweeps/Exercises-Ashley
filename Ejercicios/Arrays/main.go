package main

import "fmt"

func main() {

	var notas [6][4]float64

	var promedios []float64
	var notasAltas []float64
	var notasBajas []float64

	var sumaGeneral float64

	for p := 0; p < 6; p++ {

		fmt.Println("Estudiante", p+1)

		var suma float64

		for i := 0; i < 4; i++ {

			fmt.Print("Ingrese la nota de la materia ", i+1, ": ")
			fmt.Scan(&notas[p][i])

			suma += notas[p][i]
		}

		promedio := suma / 4

		promedios = append(promedios, promedio)

		mayor := notas[p][0]
		menor := notas[p][0]

		for i := 0; i < 4; i++ {

			if notas[p][i] > mayor {
				mayor = notas[p][i]
			}

			if notas[p][i] < menor {
				menor = notas[p][i]
			}
		}

		notasAltas = append(notasAltas, mayor)
		notasBajas = append(notasBajas, menor)

		sumaGeneral += promedio
	}

	fmt.Println("\nRESULTADOS")

	for p := 0; p < 6; p++ {

		fmt.Println("\nEstudiante", p+1)

		fmt.Println("Promedio:", promedios[p])
		fmt.Println("Nota mas alta:", notasAltas[p])
		fmt.Println("Nota mas baja:", notasBajas[p])
	}

	fmt.Println("\nPromedio general de la clase:", sumaGeneral/6)
}
