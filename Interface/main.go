package main

import (
	"errors"
	"fmt"
)

type Dispositivo interface {
	Encender() error
	Apagar() error
	Estatus() string
}

type Router struct {
	Nombre    string
	Encendido bool
}

func (r *Router) Encender() error {
	if r.Encendido {
		return errors.New("El router ya está encendido")
	}
	r.Encendido = true
	return nil
}

func (r *Router) Apagar() error {
	if !r.Encendido {
		return errors.New("El router ya está apagado")
	}
	r.Encendido = false
	return nil
}

func (r *Router) Estatus() string {
	if r.Encendido {
		return r.Nombre + ": está encendido"
	}
	return r.Nombre + ": está apagado"
}

type Servidor struct {
	Nombre string
	Activo bool
}

func (s *Servidor) Encender() error {
	if s.Activo {
		return errors.New("El servidor ya está activo")
	}
	s.Activo = true
	return nil
}

func (s *Servidor) Apagar() error {
	if !s.Activo {
		return errors.New("El servidor ya está apagado")
	}
	s.Activo = false
	return nil
}

func (s *Servidor) Estatus() string {
	if s.Activo {
		return s.Nombre + ": está activo"
	}
	return s.Nombre + ": está inactivo"
}

func OperarDispositivo(d Dispositivo) {
	fmt.Println("Estatus inicial:", d.Estatus())

	err := d.Encender()
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println("Después de encender:", d.Estatus())

	err = d.Apagar()
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println("Después de apagar:", d.Estatus())
}

func main() {
	router := &Router{
		Nombre:    "Cisco",
		Encendido: true,
	}

	servidor := &Servidor{
		Nombre: "Servidor Web",
		Activo: true,
	}

	fmt.Println("***** Router *****")
	OperarDispositivo(router)

	fmt.Println("***** Servidor *****")
	OperarDispositivo(servidor)
}
