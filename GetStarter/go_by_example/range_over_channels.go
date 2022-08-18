package main

import "fmt"

// esse exemplo mostra que mesmo com o canal fechado, o range pode ser usado para ler os elementos do canal.

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
