package main

func main() {
	num := 5
	goDoSomething(num)
}

func goDoSomething(num int) {
	num = num * num
	goDoSomethingElse(num)
}

func goDoSomethingElse(num int) {
	num = num + 10
	goPanic(num)
}

func goPanic(num int) {
	if num == 35 {
		panic("ah!!! I can't carry on if num is 35")
	}
}
