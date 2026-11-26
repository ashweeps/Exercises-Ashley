package main

import "fmt"

type InformacionProducto interface {
	MostrarInfo() string
}

type ProductoBase struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func (p ProductoBase) MostrarInfo() string {
	fmt.Printf("ID: %d | Nombre: %s | Precio: %.2f | Cantidad: %d",
		p.ID, p.Nombre, p.Precio, p.Cantidad)

	return "Producto: " + p.Nombre
}

type Electronico struct {
	ProductoBase
	GarantiaMeses int
}

func (e Electronico) MostrarInfo() string {
	producto := e.ProductoBase.MostrarInfo()

	fmt.Printf(" | Garantía: %d meses", e.GarantiaMeses)

	return producto
}

type Alimento struct {
	ProductoBase
	FechaExpiracion string
}

func (a Alimento) MostrarInfo() string {
	producto := a.ProductoBase.MostrarInfo()

	fmt.Printf(" | Fecha de expiración: %s", a.FechaExpiracion)

	return producto
}

func NuevoProductoElectronico(id int, nombre string, precio float64, cantidad int, garantia int) Electronico {
	return Electronico{
		ProductoBase: ProductoBase{
			ID:       id,
			Nombre:   nombre,
			Precio:   precio,
			Cantidad: cantidad,
		},
		GarantiaMeses: garantia,
	}
}

func NuevoProductoAlimento(id int, nombre string, precio float64, cantidad int, fechaExp string) Alimento {
	return Alimento{
		ProductoBase: ProductoBase{
			ID:       id,
			Nombre:   nombre,
			Precio:   precio,
			Cantidad: cantidad,
		},
		FechaExpiracion: fechaExp,
	}
}

type Inventario struct {
	Productos []InformacionProducto
}

func (i *Inventario) AgregarProducto(p InformacionProducto) {
	i.Productos = append(i.Productos, p)
}
func (i Inventario) ListarProductos() {

	fmt.Println("INVENTARIO DE PRODUCTOS")

	for posicion, producto := range i.Productos {
		fmt.Println("\nProducto:", posicion+1)
		producto.MostrarInfo()
	}
}
func main() {

	inventario := Inventario{}

	electronico1 := NuevoProductoElectronico(
		1, "Laptop MSI", 850.50, 5, 24,
	)

	electronico2 := NuevoProductoElectronico(
		2, "Cheescake de mora", 25.99, 15, 12,
	)

	alimento1 := NuevoProductoAlimento(
		3, "Leche", 1.25, 30, "15/12/2026",
	)

	alimento2 := NuevoProductoAlimento(
		4, "Yogurt", 0.85, 20, "10/12/2026",
	)

	inventario.AgregarProducto(electronico1)
	inventario.AgregarProducto(electronico2)
	inventario.AgregarProducto(alimento1)
	inventario.AgregarProducto(alimento2)

	inventario.ListarProductos()
}
