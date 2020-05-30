# go 相比c/c++,自带工程属性，如
* go fmt 解决code style 正义
* 目录组织和package 来解决源码组织问题，
* godoc 解决 help 问题


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
* method
* interface: 可以抽象的方式来代表一类对象
