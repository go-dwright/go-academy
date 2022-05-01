package main

import "testing"

// Lets first test the Add function
func TestAdd(t *testing.T) {
	// we know that 5+6=11 so lets test that
	firstNumber := 5
	secondNumber := 6
	expectedTotal := 11
	total := Add(firstNumber, secondNumber)
	if total != expectedTotal {
		// to fail a test we can call t.Error or t.Errorf to fail this test with a message
		t.Errorf("%d + %d should equal %d, instead got %d", firstNumber, secondNumber, expectedTotal, total)
	}

	// lets try with a couple more numbers
	firstNumber = 1000
	secondNumber = 111
	expectedTotal = 1111
	total = Add(firstNumber, secondNumber)
	if total != expectedTotal {
		t.Errorf("%d + %d should equal %d, instead got %d", firstNumber, secondNumber, expectedTotal, total)
	}

	firstNumber = 999999
	secondNumber = 1
	expectedTotal = 1000000
	total = Add(firstNumber, secondNumber)
	if total != expectedTotal {
		t.Errorf("%d + %d should equal %d, instead got %d", firstNumber, secondNumber, expectedTotal, total)
	}
}

// Let's now try to test subtract
func TestSubtract(t *testing.T) {
	// we know that 11-5=6 so lets test that
	firstNumber := 11
	secondNumber := 5
	expectedTotal := 6
	total := Subtract(firstNumber, secondNumber)
	if total != expectedTotal {
		t.Errorf("%d + %d should equal %d, instead got %d", firstNumber, secondNumber, expectedTotal, total)
	}
}

// Let's now try to test Add & Subtract at the same time
func TestAddSubtract(t *testing.T) {
	// given any starting number, if we add a number to it and then subtract that, we should get back our starting number
	startNumber := 17
	otherNumber := 5

	temp := Add(startNumber, otherNumber)
	finalNumber := Subtract(temp, otherNumber)
	if startNumber != finalNumber {
		t.Errorf("start number did not match final number? %d != %d", startNumber, finalNumber)
	}

	// we can also do the same in reverse
	startNumber = 232
	otherNumber = 11

	temp = Subtract(startNumber, otherNumber)
	finalNumber = Add(temp, otherNumber)
	if startNumber != finalNumber {
		t.Errorf("start number did not match final number? %d != %d", startNumber, finalNumber)
	}

	// now that we have tested that, lets try with some bigger numbers
	startNumber = 364723649
	otherNumber = 834221234

	temp = Add(startNumber, otherNumber)
	finalNumber = Subtract(temp, otherNumber)
	if startNumber != finalNumber {
		t.Errorf("start number did not match final number? %d != %d", startNumber, finalNumber)
	}
}

func TestPower(t *testing.T) {
	// lets start small, we know that 2 to the power 2 is 4
	number := 2
	power := 2
	expected := 4
	total := Power(number, power)
	if total != expected {
		t.Errorf("%d^%d: we expected to get %d, we got %d", number, power, expected, total)
	}

	// lets try bigger, lets say 3 power 2 is 9
	number = 3
	power = 2
	expected = 9
	total = Power(number, power)
	if total != expected {
		t.Errorf("%d^%d: we expected to get %d, we got %d", number, power, expected, total)
	}

	// lets try even bigger, lets say 4 power 3 is 64
	number = 4
	power = 3
	expected = 64
	total = Power(number, power)
	if total != expected {
		t.Errorf("%d^%d: we expected to get %d, we got %d", number, power, expected, total)
	}

	// lets try really really big, lets say 5 power 7 is 9
	number = 5
	power = 7
	expected = 78125
	total = Power(number, power)
	if total != expected {
		t.Errorf("%d^%d: we expected to get %d, we got %d", number, power, expected, total)
	}
}
