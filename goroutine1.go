package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine") // Not Executing
}
func main() {
	go hello()
	fmt.Println("main function")
}
