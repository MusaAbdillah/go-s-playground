package main 

import (
	"fmt"
)

func main() {
	string := "PanicParse"

	// make it panic
	tooFar := string[0:20]
	fmt.Println(tooFar)
	fmt.Println("PanicParse")
}