package main

import (
	"errors"
	"fmt"
)

func main() {

	// lets run normally for now
	numerator := 10
	denominator := 5
	result, err := Divide(numerator, denominator)
	if err != nil {
		fmt.Printf("unable to divide %d by %d, err: %s\n", numerator, denominator, err)
	} else {
		fmt.Printf("%d divided by %d is %d\n", numerator, denominator, result)
	}

	// what happens when we do try to divide by zero now?
	numerator = 4
	denominator = 0
	result, err = Divide(numerator, denominator)
	if err != nil {
		fmt.Printf("unable to divide %d by %d, err: %s\n", numerator, denominator, err)
	} else {
		fmt.Printf("%d divided by %d is %d\n", numerator, denominator, result)
	}
}

func Divide(numerator, denominator int) (result int, err error) {
	if denominator == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return numerator / denominator, nil
}
