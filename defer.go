package main
import(
	"fmt"
)
func first(){
	fmt.Printf("1st\n")
}
func second(){
	fmt.Printf("2nd\n")
}
func main(){
	defer second()
	first()
}