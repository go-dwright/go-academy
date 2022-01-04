LOOPS
=====

Loops are always an important part of coding and here are some examples of how to set them up and use them

WAYS OF MAKING LOOPS
--------------------

There are two mains ways of running loops in go, the first is very typical in coding and should look very familiar if you have ever coded in c or java. 

```	
for i := 0; i < 10; i++ {
		// do something
}
``` 

In this example we create an integer called i, give it a value of 0 and then while i is under 10 run the loop and at the end of each loop increment i. Very typical for loop. 

We also have for range loops which use slices, in go you can think of slices the same way other languages use arrays, the only difference is that they do not have a fixed size so they can grow or contract as required. 

```
names := []string{"Daniel","Rupert","Emma"}
for i,name := range names {
  print(i)
  print(name)
}
```

when we use the range keyword for loops, we get back 2 values; the element and the value. So in the example above in the first loop we would get i=0 and name=Daniel, the second would be i=1 and name=Rupert and finally i=2 and name=Emma. If we don't care about one of the two values, e.g. we only care about the names and we don't want to have the element number we can replace it with an underscore to throw away the value. 

```
for _,name := range names {
  print(name)
}
```


LAB TASK
--------

```
names := []string{"Daniel","Rupert","Emma"}
colours := []string{"Red","Blue","Green","Yellow"}

```

Create a program using loops that creates following output using the two slices above.  

```
Daniel
Red
Blue
Green
Yellow
Rupert
Red
Blue
Green
Yellow
Emma
Red
Blue
Green
Yellow
```

