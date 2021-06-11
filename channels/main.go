package main

//channels para organizar gorutines
import "fmt"

func say(text string, c chan<- string) { // canal de entrada de datos
	c <- text
}

func printChannelOutput(c <-chan string) { // canal de salida de datos
	fmt.Println(<-c)
	//var output string
	//output = <-c
	//fmt.Println(output)
}

func main() {
	// c := make(chan string) // dinamically accepts goroutines
	c := make(chan string, 1) // one goroutine at time
	fmt.Println("Hello")
	go say("Bye", c)

	//fmt.Println(<-c)
	printChannelOutput(c)
}
