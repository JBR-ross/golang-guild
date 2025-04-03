package main

import "fmt"

// channel serve para uma thread, se comunicar com a outra
// não compartilha memória para se comunicar, ele se comunica compartilhando memória!!

func main() {
	hello := make(chan string) // channel é um tipo de dado

	go func() { // função anônima
		hello <- "Hello, World!" // envia a string para o canal
	}()
	result := <-hello   // recebe a string do canal
	fmt.Println(result) // imprime a string
}
