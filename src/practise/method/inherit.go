package main

import "fmt"

type human struct {
	name  string
	age   int
	phone string
}

type student struct {
	human  //匿名字段
	school string
}

type employee struct {
	human   //匿名字段
	company string
}

//在human上面定义了一个method
func (h human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//override
func (e employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

func main() {
	mark := student{human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := employee{human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}
