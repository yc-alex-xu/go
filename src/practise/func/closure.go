package main

import (
	"fmt"
	"log"
	"runtime"
)

/*代替了debug的宏
 */
var where = func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}

func main() {
	//不需要时关闭调试语句
	where = func() {}
	f := adder()
	for i := 1; i < 10; i++ {
		fmt.Println("+", i, "=", f(i))
	}
	where()
	adder2()
}

func adder() func(int) int {
	var x int //起了类似静态变量的作用
	return func(delta int) int {
		//在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明的：
		x += delta
		return x
	}
}

func adder2() {
	var g int
	func(i int) {
		s := 0
		for j := 1; j < i; j++ {
			fmt.Println("+", j)
			s += j
		}
		where()
		g = s
	}(10)
	fmt.Println("=", g)
}
