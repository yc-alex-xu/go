package main

import (
	"fmt"
	"log"
	"runtime"
)

/*代替了debug的宏
 */
var logF = func() {
	_, file, line, _ := runtime.Caller(0)
	log.Printf("%s:%d", file, line)
}

func main() {
	//不需要时关闭调试语句，但不能赋 nil
	logF = func() {}
	acc := adder()
	for i := 1; i < 10; i++ {
		fmt.Println("+", i, "=>", acc(i))
	}
	acc2()
}

func adder() func(int) int {
	var x int //起了类似静态变量的作用
	return func(delta int) int {
		logF()
		x += delta //access 在外部声明的变量
		return x
	}
}

func acc2() {
	var g int
	func(i int) {
		s := 0
		for j := 1; j < i; j++ {
			fmt.Println("+", j)
			s += j
		}
		logF()
		g = s
	}(10)
	fmt.Println("=", g)
}
