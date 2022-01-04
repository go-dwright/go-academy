CONDITIONALS
============

Conditionals are the ways that our code can make decisions. So far all our code has never has had to make a descision, it has just done one step and then the next. 

IF CASE
-------

The most famous of conditionals is the `if` statement which always boils down to a boolean statement, e.g. is something true or false. If something is true, then the code within the `if` will be ran, if it is false then it will not be. 

In general when we use the `if` keyword we will want to be comparing something to something else. In a lot of cases we want to see if something is that same as something else, but not always. As such, there are a set of operators in go for comparing, you can see the most common of these in the table below.

| Operator | Name                     | Example | Result                                  |
|----------|--------------------------|---------|-----------------------------------------|
| ==       | Equal                    | x == y  | True if x is equal to y                 |
| !=       | Not equal                | x != y  | True if x is not equal to y             |
| <        | Less than                | x < y   | True if x is less than y                |
| <=       | Less than or equal to    | x <= y  | True if x is less than or equal to y    |
| >        | Greater than             | x > y   | True if x is greater than y             |
| >=       | Greater than or equal to | x >= y  | True if x is greater than or equal to y |


There is another conditional known as the switch/case. We will be looking at that in another lab. 


LAB TASK
--------

Write a program that has a variable called `myAge` and use that to determine what year the user was born in. If they were born before 2000 it should print out "You were born in the 20th Century", if they were born after 2000 it should print "You were born in the 21st Century". If they were born on the year 2000 it should print "You were born on the millenium!"

