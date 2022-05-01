ERRORS
======

We don't really want to panic when we find an issue, so instead we should use an error to tell us something is wrong. In go, errors are a type like anything else so when we call a function that could fail for some reason, we generally return the thing that we want as well as an error. For example we saw we paniced when we tried to divide by zero in a pervious lab. To get around this if we know there is a potential for something to go wrong, then we should also return an error. For example

```
func Divide(numerator, denominator int) (result int, err error) {
	if denominator == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return numerator / denominator, nil
}
```

We now always get back 2 values from when we call the Divide function, a result of type int and an err of type error. So to call it could do the following to capture both of them (be aware we have to let go know that we are getting all of the returned values)

```
numerator := 10
denominator := 5
result, err := Divide(numerator, denominator)
```

Okay, so now we have a result and an error, but how do we know if there was a problem? We do that by checking if the err is nil (golangs version of null). If it is nil then no error was thrown and we are good! If it is not nil then there is a problem and we can't carry on. We can print out errors like a string to see what the problem was. 

```
if err != nil {
		fmt.Printf("unable to divide %d by %d, err: %s\n", numerator, denominator, err)
	}
```

LAB TASK
--------

Write some tests in the `TestDivide()` function in the `main_test.go` file and prove that the divide function is working as expected. The following is our desired behaviour that we should test:

If the denominator is 0, then we should get an error back
If the denominator is not 0, then we should not get an error back
If we get an error, the result should be 0
If we do not get an error, the result should be the correct division