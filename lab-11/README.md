CUSTOM ERROR HANDLING
=====================

We saw with our divide in a previous lab that we could send back an error if something went wrong. But what about if multiple things could go wrong and some of those we could handle? In that case what if we checked what the error was in more detail and then see if we could fix this going forward. So lets extend our Divide function with a few more errors. We know it's bad to divide by 0, but what's the point of dividing by 1 as you will get the same number back? What about if the number doesn't neatly divide? Lets start making custom errors for each of these. Lets create ourselves some named errors globally


```
var (
	ErrDivByZero      = errors.New("attempt to divide by zero")
	ErrDivByOne       = errors.New("are you sure you want to divide by one")
	ErrRemainderFound = errors.New("the was a remainder")
)
```

We now have 3 errors, divide by zero, divide by one and remainder found. We can pass them around by calling them directly, for example to return an error that say you tried to divide by zero you can do

```
return ErrDivByZero
```

We can check in detail the type of our error by comparing it to our existing ones. We can do this via if or a switch statement. 

```
switch err {
	case ErrDivByZero:
	...
	case ErrDivByOne:
	...
	case ErrRemainderFound:
	...
	default:
	 doSomething()
}
```

switch statements can make things easier to read by the are the same essentially as if else chains. The above is exactly the same as

```
if err == ErrDivByZero{
	...
} else if err == ErrDivByOne{
	...
} else if err == ErrRemainderFound{
	...
} else {
	 doSomething()
}

LAB TASK
--------

Write some tests in the `TestDivide()` function in the `main_test.go` file and prove that the divide function is working as expected. The following is our desired behaviour that we should test:

If the denominator is 0, then we should get an ErrDivByZero error back
If the denominator is 1, then we should get an ErrDivByOne error back
If the numbers do not neatly divide, then we should get an ErrRemainderFound error back
If we get an ErrDivByZero error, the result should be 0
If we get an ErrDivByOne or ErrRemainderFound error, the result should be the correct division
If we do not get an error, the result should be the correct division