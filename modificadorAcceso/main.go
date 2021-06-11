package main

import (
	pk "curso_golang/modificadorAcceso/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hola mundo")
	//pk.PrintMessage()
}
