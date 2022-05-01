PANIC!
======

A panic is where our code hits a situation where it just cannot continue. For example the famous divide by zero. Here is some code to demonstrate this (you may notice we had to put the 0 in a variable because go would refuse to compile if we explitely put a divide by zero in our code)

```
package main

func main() {
	divisor := 0
	total := 5 / divisor
	println(total)
}
```

when I try to run it, my code panics and I get this as a response

```
> go run .
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.main()
        /Users/damonwright/go/src/gitlab.platform-engineering.com/golang-academy/lab-09/main.go:5 +0x11
exit status 2
```

What does this tell us? Lets break it down into sections. 

`panic: runtime error: integer divide by zero` We are told our code paniced due to a runtime error that was caused by an integer divide by zero

`goroutine 1 [running]:` This told us it ran in goroutine 1 which means the main logic and not a separate goroutine (like a thread but we will get to those later)

`main.main()` The panic happend in the main() function in the main package

`/Users/damonwright/go/src/gitlab.platform-engineering.com/golang-academy/lab-09/main.go:5 +0x11` the panic happened on line 5 of our code. 


`exit status 2` our code exited with a return code of 2 (not good)

WHATS THE POINT?
----------------

panicking is there if we know that there is a situation that we cannot continue on so we may as well exit now rather than continuing. Best practise says our code should never really crash and we should handle erros, but that's something else we will get onto shortly. We can also make our code panic whenever we like by calling the `panic()` builtin function. 


LAB TASK
--------

Run the code that is a part of this lab and have a look at the output. Make sure you understand all the information that it is telling you and how you can trace through the output of the panic to see the route the code used to get to the panic.