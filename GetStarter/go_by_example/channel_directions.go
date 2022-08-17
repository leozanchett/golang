package main

import "fmt"

// Ao usar canais como parametros de funcao, voce pode especificar se um canal
// destina-se apenas a enviar ou receber valores. Essa especificidade aumenta a seguranca de tipo do programa.

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
