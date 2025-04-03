package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 10; i++ {
		fmt.Println(tipo, i)
	}
}

func main() {
	contador("Sem go routines")
	go contador("Com go routines")

	time.Sleep(time.Second * 2)
	fmt.Println("Fim")
}
