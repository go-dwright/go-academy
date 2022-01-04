VARIABLES
=========

This lab we will be looking into variables and how to use them as part of a static compiled language.

CREATING VARIABLES
------------------

There are two ways of creating variables within go. We have the basic form using the `var` keyword

```
var age int
var name string
```

As go is statically typed, when we create a variable we provide it a type and that type is permanent. In the example above we create age as an integer and because we used the var keyword, it now has the default value for an int which is 0. There is no way we can change the type of age after it has been created, it will always be type int. 

We can also create and assign values to variables at the same time. Go is clever enough to work out what the kind of data you are providing and then automatically setting the correct type for your variable.

```
age := 27
name := "Damon"
```

We have typed a variable name that doesn't exist and we have used the `:=` assigner, so golang will take the value on the right-hand side and determine the type of that value and create a variable with that type on the left and assign the value. In this case age is still an int and name has been set as a string. 


NAMING VARIABLES
----------------

The way we name variables and functions is very important in go. We will go into more detail why later but for now variables should be named using lowerCamelCase. This means that if we want to create a variable with multiple words in the name, the first word should be lower case and then each subsequent word should start with a capital.

```
var firstName string
var lastName string

var ageInYears int
var ageInMonths int

var aVeryLongVariableNameWithMultipleWordsInIt bool
```


LAB TASK
--------

Below you can see an example program. 

```
package main

func main() {
	var myString string
	println(myString)
}
```

Extend this program so that it outputs the following output when ran, using multiple `println(myString)` statements, do not create any other variables. 

```
Hello World!
1234
Welcome To Go!
5678
```

