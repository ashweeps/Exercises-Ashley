package main

import "fmt"

type InformacionProducto interface {
	MostrarInfo() string
}

// Estructura base
type ProductoBase struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

// Método MostrarInfo
func (p ProductoBase) MostrarInfo() string {
	return fmt.Sprintln("ID:", p.ID,
		"Nombre:", p.Nombre,
		"Precio:", p.Precio,
		"Cantidad:", p.Cantidad)
}

// Estructura Electronico
type Electronico struct {
	ProductoBase
	GarantiaMeses int
}

// Método del electrónico
func (e Electronico) MostrarInfo() string {
	return e.ProductoBase.MostrarInfo() +
		fmt.Sprintln("Garantía:", e.GarantiaMeses, "meses")
}

// Estructura Alimento
type Alimento struct {
	ProductoBase
	FechaExpiracion string
}

// Método del alimento
func (a Alimento) MostrarInfo() string {
	return a.ProductoBase.MostrarInfo() +
		fmt.Sprintln("Fecha de expiración:", a.FechaExpiracion)
}

// Constructor de electrónicos
func NuevoProductoElectronico(id int, nombre string, precio float64, cantidad int, garantia int) Electronico {
	return Electronico{
		ProductoBase:  ProductoBase{id, nombre, precio, cantidad},
		GarantiaMeses: garantia,
	}
}

// Constructor de alimentos
func NuevoProductoAlimento(id int, nombre string, precio float64, cantidad int, fechaExp string) Alimento {
	return Alimento{
		ProductoBase:    ProductoBase{id, nombre, precio, cantidad},
		FechaExpiracion: fechaExp,
	}
}

// Estructura Inventario
type Inventario struct {
	Productos []InformacionProducto
}

// Agregar productos
func (i *Inventario) AgregarProducto(p InformacionProducto) {
	i.Productos = append(i.Productos, p)
}

// Listar productos
func (i Inventario) ListarProductos() {
	fmt.Println("INVENTARIO DE PRODUCTOS")

	for _, p := range i.Productos {
		fmt.Println(p.MostrarInfo())
	}
}

func main() {

	inventario := Inventario{}

	laptop := NuevoProductoElectronico(1, "Laptop", 800, 5, 24)
	mouse := NuevoProductoElectronico(2, "Mouse", 15, 10, 12)

	leche := NuevoProductoAlimento(3, "Leche", 1.50, 20, "15/12/2026")
	yogurt := NuevoProductoAlimento(4, "Yogurt", 0.80, 15, "20/12/2026")

	inventario.AgregarProducto(laptop)
	inventario.AgregarProducto(mouse)
	inventario.AgregarProducto(leche)
	inventario.AgregarProducto(yogurt)

	inventario.ListarProductos()
}
