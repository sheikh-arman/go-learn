package main

import (  
    "fmt"
    "reflect"
)

type order struct {  
    ordId      int
    customerId int
}
type employee struct {  
    name string
    id int
    address string
    salary int
    country string
}
func createQuery(q interface{}) {  
    t := reflect.TypeOf(q)
    v := reflect.ValueOf(q)
    knd := t.Kind()
    fmt.Println("Type ", t)
    fmt.Println("Value ", v)
    fmt.Println("Kind ", knd)


}
func main() {  
    o := order{
        ordId:      456,
        customerId: 56,
    }
	e := employee {
        name: "Naveen",
        id: 565,
        address: "Science Park Road, Singapore",
        salary: 90000,
        country: "Singapore",
    }
    createQuery(o)
    createQuery(e)

}