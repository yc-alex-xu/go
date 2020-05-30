package main

import "fmt"

func main() {

	type Element interface{}
	var element Element
	var i int = 5
	s := "Hello world"
	element = i //空interface 可以代表任何数据类型
	element = s
	fmt.Println("s:", s)
	fmt.Printf("element: %v ", element) //可以自动判别element
	s2 := element
	/*
		Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
	*/
	if value, ok := element.(int); ok {
		fmt.Printf("is an int and its value is %d\n", value)
	} else if value, ok := element.(string); ok {
		fmt.Printf("is a string and its value is %s\n", value)
	} else {
		fmt.Printf("is of a different type\n")

	}
	fmt.Println("s2 is:", s2)
}
