package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivByZero      = errors.New("attempt to divide by zero")
	ErrDivByOne       = errors.New("are you sure you want to divide by one")
	ErrRemainderFound = errors.New("the was a remainder")
)

func main() {

	// lets run normally for now
	numerator := 10
	denominator := 5
	result, err := Divide(numerator, denominator)
	switch err {
	case ErrDivByZero:
		fmt.Printf("unable to divide %d by %d, err: %s\n", numerator, denominator, err)
	case ErrDivByOne:
		fmt.Printf("able to divide %d by %d, but should be have?\n", numerator, denominator)
	case ErrRemainderFound:
		fmt.Printf("%d is not completely divisible by %d\n", numerator, denominator)
	default:
		fmt.Printf("%d divided by %d is %d\n", numerator, denominator, result)
	}

	// what happens when we do try to divide by zero now?
	numerator = 4
	denominator = 0
	result, err = Divide(numerator, denominator)
	switch err {
	case ErrDivByZero:
		fmt.Printf("unable to divide %d by %d, err: %s\n", numerator, denominator, err)
	case ErrDivByOne:
		fmt.Printf("able to divide %d by %d, but should be have?\n", numerator, denominator)
	case ErrRemainderFound:
		fmt.Printf("%d is not completely divisible by %d, got %d\n", numerator, denominator, result)
	default:
		fmt.Printf("%d divided by %d is %d\n", numerator, denominator, result)
	}

	// what happens when we do try to divide a prime number?
	numerator = 23
	denominator = 7
	result, err = Divide(numerator, denominator)
	switch err {
	case ErrDivByZero:
		fmt.Printf("unable to divide %d by %d, err: %s\n", numerator, denominator, err)
	case ErrDivByOne:
		fmt.Printf("able to divide %d by %d and got %d, but should we have?\n", numerator, denominator, result)
	case ErrRemainderFound:
		fmt.Printf("%d is not completely divisible by %d, got %d\n", numerator, denominator, result)
	default:
		fmt.Printf("%d divided by %d is %d\n", numerator, denominator, result)
	}

	// what happens when we do try to divide by 1?
	numerator = 37
	denominator = 1
	result, err = Divide(numerator, denominator)
	switch err {
	case ErrDivByZero:
		fmt.Printf("unable to divide %d by %d, err: %s\n", numerator, denominator, err)
	case ErrDivByOne:
		fmt.Printf("able to divide %d by %d and got %d, but should we have?\n", numerator, denominator, result)
	case ErrRemainderFound:
		fmt.Printf("%d is not completely divisible by %d, got %d\n", numerator, denominator, result)
	default:
		fmt.Printf("%d divided by %d is %d\n", numerator, denominator, result)
	}
}

func Divide(numerator, denominator int) (result int, err error) {
	if denominator == 0 {
		return 0, ErrDivByZero
	}
	if denominator == 1 {
		return numerator / denominator, ErrDivByOne
	}
	remainder := numerator % denominator
	if remainder != 0 {
		return numerator / denominator, ErrRemainderFound
	}
	return numerator / denominator, nil
}
