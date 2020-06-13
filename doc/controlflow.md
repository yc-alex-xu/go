# if

# switch

# goto

# for

# func
[code example](https://github.com/yc-alex-xu/go/tree/master/src/practise/func)
* byPointer.go  参数传值
* byValue.go 　　参数传指针 
* funcPointer.go
* closure.go    闭包函数

## 不定数量的参数
```go
func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}
```
alex: 实际就传一slice，只是syntax上与c语言类似一下

## 函数作为值、类型

## closure
闭包（Closure）是词法闭包（Lexical Closure）的简称。对闭包的具体定义有很多种说法,本文也无意学术气息的讨论,只是介绍一下自己的理解.
* Closure 不是所有的语言支持的,如Python,javascript是支持
* Closure 的目的是为了保存状态/历史, Java,C++ 中的对象也有类似的作用
* Closure 有代价的,往往意味着所在的fuction object 得延迟释放内存.

# defer
有点类似
```c
#include <stdlib.h>

int on_exit(void (*function)(int , void *), void *arg);
```
不过defer 是func级别的，并且组成一个stack. [example](https://github.com/yc-alex-xu/go/tree/master/src/practise/defer)

# Panic和Recover
Go没有像Java那样的异常机制，它不能抛出异常，而是使用了panic和recover机制。一定要记住，你应当把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有panic的东西。这是个强大的工具，请明智地使用它。那么，我们应该如何使用它呢？

# main函数和init函数
Go里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）。这两个函数在定义时不能有任何的参数和返回值。虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中每个文件只写一个init函数。

# builtin function
```bash
$ go doc builtin
package builtin // import "builtin"
func close(c chan<- Type)
func delete(m map[Type]Type1, key Type)
func panic(v interface{})
func print(args ...Type)
func println(args ...Type)
func recover() interface{}
func cap(v Type) int
func copy(dst, src []Type) int
func len(v Type) int
type ComplexType complex64
    func complex(r, i FloatType) ComplexType
type FloatType float32
    func imag(c ComplexType) FloatType
    func real(c ComplexType) FloatType
type IntegerType int
type Type int
    var nil Type
    func append(slice []Type, elems ...Type) []Type
    func make(t Type, size ...IntegerType) Type
    func new(Type) *Type
```