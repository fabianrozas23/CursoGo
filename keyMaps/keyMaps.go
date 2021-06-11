package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar valor
	value := m["Jose"]
	fmt.Println(value)

	//Ver si no esta en el diccionario
	value, ok := m["Josep"]
	fmt.Println(value, ok)

	//eliminar elemento
	delete(m, "Jose")
	v, ok := m["Jose"]
	fmt.Println(v, ok)
}
