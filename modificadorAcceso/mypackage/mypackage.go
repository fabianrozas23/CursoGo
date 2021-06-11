package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

func PrintMessage(text string) {
	fmt.Println(text)
	//fmt.Println("Hola")
}
