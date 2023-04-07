package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
var(
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)
type Pair struct{
	First int
	Second int
}
func main(){
	defer func() {
		out.Flush()
	}()
	var tcase=0
	fmt.Fscan(in,&tcase)
	var test=0;
	for test<tcase {
		var (
			n=0
			s1=0
			s2=0
		)
		fmt.Fscan(in,&n,&s1,&s2)
		var V1,V2 =[]int{},[]int{}
		V := make([]Pair, n)
		for i:=0; i<n; i++ {
			var a int=0
			fmt.Fscan(in,&a)
			V[i].First=a
			V[i].Second=i+1
		}
		sort.Slice(V,func(i,j int) bool{
			return V[i].First>V[j].First
		}) 
		sum1:=s1
		sum2:=s2
		for i:=0; i<n;i++ {
			if sum1<sum2 {
				V1=append(V1,V[i].Second)
				sum1+=s1
			}else{
				V2=append(V2,V[i].Second)
				sum2+=s2
			}
		}

		fmt.Fprintf(out,"%v",len(V1))
		for i:=0; i<len(V1);i++ {
			fmt.Fprintf(out," %v",V1[i])
		}
		fmt.Fprintf(out,"\n")
		fmt.Fprintf(out,"%v",len(V2))
		for i:=0 ;i<len(V2);i++ {
			fmt.Fprintf(out," %v",V2[i])
		}
		fmt.Fprintf(out,"\n")
		test+=1
	}
}