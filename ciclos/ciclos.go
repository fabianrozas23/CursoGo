package main

import "fmt"

func main() {
	fmt.Println("Ciclo For")
	forCondicional()
	fmt.Println("Ciclo for while")
	forWhile()
	//fmt.Println("Ciclo for forever")
	//forForever()
	fmt.Println("Ciclo for decremento")
	forDecremento()
	fmt.Println("For con Lista")
	forLista()
}

func forCondicional() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}
}

func forWhile() {
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
}

/*func forForever() {
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
*/

func forDecremento() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}

func forLista() {
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}
}
