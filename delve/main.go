package main

import(
	"fmt"
	"time"
)

type Person struct{
	Name string
}

func main(){
	fmt.Println("Delve")
	go sayHello()
    time.Sleep(1 * time.Second)
	fmt.Println("... Waiting for go routine")
}

func sayHello(){
	p := Person{Name: "Musa"}
	fmt.Println("Hello ", p.Name)
}
