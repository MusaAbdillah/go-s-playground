// TL;DR:
// The type person is also type human
// The type student is also type human
package main

import(
	"fmt"
)

type person struct {
    name string
    age  int
}

type student struct {
    alumni person
    school string
}

func (p person) sayHello() string {
    return "Hello, my name is " + p.name
}

func (s student) sayHello() string {
    return "Hello, my name is " + s.alumni.name
}

// Any type that has a sayHello() function will also be type human
// This means that person and student are also type human
type human interface {
    sayHello() string
}

func humanCanPass(h human) {
    fmt.Println("Hey! I'm a human", h)
}

func main() {
    p := person{
        name: "John Doe",
        age:  40,
    }
    s := student{
        alumni: person{
            name: "Ana Doe",
            age:  30,
        },
        school: "MIT",
    }
    fmt.Println(p.sayHello())
    fmt.Println(s.sayHello())

    humanCanPass(p)
    humanCanPass(s)
}