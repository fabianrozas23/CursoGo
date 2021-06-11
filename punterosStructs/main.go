package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

//funcion que utiliza al struct
func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	a := 50
	b := &a

	//El símbolo " & " accede a la memoria de la variable.
	//El símbolo " * " accede al valor que está guardado en memoria, usando la dirección de memoria de la variable.
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	//al modificar " *b " el valor de a tambien cambia
	//ya que los dos apuntan a la misma direccion de memoria
	*b = 100
	fmt.Println(a)

	//Llamar a struct
	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)
	//Llamar a funcion ping
	myPC.ping()

	//imprimir estado actual
	fmt.Println(myPC)
	//duplicamos la ram
	myPC.duplicateRAM()
	//imprimimos nuevamente
	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)
}
