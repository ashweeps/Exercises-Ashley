package main

import "fmt"

type Persona struct {
	nombre string
}

func (P Persona) GetNombre() string {
	return (P.nombre)
}
func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func main() {
	persona := Persona{}

	persona.SetNombre("Ashley")
	fmt.Println("Su nombre es: ", persona.GetNombre())

	persona.SetNombre("Michelle")
	fmt.Println("El otro nombre es: ", persona.GetNombre())

}
