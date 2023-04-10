package main

import(
	"fmt"
)

func main(){
	var m map[string]int
	p:=make(map[string]int)

	a:= p["the"]
	b:=m["the"]

	//m["and"]=1
	m=p
	m["and"]++
	c:=p["and"]
	fmt.Print(a,b,c)
}