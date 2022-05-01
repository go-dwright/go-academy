IMPORTS
=======

There is a very large amount of code out there that can make common tasks a lot easier. To use this code we will need to import this into our code, a very famous exmaple is the `fmt` or the format package.  

IMPORTING A PACKAGE
-------------------

We can import dependent code by using the `import` keyword. Our imports always come at the top of our file immediately after we define our package.  

```
package main

import "fmt"
...
```

After we have imported a package we can use any function within it. 

```
fmt.Println("Hello World!")
```

There are 2 important parts here, first is that we have to prefix the function with the pacakage it comes from in this case fmt. The second part is that the function starts with a capital and is in CamelCase. This is how go differs between private functions, which can only be used inside the code where it is defined, and public functions which can be used after being imported. 

```
package vehicle

func NewCar() Car    // this can be used outside of the package as it starts with a capital
func newBike() bike  // this cannot be used outside the package as it starts with a lower case character
```

We also now have the concept of OOP or object orientated programming. We can create an object and then have specific functions that are tied to that instance of the object. 

```
myCar := NewCar(4) // create a car with 4 wheels
mySecondCar := NewCar(5) // create a car with 5 wheels
fmt.Println("my first car has %d wheels", myCar.GetWheels())
fmt.Println("my second car has %d wheels", mySecondCar.GetWheels())
```

In the example above we created 2 cars each with a different number of wheels. We can then call a function on the object itself rather than calling a function with the car object as an arguement. To do this we add in the object almost like an aguement in the function definition but before we provide the name, like so:

```
func (c *Car) GetWheels() int {
	return c.wheels
}
```

Becuase we have called it this way the function has access to the private fields within the object (as this fucntion was defined in the same package as the struct). 


LAB TASK
--------

So far we have been able to create cars and bikes and see what their numbers of wheels are using the GetWheels() function. Extend the code to include a SetWheels function that takes in an integer and sets the amount of wheels that car or bike has to the provided amount. Then extend the main code to print out the new value. 




