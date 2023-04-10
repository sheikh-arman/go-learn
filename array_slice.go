package main
import(
	"fmt"
)
func do(a [3]int, b []int) []int{
	//a=b;
	a[0]=4;
	b[0]=3
	c:=make([]int,5)
	c[4]=42
	copy(c,b)
	return c
}
 func main(){
	var w=[...]int{1,2,3}
	var x=[]int{0,0,0}
	
	y:=do(w,x)
	fmt.Println(w,x,y)
 }