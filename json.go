package main

import (
    "encoding/json"
    "fmt"
)
// Keys must be capitalized
type person struct {
    Name     string
    LastName string
    Age      int
}

func main() {
    p1 := person{
        Name:     "John",
        LastName: "Doe",
        Age:      40,
    }
    p2 := person{
        Name:     "Ana",
        LastName: "Doe",
        Age:      32,
    }
    people := []person{p1, p2}

    bs, err := json.Marshal(people)

    if err != nil {
        fmt.Println("The error => ", err)
    }

    fmt.Println("--- []people{p1, p2}")   
    fmt.Println(people)

    fmt.Println("--- json.Marshal Result bs")
    fmt.Println(bs)

    fmt.Println("--- string convert json.Marshal Result bs")
    fmt.Println(string(bs)) // [{"Name":"John","LastName":"D....
}