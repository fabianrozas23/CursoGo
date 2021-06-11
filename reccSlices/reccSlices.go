package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palíndromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	//ama
	//amor a roma
	var palabra string
	fmt.Scan(&palabra)
	isPalindromo(palabra)
	//	minus := strings.ToLower(palabra)
	//	isPalindromo(minus)
	/*
		isPalindromo("ama")
		isPalindromo("amor a roma")
		isPalindromo("casas")*/
}
