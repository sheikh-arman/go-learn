package main

import (
	"fmt"
)

type base struct {
	a string
	b int
}

func (this base) xyz() {
	fmt.Println("xyz, a is :", this.a)
}
func (this base) display() {
	fmt.Println("base, a is:", this.a)
}

type derived struct {
	base
	d int
	a float32
}

func (this derived) display() {
	fmt.Println("derived a is :", this.a)
}
func main() {
	var a derived = derived{base{"base-a", 10}, 20, 2.5}
	a.display()

	a.base.display()

	a.xyz()
}
