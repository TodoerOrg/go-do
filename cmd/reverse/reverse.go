package main

import "fmt"
import "github.com/todoer-org/todoer/src/lib"

func main() {
	result := lib.Reverse("abc")
	fmt.Printf("result: %s \n", result)
}
