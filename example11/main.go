package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			queue <- i //lembra do mutex? esse queue é um chanal, então a thread vai ficar aguardando até ele ser esvaziado
			i++
		}
	}()

	for x := range queue {
		fmt.Println(x)
	}
}
