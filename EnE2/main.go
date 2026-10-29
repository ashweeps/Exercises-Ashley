package main

import (
	"errors"
	"fmt"
)

type Estudiante struct {
	edad int
}

func (e Estudiante) GetEdad() int {
	return e.edad
}
func (e *Estudiante) SetEdad(edad int) error {
	if edad < 0 || edad > 80 {
		return errors.New("Error, la edad debe de estar entre 1 y 79 años")
	}
	e.edad = edad
	return nil
}
func NewStudent(edad int) (*Estudiante, error) {
	est := &Estudiante{}
	err := est.SetEdad(edad)
	if err != nil {
		return nil, err
	}
	return est, nil
}

func main() {
	student := Estudiante{}
	err := student.SetEdad(100)
	if err != nil {
		fmt.Println("No se puede modificar", err)
	}
	fmt.Println("Hola, su edad es: ", student.GetEdad())

}
