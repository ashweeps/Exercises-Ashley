package vocales

func Contar(frase string) (int, int, int, int, int) {

	a := 0
	e := 0
	i := 0
	o := 0
	u := 0

	for _, letra := range frase {

		switch letra {

		case 'a', 'A':
			a++

		case 'e', 'E':
			e++

		case 'i', 'I':
			i++

		case 'o', 'O':
			o++

		case 'u', 'U':
			u++
		}
	}

	return a, e, i, o, u
}
