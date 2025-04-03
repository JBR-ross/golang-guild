package main

import "fmt"

func contador() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	contador()
	fmt.Println("Hello, World!1")
	fmt.Println("Hello, World!2")
}
