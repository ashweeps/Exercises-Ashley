package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	producto := Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}

	return producto
}

func (producto *Producto) ActualizarPrecio(nuevoPrecio float64) {
	producto.Precio = nuevoPrecio
}

func (producto *Producto) ActualizarCantidad(nuevaCantidad int) {
	producto.Cantidad = nuevaCantidad
}

func (producto Producto) MostrarInfo() {
	fmt.Println("ID:", producto.ID)
	fmt.Println("Nombre:", producto.Nombre)
	fmt.Println("Precio: $", producto.Precio)
	fmt.Println("Cantidad:", producto.Cantidad)
}

type Inventario struct {
	Productos map[int]Producto
}

func (inventario *Inventario) AgregarProducto(producto Producto) {
	inventario.Productos[producto.ID] = producto
}

func (inventario *Inventario) EliminarProducto(id int) {
	delete(inventario.Productos, id)
}

func (inventario Inventario) BuscarProducto(id int) (Producto, bool) {
	producto, existe := inventario.Productos[id]
	return producto, existe
}

func (inventario Inventario) ListarProductos() {

	if len(inventario.Productos) == 0 {
		fmt.Println("No hay productos en el inventario")
	} else {
		fmt.Println("LISTA DE PRODUCTOS")

		for _, producto := range inventario.Productos {
			fmt.Println("--------------------")
			producto.MostrarInfo()
		}
	}
}

func main() {

	inventario := Inventario{
		Productos: make(map[int]Producto),
	}

	var opcion int
	var id int
	var nombre string
	var precio float64
	var cantidad int

	for {
		fmt.Println("INVENTARIO DE TIENDA")
		fmt.Println("1. Agregar producto")
		fmt.Println("2. Mostrar productos")
		fmt.Println("3. Buscar producto")
		fmt.Println("4. Actualizar precio")
		fmt.Println("5. Actualizar cantidad")
		fmt.Println("6. Eliminar producto")
		fmt.Println("7. Salir")
		fmt.Print("Seleccione una opcion: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			fmt.Print("Ingrese el ID: ")
			fmt.Scan(&id)

			_, existe := inventario.BuscarProducto(id)

			if existe {
				fmt.Println("Ya existe un producto con ese ID")
			} else {
				fmt.Print("Ingrese el nombre: ")
				fmt.Scan(&nombre)

				fmt.Print("Ingrese el precio: ")
				fmt.Scan(&precio)

				fmt.Print("Ingrese la cantidad: ")
				fmt.Scan(&cantidad)

				if precio < 0 || cantidad < 0 {
					fmt.Println("El precio y la cantidad no pueden ser negativos")
				} else {
					producto := NuevoProducto(id, nombre, precio, cantidad)
					inventario.AgregarProducto(producto)

					fmt.Println("Producto agregado correctamente")
				}
			}

		case 2:
			inventario.ListarProductos()

		case 3:
			fmt.Print("Ingrese el ID del producto: ")
			fmt.Scan(&id)

			producto, existe := inventario.BuscarProducto(id)

			if existe {
				fmt.Println("Producto encontrado:")
				producto.MostrarInfo()
			} else {
				fmt.Println("Producto no encontrado")
			}

		case 4:
			fmt.Print("Ingrese el ID del producto: ")
			fmt.Scan(&id)

			producto, existe := inventario.BuscarProducto(id)

			if existe {
				fmt.Print("Ingrese el nuevo precio: ")
				fmt.Scan(&precio)

				if precio < 0 {
					fmt.Println("El precio no puede ser negativo")
				} else {
					producto.ActualizarPrecio(precio)
					inventario.AgregarProducto(producto)

					fmt.Println("Precio actualizado correctamente")
				}
			} else {
				fmt.Println("Producto no encontrado")
			}

		case 5:
			fmt.Print("Ingrese el ID del producto: ")
			fmt.Scan(&id)

			producto, existe := inventario.BuscarProducto(id)

			if existe {
				fmt.Print("Ingrese la nueva cantidad: ")
				fmt.Scan(&cantidad)

				if cantidad < 0 {
					fmt.Println("La cantidad no puede ser negativa")
				} else {
					producto.ActualizarCantidad(cantidad)
					inventario.AgregarProducto(producto)

					fmt.Println("Cantidad actualizada correctamente")
				}
			} else {
				fmt.Println("Producto no encontrado")
			}

		case 6:
			fmt.Print("Ingrese el ID del producto: ")
			fmt.Scan(&id)

			_, existe := inventario.BuscarProducto(id)

			if existe {
				inventario.EliminarProducto(id)
				fmt.Println("Producto eliminado correctamente")
			} else {
				fmt.Println("Producto no encontrado")
			}

		case 7:
			fmt.Println("Saliendo del programa")
			return

		default:
			fmt.Println("Opcion incorrecta")
		}

		fmt.Println()
	}
}
