package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("at the begin  x=:", x)
	v := reflect.ValueOf(x)        // 首先需要把它转化成reflect对象, reflect.TypeOf 或者 reflect.ValueOf
	fmt.Println("type:", v.Type()) //获取反射值能返回相应的类型和数值
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	p := reflect.ValueOf(&x) //传引用
	v = p.Elem()             //reflect对象转化成相应的值
	v.SetFloat(7.1)
	fmt.Println("at the end x=:", x)
}
