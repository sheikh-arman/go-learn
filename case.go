package main

import (
	"fmt"
	"bufio"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewReader(os.Stdout)
)

func main() {
	var num=3
	switch num {
	case 1:
		fmt.Printf("One\n")
		fallthrough
	case 2:
		fmt.Printf("Two\n")
		fallthrough
	case 3:
		fmt.Printf("Three\n")
	}
}
