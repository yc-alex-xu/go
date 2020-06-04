# first program
* [lib](https://github.com/yc-alex-xu/go/tree/master/src/practise/math)
* [main program](https://github.com/yc-alex-xu/go/blob/master/src/practise/app1/main.go)

note:
* 包名和包所在的文件夹名可以是不同的，此处的<pkgName>即为通过package <pkgName>声明的包名，而非文件夹名。

# keywords
```
break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var
```

# package
Go程序是通过package来组织的

package <pkgName>（在我们的例子中是package main）这一行告诉我们当前文件属于哪个包，而包名main则告诉我们它是一个可独立运行的包，它在编译后会产生可执行文件。除了main包之外，其它的包最后都会生成*.a文件（也就是包文件）并放置在$GOPATH/pkg/$GOOS_$GOARCH中.

可见性规则：

* 大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
* 大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

# encoding
Go是天生支持UTF-8的，任何字符都可以直接输出，你甚至可以用UTF-8中的任何字符作为标识符。
# var
```go
//定义一个名称为“variableName”，类型为"type"的变量
var variableName type
//定义三个类型都是“type”的变量
var vname1, vname2, vname3 type
//初始化“variableName”的变量为“value”值，类型是“type”
var variableName type = value
/*
	定义三个变量，它们分别初始化为相应的值
	vname1为v1，vname2为v2，vname3为v3
	然后Go会根据其相应值的类型来帮你初始化它们
*/
var vname1, vname2, vname3 = v1, v2, v3
```
* Go 语言也是区分大小写的，这与 C 家族中的其它语言相同
* **_**（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。

# const
```go
const constantName = value
```
# builtin数据类型
go 是强类型，并且没有缺省转换的。
## Boolean
在Go中，布尔值的类型为bool，值是true或false，默认为false。

## int 
有无符号和带符号两种。Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。Go里面也有直接定义好位数的类型：rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。其中rune是int32的别称，byte是uint8的别称。

## 浮点数
的类型有float32和float64两种（没有float类型），默认是float64

## 复数
complex128（64位实数+64位虚数）。如果需要小一些的，也有complex64(32位实数+32位虚数)。复数的形式为RE + IMi

需要注意的一点是，这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错

## string
我们在上一节中讲过，Go中的字符串都是采用UTF-8字符集编码。字符串是用一对双引号（""）或反引号（` `）括起来定义，它的类型是string。

string 是Immutable的，如要修改需要转换成slice.

```go
s := "hello"
c := []byte(s)  // 将字符串 s 转换为 []byte 类型
c[0] = 'c'
s2 := string(c)  // 再转换回 string 类型
fmt.Printf("%s\n", s2)
```
感觉跟Python很像,不过Go的类型转化写法还是挺规整的的　type(var) 即可。

## error类型
Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误：

## iota枚举
Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1：

# array、slice、map
## array
array就是数组，它的定义方式如下：
```go
var arr [n]type
```
由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。如果要使用指针，那么就需要用到后面介绍的slice类型了。

## slice
在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice

slice并不是真正意义上的动态数组，而是一个**引用类型**(alex:此引用跟c++的reference概念应该是完全不同。内部实现是一个管理用struct吧，毕竟它自己的address得固定)。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度(alex:估计类似c++ string的实现，可以自动扩充数组大小)。
```go
// 和声明array一样，只是少了长度
var fslice []int
```
从概念上面来说slice像一个结构体，这个结构体包含了三个元素：
* 一个指针，指向数组中slice指定的开始位置
* 长度，即slice的长度
* 最大长度，也就是slice开始位置到数组的最后位置的长度

slice有一些简便的操作

* slice的默认开始位置是0，ar[:n]等价于ar[0:n]
* slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
* 如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]

slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，
[main program](https://github.com/yc-alex-xu/go/blob/master/src/practise/app1/main.go)
的a和b，如果修改了a Slice中元素的值，那么b Slice相对应的值也会改变。

从概念上面来说slice像一个结构体，这个结构体包含了三个元素：

* 一个指针，指向数组中slice指定的开始位置
* 长度，即slice的长度
* 最大长度，也就是slice开始位置到数组的最后位置的长度

对于slice有几个有用的内置函数：

* len 获取slice的长度
* cap 获取slice的最大容量
* append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
* copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数

## map
map也就是Python中字典的概念，它的格式为map[keyType]valueType

我们看下面的代码，map的读取和设置也类似slice一样，通过key来操作，只是slice的index只能是｀int｀类型，而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。

使用map过程中需要注意的几点：

* map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
* map的长度是不固定的，也就是和slice一样，也是一种引用类型
* 内置的len函数同样适用于map，返回map拥有的key的数量
* map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
* map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制

# make、new操作
make用于内建类型（map、slice 和channel）的内存分配。new用于其他各种类型的内存分配。

## new
本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：

new返回指针。

## make(T, args)
与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。
Alex: new 分配该数据类型内存，然后内存清0; 而make 对应的slice、map和channel不仅仅是分配内存清0这么简单，它还有创建对应的管理用数据结构。

# 值类型和引用类型
* 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值.array 和strut 也是值类型。你可以通过 &i 来获取变量 i 的内存地址。值类型的变量的值存储在栈中。
* 一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置. 在 Go 语言中，指针属于引用类型，其它的引用类型还包括 slices（，maps和 channel。被引用的变量会存储在堆中，以便进行垃圾回收，且比栈拥有更大的内存空间




# struct
类似C++概念。


# 自定义数据类型
```go
type typeName typeLiteral
```
struct只是自定义类型里面一种比较特殊的类型而已，还有其他自定义类型申明，e.g.
```go
type ages int

type money float32

type months map[string]int

m := months {
	"January":31,
	"February":28,
	...
	"December":31,
}
```

# pass by value or address
* value:基本数据类型，int,float,bool,string, 以及数组和 struct 特点：变量直接存储值，内存通常在栈中分配，栈在函数调用完会被释放
* address: 变量存储的是一个地址，这个地址存储最终的值。内存通常在 堆上分配。通过 GC 回收。引用类型： 指针、slice 切片、管道 channel、接口 interface、map、函数等

# pointer
[rbt & its testing](https://github.com/yc-alex-xu/go/tree/master/src/practise/rbtree)