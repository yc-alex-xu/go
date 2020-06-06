package main

import "fmt"

type human struct {
	name  string
	age   int
	phone string // human类型拥有的字段
}

type employee struct {
	human      // 匿名字段human
	speciality string
	phone      string // 雇员的phone字段
}

func main() {
	Bob := employee{human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 访问human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.human.phone)
}
