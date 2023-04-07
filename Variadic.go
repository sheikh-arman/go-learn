package main

import "fmt"

func first() {
  fmt.Println("1st")
}
func second() {
  fmt.Println("2nd")
}
type circle struct{
	x,y,z float64
}
func main() {
	var a circle
	a.x=1
	a.y=2
	a.z=3
	fmt.Println(a)
}