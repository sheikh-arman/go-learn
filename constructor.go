package main

import (
	"fmt"
)

type Rectangle struct {
	Name          string
	Width, Height float64
}

func main() {
	var a Rectangle
	var b = Rectangle{"Im Arman", 12.0, 20.0}
	var c = Rectangle{Height: 12, Width: 14}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
