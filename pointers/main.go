package main

import "fmt"

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

type Fish struct {
	Name string
}

func (c Cat) SetName(name string) {
	c.Name = name
}

func (c Cat) GetName() string {
	return c.Name
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d *Dog) GetName() string {
	return d.Name
}

func GetName(f Fish) string {
	return f.Name
}

func SetName(f Fish, name string) Fish {
	f.Name = name
	return f
}

func SetNamePntr(f *Fish, name string) {
	f.Name = name
}

func main() {
	dog := Dog{}
	cat := Cat{}
	fish := Fish{}
	fish2 := &Fish{}

	dog.SetName("Rudolph")
	cat.SetName("Tibby")
	fish = SetName(fish, "Bubbles")
	SetNamePntr(fish2, "Nemo")

	fmt.Println("the dogs name is", dog.GetName())
	fmt.Println("the cats name is", cat.GetName())
	fmt.Println("the fishes name is", GetName(fish))
	fmt.Println("the seconds fishes name is", GetName(*fish2))
}
