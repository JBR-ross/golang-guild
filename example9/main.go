package main

import "fmt"

func main() {

	forever := make(chan string)

	go func() {
		x := true //da pra remover
		for {
			if x == true { //da pra remover
				continue
			}
		}
	}() //funciona de forma preemptiva

	fmt.Println("Aguardando para sempre!")

	<-forever // bloqueia o programa, ou seja, a nossa chanel para a execução até ela receber um valor.
}
