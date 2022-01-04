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
func newBoke() bike  // this cannot be used outside the package as it starts with a lower case character
```


LAB TASK
--------

In the code for this lab we have created the functions for add and divide but we haven't filled out any of the logic for this. Update the `add` and `divide` functions to ensure they work as we expect. Hint: for the divide function you will need the `%` operator. 




