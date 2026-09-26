package vocales

func Contar(frase string) (int, int, int, int, int) {

	a := 0
	e := 0
	i := 0
	o := 0
	u := 0

	for _, letra := range frase {

		switch letra {

		case 'a', 'A', 'á':
			a++

		case 'e', 'E', 'é':
			e++

		case 'i', 'I', 'í':
			i++

		case 'o', 'O', 'ó':
			o++

		case 'u', 'U', 'ú', 'ü':
			u++
		}
	}

	return a, e, i, o, u
}
