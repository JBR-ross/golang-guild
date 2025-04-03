package main

import (
	"fmt"
	"runtime"
	"time"
)

// usar o 1.13 para exemplificar

func main() {
	runtime.GOMAXPROCS(1) // Set the number of OS threads to 1
	fmt.Println("Começou")
	//forma cooperativa

	go func() { //forma preemptiva
		for {

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Fim")
}
