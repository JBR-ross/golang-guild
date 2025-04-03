package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Printf("Worker %d: %d\n", workerId, res)
		time.Sleep(time.Second)
	}
}

//Apache é comuim colocar só 5 para rodar em um webserver, no go o ceu é o limite

func main() {

	msg := make(chan int)
	go worker(1, msg) //lembrar de enfatizar o go routine

	for i := 0; i < 10; i++ {
		msg <- i // faz sentido para voce que uma API tenha que esperar para enviar uma mensagem?
	}

}
