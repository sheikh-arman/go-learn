package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewReader(os.Stdout)
)
func f() (int, int) {
	fmt.Println("")
	fmt.Println("slapuj")
	return 5, 6
}
func main() {
	fmt.Print("Hi\n")

	var slice1 = []int{1, 2, 3}
	var slice2 = make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
	var x = make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	fmt.Println(len(x))
	delete(x, "key")
	fmt.Println(len(x))

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["H"] = "Helium"
	elements["li"] = "lithum"
	fmt.Println(elements["li"])

	name, ok := elements["Un"]
	fmt.Println(name, ok)

	xx, yy := f()
	fmt.Println(xx,yy)
	
}
