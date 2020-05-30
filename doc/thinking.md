# go 相比c/c++,自带工程属性，如
* go fmt 解决code style 正义
* 目录组织和package 来解决源码组织问题，
* godoc 解决 help 问题

# 语法检查严格
* 变量定义了未使用，包引入了未使用　　　　这些在c++中warning 都不算，但在go 就是error了。
* Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错,所以只能使用**_**来丢弃不需要的返回值

# 谨慎
* new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值. 而C/C++是不浪费时间做零值填充的，除非你定义了构造函数。
* 

# 借鉴其他语言
* array、slice、map
* make、new操作
* defer
* Panic和Recover
* init函数
* smart pointer ????

# 函数
* 保留c语言的不定数量的参数
* 可以返回多个值，不用定义tuple

# OO
* struct
C++ 中继承是用父类之类的概念，而go 中直接的多，是用struct B 包含匿名的 struct B,C的方法，表示struct B 包含了struct B,C的所有field和方法。
```go
type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
	loan float32
}

type Employee struct {
	Human //匿名字段
	company string
	money float32
}

```
# TDD
go test, 会自动读取源码目录下面名为*_test.go的文件，生成并运行测试用的可执行文件。默认的情况下，不需要任何的参数，它会自动把你源码包下面所有test文件测试完毕，当然你也可以带上参数，详情请参考go help testflag


* method
* interface: 可以抽象的方式来代表一类对象
