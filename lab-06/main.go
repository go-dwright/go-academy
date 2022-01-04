package main

// we import the fmt & math package so we can use any public functions within it
import (
	"fmt"
)

func main() {
	// to use a package function we have to first call the pacakge name,
	// in this case "fmt", then it is followed by the function name
	fmt.Println("Hello World!")

	// fmt gives us more power to write strings than the inbuilt function
	// fmt has a large amount of format chracters: %d in this case is for
	// printing integer arguments
	fmt.Printf("Now you have %d problems, but go ain't %d.\n", 99, 1)
}
