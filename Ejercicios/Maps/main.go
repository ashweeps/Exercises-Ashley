package main

import "fmt"

func ActividadGanadora(votos map[string]int) string {

	var ganadora string
	mayor := -1

	for actividad, cantidad := range votos {

		if cantidad > mayor {

			mayor = cantidad
			ganadora = actividad
		}
	}

	return ganadora
}

func main() {

	votos := map[string]int{
		"Deportes":    0,
		"Videojuegos": 0,
		"Cine":        0,
		"Musica":      0,
	}

	var opcion int

	for i := 1; i <= 5; i++ {

		fmt.Println("\nVOTO", i)

		fmt.Println("1. Deportes")
		fmt.Println("2. Videojuegos")
		fmt.Println("3. Cine")
		fmt.Println("4. Musica")

		fmt.Print("Seleccione una actividad: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			votos["Deportes"]++

		case 2:
			votos["Videojuegos"]++

		case 3:
			votos["Cine"]++

		case 4:
			votos["Musica"]++

		default:
			fmt.Println("Opcion no valida")
			i--
		}
	}

	fmt.Println("\nRESULTADOS DE LA VOTACION")

	for actividad, cantidad := range votos {

		fmt.Println(actividad, ":", cantidad, "votos")
	}

	ganadora := ActividadGanadora(votos)

	fmt.Println("\nActividad con mas votos:", ganadora)
}
