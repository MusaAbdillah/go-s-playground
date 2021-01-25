package main

import (
    "encoding/json"
    "fmt"
)

type person struct {
    Name     string
    LastName string
    Age      int
}

func main() {
    jsonString := `[{"Name":"John","LastName":"Doe","Age":40},{"Name":"Ana","LastName":"Doe","Age":32}]`
    jsonBytes := []byte(jsonString)
    people := []person{}

    err := json.Unmarshal(jsonBytes, &people)

    if err != nil {
        fmt.Println("The error => ", err)
    }

    fmt.Println(people[0].Name) // John
    fmt.Println(people[0].LastName) // Doe
    fmt.Println(people[1].Name) // Ana
}