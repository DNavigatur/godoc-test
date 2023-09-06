package main

import (
	"fmt"

	"MyModule/HenMath"
)

// we are calling the function from a different package
func main() {
	fmt.Println(HenMath.Add(6, 9))
}
