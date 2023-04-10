package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)

	for i := 0; i < 10; i++ {
		go f(i)
	}
	fmt.Scanln(&input)
}
