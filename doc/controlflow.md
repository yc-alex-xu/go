# if
也许是各种编程语言中最常见的了，它的语法概括起来就是：如果满足条件就做某事，否则做另一件事。Go里面if条件判断语句中不需要括号，如下代码所示
```go
if x > 10 {
	fmt.Println("x is greater than 10")
} else {
	fmt.Println("x is less than 10")
}

if integer := func();integer == 3 {//条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内
	fmt.Println("The integer is equal to 3")
} else if integer < 3 {
	fmt.Println("The integer is less than 3")
} else {
	fmt.Println("The integer is greater than 3")
}

```

# switch
```go
switch sExpr {
case expr1:
	some instructions
case expr2:
	some other instructions
case expr3:
	some other instructions
default:
	other code
}
```
Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。

# goto
请明智地使用它。用goto跳转到必须在当前函数内定义的标签。例如假设这样一个循环：
```go
func myFunc() {
	i := 0
Here:   //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	goto Here   //跳转到Here去
}
```
# for
它既可以用来循环读取数据，又可以当作while来控制逻辑，还能迭代操作。它的语法如下：
```go
for expression1; expression2; expression3 {
	//...
}

for k,v:=range map { //for配合range可以用于读取slice/map/channel的数据：
	fmt.Println("map's key:",k)
	fmt.Println("map's val:",v)
}
```
在循环里面有两个关键操作break和continue ,break操作是跳出当前循环，continue是跳过本次循环。当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置，详细参考如下例子：
```go
for index := 10; index>0; index-- {
	if index == 5{
		break // 或者continue
	}
	fmt.Println(index)
}
// break打印出来10、9、8、7、6
// continue打印出来10、9、8、7、6、4、3、2、1
```
note:

由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃不需要的返回值 例如
```go
for _, v := range map{
	fmt.Println("map's val:", v)
}
```

# func
[code example](https://github.com/yc-alex-xu/go/tree/master/src/practise/func)
* byPointer.go  参数传值
* byValue.go 　　参数传指针 
* funcPointer.go
* closure.go    闭包函数

```go
func funcName(input1 type1, input2 type2) (type1, type2) {
	//这里是处理逻辑代码
	//返回多个值
	return output1, output2
}
/*如果只返回一个值的话，括号可以省略*/
func abc(x,y,z int) error{
	
}
```

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
在Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型
```go
type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
```
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