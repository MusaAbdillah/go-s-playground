package main

import (
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

func main() {
    x := student{
        alumni: person{
            name: "John Doe",
            age:  40,
        },
        school: "MIT",
    }
    fmt.Println(x.school)
    fmt.Println(x.alumni.name, x.alumni.age)
}
