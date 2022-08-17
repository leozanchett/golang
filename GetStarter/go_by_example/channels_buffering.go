package main

import "fmt"

func main() {
	// os buffers aceitam uma quantidade limitada de inputs
	// os outputs seguem a mesma ordem de inputs
	// ao executar um output, libera um input
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)

	messages <- "teste"
	fmt.Println(<-messages)

	fmt.Println(<-messages)

}
