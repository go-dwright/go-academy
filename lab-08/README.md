TESTS
=====

Tests! How do we know our code does what it is supposed to do? Well the best way of doing that is by testing it! In principle we should always be working in whats called TDD or Test Driving Development, to do this we write our tests first and then write our code. Why, so that we can accurate describe what our code should do before we even write it, and therefore prove once we have written it that is does what we want.

WRITING TESTS
-------------

Tests a writen in dedicated files separate from the core code of our package. We create a file with the suffix `_test.go` to tell go that this is a file that contains tests. In this lab we will be using `main_test.go` to store our tests.

Go comes with a good testing suite builtin to it. To use this we need to first import the testing package. 

```
import "testing"
```

After that there are a few ways we can define a test in go, but lets start simple. 

```
func TestMyFunction(t *testing.T) {
  ...
}
```

We create a function that is public and starts with the keyword Test, then is has an arguement of *testing.T and returns nothing, in general you would have Test then follow it with the name of the function or functuality you want to test. If we define a function like that then go knows that it is a test and then can be run using the `test` subcommand. 

```
> go test . -v 1
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestSubtract
--- PASS: TestSubtract (0.00s)
=== RUN   TestAddSubtract
--- PASS: TestAddSubtract (0.00s)
=== RUN   TestPower
    main_test.go:96: 3^2: we expected to get 9, we got 6
    main_test.go:105: 4^3: we expected to get 64, we got 12
    main_test.go:114: 5^7: we expected to get 78125, we got 35
--- FAIL: TestPower (0.00s)
FAIL
FAIL    gitlab.platform-engineering.com/golang-academy/lab-08   0.122s
FAIL
```

And now we see we have run all our tests. Looks like some have passed and some have failed with a message. 


LAB TASK
--------

We have tested the Add and Subtract functions and they looks like they are working exactly as expected. Power does not seem to be working and we don't have any tests for Multiply. 

First: look at the logic for the Power function and fix it so that the tests pass, do not edit the tests. 

Second: create a new TestMultiply function that tests the Multiply function and write several tests in the same way as the other example tests. Then try running them, the tests should fail for now and thats a good thing!

Finally: Write the code in the Multiply function so that is passes your tests. 




