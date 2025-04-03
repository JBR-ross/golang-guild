package main

import "fmt"

func contador(tipo string) {
	for i := 0; i < 10; i++ {
		fmt.Println(tipo, i)
	}
}

func main() {
	contador("Sem go routines")
	go contador("Com go routines")
}
