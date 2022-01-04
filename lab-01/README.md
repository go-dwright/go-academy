HELLO WORLD
===========

The first program anyone writes in any language is hello world. This will be no different. You can see the entirety of the program in the `main.go` file in this directory. There are a few things to note here, any golang language files should always end in `.go` and by convention we have named this file with the prefix `main` so we expect to find our `main()` function here. 

main() is a convention that goes all the way back to the first versions of `C`. When we run a golang program the main function is the start point (in most cases) of where we run code. 

BUILDING & RUNNING
------------------

Ensure your terminal is in this directory then you can either build and execute by using the following commands:
```
go build -o app . 
./app
``` 

There are several parts to the first command and lets go through them all:

`go build` : This is a command and subcommand, `go` being our command and `build` being our subcommand. We use this subcommand if we want to compile (build) go code. 

`-o app` : This is the output flag, so we can explictly name the file that is built by `go build`. In this case it will be `app`, if we didn't provide this flag it would take the name of the directory.

`.` : This is the file OR directory that we want go to look in for our main function and then build from there. Note that we can only compile code if golang finds a `main()` function within a file that is a part of the main package. 


You can also run the shorthand command which compiles and runs the code without creating a compiled file:
```
go run .
```

LAB TASK
--------

Create a program that outputs the following when ran

```
Hello World!
Welcome To Go!
```



