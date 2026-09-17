package main

import "fmt"

func main() {
	var edad int = 15
	var temperatura float32 = 21.3
	var activo bool = true
	var mensaje string = "Bienvenida"
	var dato byte = 255
	var dias [5]string = [5]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes"}
	var numeros []float64 = []float64{1.1, 2.2, 3.3}

	fmt.Printf("Edad: %v---TipoDato: %T\n", edad, edad)
	fmt.Printf("Temperatura: %v---TipoDato: %T\n", temperatura, temperatura)
	fmt.Printf("Activo: %v---TipoDato: %T\n", activo, activo)
	fmt.Printf("Mensaje: %v---TipoDato: %T\n", mensaje, mensaje)
	fmt.Printf("Dato: %v---TipoDato: %T\n", dato, dato)
	fmt.Printf("Días: %v---TipoDato: %T\n", dias, dias)
	fmt.Printf("Números: %v---TipoDato: %T\n", numeros, numeros)

	fmt.Println("Tu edad es de: ", edad)
	fmt.Println("La temperatura es: ", temperatura)
	fmt.Println("El activo es: ", activo)
	fmt.Println("El mensaje es : ", mensaje)
	fmt.Println("El dato es: ", dato)
	fmt.Println("Los días son: : ", dias)
	fmt.Println("Los números son: ", numeros)

	fmt.Println("Tu edad es de : ", edad, "Tu estado es: ", activo)

	var i int
	var j float64
	fmt.Print("Ingresa dos valores: ")
	fmt.Scanf("%d %f", &i, &j)
	fmt.Println("Resultado: ", (float64(i) * j * j))

}
