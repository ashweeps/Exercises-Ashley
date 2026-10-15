package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func MostrarProducto(p Producto) {
	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Precio: $", p.Precio)
	fmt.Println("Cantidad:", p.Cantidad)
}

func CalcularTotal(p Producto) float64 {
	total := p.Precio * float64(p.Cantidad)
	return total
}

func VerificarStock(p Producto) string {

	if p.Cantidad > 0 {
		return "Disponible"
	} else {
		return "Sin stock"
	}
}

func main() {

	producto1 := Producto{"Arroz", 1.25, 10}
	producto2 := Producto{"Leche", 0.95, 5}
	producto3 := Producto{"Pan", 0.50, 0}

	productos := []Producto{producto1, producto2, producto3}

	fmt.Println("LISTA DE PRODUCTOS")

	for _, producto := range productos {

		MostrarProducto(producto)

		fmt.Println("Valor total: $", CalcularTotal(producto))

		fmt.Println("Estado:", VerificarStock(producto))
	}
}
