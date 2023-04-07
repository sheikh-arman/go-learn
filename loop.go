package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	fmt.Fprintf(out, "Tut\n")

	var x=make([]int,6);
	for _, v:= range x {
		fmt.Fprintf(out,"%v\n",v)
	}
	out.Flush()
}
