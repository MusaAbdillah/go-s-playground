package main

import "fmt"

func main() {
	variadicExample("red", "blue", "green", "yellow")
}

func variadicExample(s ...string) {
	for i, ss := range s {
		fmt.Println(i, ss)
	}
	
	fmt.Println(s[0])
	fmt.Println(s[3])
	
	
}
