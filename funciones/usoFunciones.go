package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	/*	 si solo quieres ocupar un valor
	_, value2 := doubleReturn(2)
	fmt.Println("Value2:", value2)*/
	value1, value2 := doubleReturn(2)
	fmt.Printf("Value1: %v - Value2: %v\n", value1, value2)
}
