package main

import (
	"fmt"
	"time"
)

func main() {

	hello := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		hello <- "Hello, World!"
	}() //funciona de forma preemptiva

	select {
	case result := <-hello:
		fmt.Println(result)
	default:
		fmt.Println("não foi dessa vez")
	}

	time.Sleep(time.Second * 5)
}
