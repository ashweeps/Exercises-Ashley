package conversor

func Convertir(dolares float64, moneda string) float64 {

	var resultado float64

	switch moneda {

	case "EUR":
		resultado = dolares * 0.85

	case "LB":
		resultado = dolares * 0.74

	case "WON":
		resultado = dolares * 1400

	case "BTC":
		resultado = dolares / 110000

	default:
		resultado = 0
	}

	return resultado
}
