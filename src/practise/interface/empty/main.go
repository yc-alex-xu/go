package main

import (
	"fmt"
	"strconv"
)

//List  type: slice of empty interfae
type List []interface{} //slice

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	var e interface{}
	var i int = 5
	s := "Hello world"
	e = i                     //空interface 可以代表任何数据类型
	fmt.Printf("e: %v \n", e) //可以自动判别element
	e = s
	fmt.Printf("e: %v \n", e) //可以自动判别element

	fmt.Print("e: ")
	/*
		value, ok = e.(T) 语法
	*/
	if value, ok := e.(int); ok {
		fmt.Printf("is an int and its value is %d\n", value)
	} else if value, ok := e.(string); ok {
		fmt.Printf("is a string and its value is %s\n", value)
	} else {
		fmt.Printf("is of a different type\n")

	}
	interfaceType()
}

/*.(type) 用法
 */
func interfaceType() {
	l := make(List, 3) //make a slice
	l[0] = 1           //an int
	l[1] = "Hello"     //a string
	l[2] = person{"Dennis", 70}

	for index, element := range l {
		switch value := element.(type) {
		case int:
			fmt.Printf("l[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("l[%d] is a string and its value is %s\n", index, value)
		case person:
			fmt.Printf("l[%d] is a person and its value is %s\n", index, value)
		default:
			fmt.Println("l[%d] is of a different type", index)
		}
	}
}
