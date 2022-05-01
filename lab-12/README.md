CONSUMING DATA
==============

A lot of the data we like our programs to consume comes in some well known types like json or yaml. Wouldn't it be great if go understood yaml or json and allowed us to converted directly to or from json to a struct? 

TAGS
----

When we create a struct we can add tags if we know we are going to be converting to/from a given data type. For example 

```
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
```

This Person struct has json tags definied for each of its fields, which means that if we get sent a json message with those fields, we can immediately convert that into a Person struct that we can then manipluate. To do that we can do the following 


```
jsonString := `{"first_name":"John","last_name":"Doe","age":32}`
var person Person
json.Unmarshal([]byte(jsonString), &person)
```

In this example we have a json string that contains all our fields, and we create ourselves a new Person using the var keyword that is currently only using the default values. We then call json.Unmarshal where we take our jsonString (and cast it to a slice of bytes) and then the result of that is put into the person variable we just created. We can then use that object like normal and all the fields will be set. 

```
fmt.Println(person.FirstName)
// John
```