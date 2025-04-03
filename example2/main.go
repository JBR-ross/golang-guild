package main

import "fmt"

func contador(tipo string) {
	for i := 0; i < 10; i++ {
		fmt.Println(tipo, i)
	}
}

func main() {
	contador("Sem go routines")
	fmt.Println("Hello, World!1")
	fmt.Println("Hello, World!2")
}
