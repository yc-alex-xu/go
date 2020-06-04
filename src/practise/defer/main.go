/*
defer 对普通函数与闭包的影响

Go functions may be closures. A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.


*/
package main

import "fmt"

func main() {
	f()
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
		//a lambda function as closure
		defer func(n int) {
			fmt.Println("lamda=>", n*2)
		}(i) //
	}
}

func f() {
	for i := 0; i < 4; i++ {
		//function与context是联系在一起，所以每一次循环都是一个新函数
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
