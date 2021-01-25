package main

import(
	"fmt"
)

func main() {
    x := 21
    fmt.Println(x) // 21
    fmt.Println(&x) // 0xc0000b6020 (of course this will vary on every machine)

    y := &x
    fmt.Println(y) // 0xc0000b6020
    fmt.Println(*y) // 21

    // This is valid too
    fmt.Println(*&x) // 21
}