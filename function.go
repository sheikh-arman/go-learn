package main
import(
	"fmt"
)
func avarage(xs []float64)float64{
	total:=0.0
	for _, v:=range xs{
		total+=v
	}
	return total/float64(len(xs))
}
func f()(int,int){
	return 5,6
}

func add(args ...int)int{
	total:=0
	for _,v:=range args{
		total+=v
	}
	return total
}
func main(){
	xs:=[]float64{98,93,77,82,83}
	fmt.Println(avarage(xs))
	x,y:=f()
	fmt.Printf("%v %v\n",x,y)
	fmt.Printf("%v\n",add(1,2,3))

	a:=0

	increment:=func() int{
		a++
		return a
	}
	fmt.Printf("%v\n",increment())
	fmt.Printf("%v\n",increment())
}