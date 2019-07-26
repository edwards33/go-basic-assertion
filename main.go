package main

import "fmt"

func main() {
	var name interface{} = "NY"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("it is not a string\n")
	}
}
