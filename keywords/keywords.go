package main

import "fmt"

func main() {
	//Defer
	defer fmt.Println("Me muestra de los ultimos")
	fmt.Println("Hola voy primero")
	defer fmt.Println("Me muestra antes del ultimo")
	fmt.Println("Voy segundo")

	//continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		//Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
