package main

import "fmt"

func main() {

	var n int

	fmt.Println("Ingresa un numero entero positivo:")
	fmt.Scan(&n)

	fmt.Println("*****FizzBuzz*****")

	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	switch {
	case n >= 1 && n <= 10:
		fmt.Println("Numero pequeño")

	case n >= 11 && n <= 100:
		fmt.Println("Numero mediano")

	case n > 100:
		fmt.Println("Numero grande")

	}
}
