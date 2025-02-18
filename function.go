package main

import(
	"fmt"
)

type person struct {
    name string
    age  int
}
// This function will be attached to any variable whose type is person
func (p person) sayHello() string {
    return "Hello, my name is " + p.name
}

func main() {
    x := person{
        name: "John Doe",
        age:  40,
    }
    y := person{
        name: "Ana Doe",
        age:  40,
    }
    fmt.Println(x.sayHello()) // Hello, my name is John Doe
    fmt.Println(y.sayHello()) // Hello, my name is Ana Doe
}
