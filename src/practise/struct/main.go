package main

import (
	"fmt"
)

// 声明一个新的类型
type person struct {
	name string
	age  int
}

// older 比较两个人的年龄，返回年龄大的那个人
func older(p1, p2 person) person {
	if p1.age > p2.age { // 比较p1和p2这两个人的年龄
		return p1
	}
	return p2

}

func main() {
	var tom person

	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化
	bob := person{age: 25, name: "Bob"}

	// 按照struct定义顺序初始化值
	paul := person{"Paul", 43}

	senior := older(older(tom, bob), paul)
	fmt.Printf("%s is the oldest as age %d \n",
		senior.name, senior.age)

}
