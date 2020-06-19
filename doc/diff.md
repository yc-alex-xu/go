
##  package 
可以说是source code file的逻辑名，同一目录下多个source code file 只要用同样的package name，那就跟合成一个文件没区别。不用**include**啦。Packages in Go serve the same purposes as libraries or modules in other languages, supporting modularity, encapsulation, separate compilation, and reuse.

## tuple assignment
**c++ tuple**类型就没有必要了。

## pointer
C++ References are often confused with pointers, but three major differences between references and pointers are 
   - You cannot have NULL references. You must always be able to assume that a reference is connected to a legitimate piece of storage.
   - Once a reference is initialized to an object, it cannot be changed to refer to another object.  Pointers can be pointed to another object at any time.
   - A reference must be initialized when it is created. Pointers can be initialized at any time.
  
we can read or update the value of a variable **indirectly** via a pointer, without using or even knowing the name of the variable, if indeed it has a name. pointer of 像是综合了**pointer of c 的写法和reference in C++的用法**. 

## local variables 
have dynamic lifetimes。 A compiler may choose to allocate local variables on the heap or on the stack but, perhaps surprisingly, this choice is not detemined by whether var or new was used to declare the variable. 也就是说，程序员不用操心，一个func可以返回其locall variable,放心用，gc不会愚蠢的recycle其占用空间。Garbage collection is a tremendous help in writing correct programs, but it does not relieve you of the burden of thinking about memory. You don’t need to explicitly allocate and free memory, but to write efficient programs you still need to be aware of the lifet ime of variables.For example, keeping unnecessary pointers to short-lived objects within long-lived objects,especially global variables, will prevent the garbage collector from reclaiming the short-lived objects.

![a good exmpale of scope](images/scope.png)

## types 
 go 一方面是strict type, “type Celsius float64” 就定义了一种新type,不允许自动转换；另外一方面 In any case, a conversion never fails at run time.

 Go’s types fall into four categories: **basic types, aggregate types, reference types, and interface types**. 前两种各种语言都差不多。Reference types are a diverse group that includes pointers,slices, maps, **functions**, and channels , but what they have in common is that they all refer to program variables or state **indirectly**, so that the effect of an operation applied to one reference is observed by all copies of that reference.

array:

array as fuction parameter: When a function is called, a copy of each argument value is assigned to the corresponding parameter variable, so the function receives a copy, not the original. C/C++这时传的是地址

go 的type declaration 没有次序概念，所以也不需要forward declaration,e.g.
```go
func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

```

function type:

```go
package http // import "net/http"

type HandlerFunc func(ResponseWriter, *Request)
    The HandlerFunc type is an adapter to allow the use of ordinary functions as
    HTTP handlers. If f is a function with the appropriate signature,
    HandlerFunc(f) is a Handler that calls f.

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

```


## comparable

basic type: 

OK

aggreagte type: 
* **array** 只要类型（element type, len）相同即可， 
* If all the fields of a **struct** are comparable, the struct itself is comparable, so two expressions of that type may be compared using == or !=.

reference type: 
* As with **slices, maps** cannot be compared to each other ; the only legal comparison is with nil.
* **pointer** only can be compared with nil. 
* The **function** values are not just code but can have state.The anonymous inner function can access and update the local variables of the enclosing function squares. These hidden variable references are why we classify functions as reference types and why function values are not comparable. 

interface type:  

it is comparable; otherwise if two interface values are compared and have the same dynamic type, and  that type is comparable, the **interface** is comparable,  if not comparable (a slice, for instance), then the comparison fails with a panic. 


* [example of string comparation](../src/practise/comparable/main.go)

## syntactic sugar about type

struct: The dot notation also works with a pointer to a struct. 也就是说c/c++中的->符号不需要了。 
```go
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
```

 method: implicit conversion: if the receiver argument is a variable of type T and the receiver parameter has type *T. The compiler implicitly takes the address of the variable: p.ScaleBy(2) // implicit (&p). Or the receiver argument has type *T and the receiver parameter has type T. The compiler implicitly dereferences the receiver, in other words, loads the value: pptr.Distance(q) // implicit (*pptr) . 所以 So to avoid ambiguities, method declarations are not permitted on named types that are themselves pointer tyes；并且对同一type T,同时定义T和*T的方法，会报error: method redeclared. 但这语法糖的基础是，调用语句用的是一个variable,这样才可以取到地址，直接用type来调用无法取得地址不行。

