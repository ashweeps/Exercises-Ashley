package main

import "fmt"

type Producto struct {
	Nombre string
	precio float64
	stock  int
}

func NewProduct(nombre string, precio float64, stock int) Producto {
	return Producto{
		Nombre: nombre,
		precio: precio,
		stock:  0,
	}

}

func calcularStock(producto Producto) float64 {
	return float64(producto.stock) * producto.precio
}
func (p *Producto) NuevoStock(cantidad int) {
	if cantidad > 0 {
		p.stock += cantidad
	}
}
func (producto *Producto) Vender(cantidad int) bool {
	if cantidad <= 0 || cantidad > producto.stock {
		return false
	}
	producto.stock -= cantidad
	return true
}

func main() {
	product := Producto{}
	product.Nombre = "Intel Core Ultra 7"
	product.precio = 850.90
	product.stock = 200

	product2 := Producto{
		Nombre: "Teclado",
		precio: 30.5,
		stock:  15,
	}
	fmt.Println(product)

	fmt.Println(product2)
	fmt.Println("Valor del inventario:", calcularStock(product))
	fmt.Println("Valor del inventario:", calcularStock(product2))

	cantidadVendida := 50
	ventaRealizada := product.Vender(cantidadVendida)
	if ventaRealizada {
		fmt.Println("Verificación de venta:", product.Nombre)
		fmt.Println("Vendido:", cantidadVendida)
		fmt.Println("Stock sobrante:", product.stock)
		fmt.Println("Valor del inventario restante:", calcularStock(product))
	} else {
		fmt.Println("La venta no se pudo realizar.")
		fmt.Println("Stock disponible:", product.stock)
	}

	cantidadVendida = 20
	ventaRealizada = product2.Vender(cantidadVendida)
	if ventaRealizada {
		fmt.Println("Verificación de venta:", product2.Nombre)
		fmt.Println("Vendido:", cantidadVendida)
		fmt.Println("Stock sobrante:", product2.stock)
		fmt.Println("Valor del inventario restante:", calcularStock(product2))
	} else {
		fmt.Println("La venta no se pudo realizar para:", product2.Nombre)
		fmt.Println("Solicitado:", cantidadVendida)
		fmt.Println("Stock disponible:", product2.stock)
	}
	product3 := NewProduct("mouse", 12.91, 50)
	fmt.Println(product3)
	//La tienda quiere q todos los prodcutos se queden con un stock 0

}
