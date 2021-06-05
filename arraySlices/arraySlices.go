package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	fmt.Println(array, len(array), cap(array))

	//Array
	x := [5]int{0, 1, 2, 3, 4}
	fmt.Println(x)

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Slice
	l := make([]float64, 5)
	fmt.Println(l)

	//Slice
	n := make([]float64, 5, 10)
	fmt.Println("n: ", n)

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append
	//agrega un numero al slice ej: el 10
	slice = append(slice, 10)
	fmt.Println(slice)

	//Append
	//nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
