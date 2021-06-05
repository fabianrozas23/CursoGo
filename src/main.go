package main

import (
	"fmt"
	"math"
)

func main() {
	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	//División
	result = y / x
	fmt.Println("División:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incrementar
	x++
	fmt.Println("Incremental:", x)

	//Decremento
	x--
	fmt.Println("Decremental:", x)

	//Retos Área
	//Rectángulo, trapecio y de un Círculo

	//Área rectángulo
	const baseRectangulo = 18
	const alturaRectangulo = 12
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Área rectángulo:", areaRectangulo)

	//Área trapecio
	const baseMenorTrapecio = 3
	const baseMayorTrapecio = 5
	const alturaTrapecio = 4
	areaTrapecio := ((baseMayorTrapecio + baseMenorTrapecio) / 2) * alturaTrapecio
	fmt.Println("Área trapecio:", areaTrapecio)

	//Área círculo
	const radioCirculo = 5
	areaCirculo := math.Pi * math.Pow(radioCirculo, 2)
	fmt.Println("Área circulo:", areaCirculo)
}