## OO
*  虽然go 没有c++中class概念，但不限于struct,任何named type都可以看成一个class,如在其上定义method
* The zero-value mechanism: ensures that a variable always holds a well-defined value of its type。加上the init function mechanism 就不用**c++ constructor** 啦.
* Encapsulation: package/Struct/interface level 的varialbe,如果不是大写字母开头，类似**static variable in C**, method 也类似。
* Composition is central to object-oriented programming in Go,
* In Go, we don’t use a special name like this or self for the (method) receiver; we choose receiver names just as we would for any other parameter. 
* Nil Is a Valid (method) Receiver Value
* Composing Types by Struct Embedding： Distance has a parameter of type Point, and q is not a Point, so although q does have an embedded field of that type, we must explicitly select it. e.g. p.Distance(q.Point)
* Method Values 可以作为函数指针用
* As shorthand, Go programmers often say that a concrete type ‘‘is a’’ particular interface type, meaning that it satisfies the interface. interface is abstract type, 同时也有动态语言的特点：interface’s dynamic type and dynamic value；对于interface定义的method, 也是dynamic dispatch。Internally, Go can uses reflection to obtain the name of the interface’s dynamic type. 
* These interfaces are but one (idion: only one) useful way to group related concrete types together and express the facets they share in common.不能把它看成c++的abstract class. in Go we can define new abstractions or groupings of interest when we need them, without modifying the declaration of the concrete type. 这里有点方法论的味道，也就是说先有concrete type，由于refactor 需要，再创建用于grouping共性的interface type. Interfaces are only needed when there are two or more concrete types that must be dealt with in a uniform way.


## exception
 Go programs use ordinary control-flow mechanisms like **if** and **return** to respond to errors, no exception mechanism.

## programming paradigm:Functional Programming
闭包和高阶函数

函数编程支持函数作为第一类对象，有时称为闭包或者仿函数（functor）对象。实质上，闭包是起函数的作用并可以像对象一样操作的对象。与此类似，FP 语言支持高阶函数。高阶函数可以用另一个函数（间接地，用一个表达式） 作为其输入参数，在某些情况下，它甚至返回一个函数作为其输出参数。这两种结构结合在一起使得可以用优雅的方式进行模块化编程，这是使用 FP 的最大好处。 
惰性计算

除了高阶函数和仿函数（或闭包）的概念，FP 还引入了惰性计算的概念。在惰性计算中，表达式不是在绑定到变量时立即计算，而是在求值程序需要产生表达式的值时进行计算。延迟的计算使您可以编写可能潜在地生成无穷输出的函数。因为不会计算多于程序的其余部分所需要的值，所以不需要担心由无穷计算所导致的 out-of-memory 错误。一个惰性计算的例子是生成无穷 Fibonacci 列表的函数，但是对第n个Fibonacci 数的计算相当于只是从可能的无穷列表中提取一项。

Functions are first-class values in Go: like other values, function values have types, and they may be assigned to variables or passed to or returned from functions. 这说法跟Python非常类似。  The function values are not just code but can have state.The anonymous inner function can access and update the local variables of the enclosing function squares. These hidden variable references are why we classify functions as reference
types and why function values are not comparable. Function values like these are implemented using a technique called closures , and Go programmers often use this term for function values. Here again we see an example where the lifetime of a variable is not determined by its scope.

* [when closure capture loop variable](../src/practise/func/closureLoopVariable/main.go)
  
 该程序还涉及goroutine 的类似linux/c的waitpid的用法

## Goroutines
In Go, each concurrently executing activity is called a goroutine. When a program starts, its only goroutine is the one that calls the main function, so we call it the main goroutine. New goroutines are created by the go statement. There is no programmatic way for one goroutine to stop another, but as we will see later, there are ways to communicate with a goroutine to request that it stop itself.

If goroutines are the activities of a concurrent Go program, channels are the connections between them.

```go
ch = make(chan int) // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```
A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel.When a value is sent on an unbuffered channel, the receipt of the value happens before the reawakening of the sending goroutine.


# note:
## difference of printf 
```go
	var r rune = '国'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
```   
## more kinds of const declaration

const: 
* [some example](../src/practise/const/main.go)

  
# todo:
vs code debug时，如何输入stdin的内容？？？

https://stackoverflow.com/questions/50884981/how-to-read-input-when-debugging-go-in-visual-studio-code




