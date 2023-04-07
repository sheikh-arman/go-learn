package main

import (
	"fmt"
)

func even() func() int {
	a := int(0)
	return func() int {
		a += 2
		return a
	}
}
func main() {
	a := even()
	for i := 0; i <= 10; i++ {
		fmt.Println(a())
	}
}
