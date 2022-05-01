STRUCTS
=======

Structs are a way where we can create a group of fields together to define an object. Structs are go's implementation of Object Orientated programming. 


CREATING A NEW STRUCT
---------------------

We create all structs the same way:


we start with the `type` keyword meaning we are creating a new variable type

this is immediately followed by the name of the type we are creating

then we end the statement with struct to show it is a collection of other types
 
between the curly braces we can then define the fields that make up this struct

```
type MyStruct struct {
  ...
}
```

For example if we wanted to create a shopping service then we could create a user with the following details. 

```
type User struct {

  Email string
  
  FirstName string
  LastName string
  
  AddressFirstLine string
  AddressSecondLine string
  City string
  Postcode string
}
```

we can then create a variable of that type in the same way we work with builtin variables and

```
// using the var keyword to create the variable
var myUser User

// we can then use dot notation to set fields within the variable or print them out
myUser.Email = "foo@bar.com"
println(myUser.Email) // foo@bar.com

// we can also use the := notation for creation and assignment like other variables
myUserTwo := User{
  Email:     "john.doe@bar.com",
  FirstName: "John",
  LastName:  "Doe",
}

println(myUserTwo.Email) // john.doe@bar.com

```

LAB TASK
--------

In this directory you can find a file called `lab.go` which has some broken code inside of it. Fix the definition of the Person struct so that when you run `go run lab.go` you get the following output. 

```
> go run lab.go
Hello I am John Doe and I am 30 years old
```




