package main

import "fmt"

func main() {
	var e interface{}
	var i int = 5
	s := "Hello world"
	e = i                     //空interface 可以代表任何数据类型
	fmt.Printf("e: %v \n", e) //可以自动判别element
	e = s
	fmt.Printf("e: %v \n", e) //可以自动判别element
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
}
