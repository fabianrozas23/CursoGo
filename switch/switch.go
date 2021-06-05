package main

import "fmt"

func main() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	switch mod := 4 % 2; mod {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	// Sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor que 100")
	default:
		fmt.Println("Es menor que 100")
	}

	val := 50
	switch {
	case val > 100:
		fmt.Println("Es mayor que 100")
	case val < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("Es menor que 100 y mayor que 0")
	}

}
