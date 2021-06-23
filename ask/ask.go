package main

import (
	"encoding/json"
	"fmt"

	"github.com/simonnilsson/ask"
)

func main() {

	// Use parsed JSON as source data
	var object map[string]interface{}
	json.Unmarshal([]byte(`{ "a": [{ "b": { "c": 3, "d": "D value" } }] }`), &object)

	// Extract the 3
	res, ok := ask.For(object, "a[0].b.c").Int(0)
	fmt.Println("--- res ---")
	fmt.Println(res, ok)
	// Output: 3 true

	// Attempt extracting a string at path .d that does not exist
	res2, ok := ask.For(object, "a[0].b.d").String("default nil")

	fmt.Println("--- res2 ---")
	fmt.Println(res2, ok)
	// Output: nil false

	res3, ok := ask.For(object, "a[0]").Int(0)
	fmt.Println("--- res3 ---")
	fmt.Println(res3)
}
