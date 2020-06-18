就像学外语一样，死记硬背语法是一个方法，模仿成正确的习惯也是一种方法。

这里就记录一些有趣的例子

1. package 可以说是source code file的逻辑名，同一目录下多个source code file 只要用同样的package name，那就跟合成一个文件没区别。不用**include**啦。Packages in Go serve the same purposes as libraries or modules in other languages, supporting modularity, encapsulation, separate compilation, and reuse.
2. The zero-value mechanism ensures that a var iable always holds a wel l-define d value of its typ。加上the init function mechanism 就不用**c++ constructor** 啦.
3. 支持tuple assignment，**c++ tuple**类型就没有必要了。
4. With a pointer, we can read or update the value of a variable **indirectly**, without using or even knowing the name of the variable, if indeed it has a name. pointer of 像是综合了**pointer of c 的写法和reference in C++的用法**. 
5. local variables have dynamic lifetimes。 A compiler may choose to allocate local variables on the heap or on the stack but, perhaps surprisingly, this choice is not detemined by whether var or new was used to declare the variable. 也就是说，程序员不用操心，一个func可以返回其locall variable,放心用，gc不会愚蠢的recycle其占用空间。Garbage collection is a tremendous help in writing correct programs, but it does not relieve you of the burden of thinking about memory. You don’t need to explicitly allocate and free memory, but to write efficient programs you still need to be aware of the lifet ime of variables.For example, keeping unnecessary pointers to short-lived objects within long-lived objects,especially global variables, will prevent the garbage collector from reclaiming the short-lived objects.
7. go 一方面是strict type, “type Celsius float64” 就定义了一种新type,不允许自动转换；另外一方面 In any case, a conversion never fails at run time.
8. 虽然go 没有c++中class概念，但任何named type都可以看成一个class
> func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
9. package level 的varialbe,如果不是大写字母开头，类似**static variable in C**.
10. Go’s types fall into four categories: basic types, aggregate types, reference types, and interface types. 前两种各种语言都差不多。Reference types are a diverse group that includes pointers,slices, maps, **functions**, and channels , but what they have in common is that they all refer to program variables or state **indirectly**, so that the effect of an operation applied to one reference is observed by all copies of that reference.
11. OO: 
* Composition is central to object-oriented programming in Go,
* In Go, we don’t use a special name like this or self for the (method) receiver; we choose receiver names just as we would for any other parameter. 
* (syntactic sugar) implicit conversion: if the receiver argument is a variable of type T and the receiver parameter has type *T. The compiler implicitly takes the address of the variable: p.ScaleBy(2) // implicit (&p). Or the receiver argument has type *T and the receiver parameter has type T. The compiler implicitly dereferences the receiver, in other words, loads the value: pptr.Distance(q) // implicit (*pptr) . 所以 So to avoid ambiguities, method declarations are not permitted on named types that are themselves pointer tyes；并且对同一type T,同时定义T和*T的方法，会报error: method redeclared.
* Nil Is a Valid (method) Receiver Value
* Composing Types by Struct Embedding： Distance has a parameter of type Point, and q is not a Point, so although q does have an embedded field of that type, we must explicitly select it. e.g. p.Distance(q.Point)
* Method Values 可以作为函数指针用
* interface is abstract type,
12.  Go programs use ordinary control-flow mechanisms like if and return to respond to errors. 
13.  Functions are first-class values in Go: like other values, function values have types, and they may be assigned to variables or passed to or returned from functions. 这说法跟Python非常类似。  The function values are not just code but can have state.The anonymous inner function can access and update the local variables of the enclosing function squares. These hidden variable references are why we classify functions as reference
types and why function values are not comparable. Function values like these are implemented using a technique called closures , and Go programmers often use this term for function values. Here again we see an example where the lifetime of a variable is not determined by its scope:
the variable x exists after squares has returned within main, even though x is hidden inside f. 但是在loop 中,capture的一直是同一值，所以square2输出的一直是25. 这跟书中不符，跟Python的情况也不同。可以单步debug发现其奥秘。总之这一种非常有歧义的用法，GOPL的作者都无法掌握，大家就不要用了。
```go
// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func square2() []func() int {
	res := make([]func() int, 5)
	var x int
	for i := 0; i < 5; i++ {
		res[i] = func() int {
			x = i
			return x * x
		}
	}
	return res
}

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

	for _, f := range square2() {
		fmt.Println(f())
	}
}
```
14.syntactic sugar:The dot notation also works with a pointer to a struct. 也就是说c/c++中的->符号不需要了。 
```go
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
```
15. 







note:

1. C++ References are often confused with pointers, but three major differences between references and pointers are 
   - You cannot have NULL references. You must always be able to assume that a reference is connected to a legitimate piece of storage.
   - Once a reference is initialized to an object, it cannot be changed to refer to another object.  Pointers can be pointed to another object at any time.
   - A reference must be initialized when it is created. Pointers can be initialized at any time.2.
2. a good exmpale of scope

![scop](images/sope.png)

3. printf也是有差异的
```go
	var r rune = '国'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
```   
4. const: the previous expression and its type can be used again in a group
```go
	const (
		a = 1
		b
		c = 2
		d
	)
   fmt.Println(a, b, c, d) // "1 1 2 2"
	type Flags byte
	const (
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
   )
   //"1 10 100 1000 10000"
	fmt.Printf("%b %b %b %b %b\n", FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)   
```
5. array 只要类型（element type, len）相同，是可以用一句话来判断是否相等的，slice/map不行。As with slices, maps cannot be compared to each other ; the only legal comparison is with nil. If all the fields of a struct are comparable, the struct itself is comparable, so two expressions of that type may be compared using == or !=. 
```go
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b != c) // "true false true"
```
6. array as fuction parameter: When a function is called, a copy of each argument value is assigned to the corresponding parameter variable, so the function receives a copy, not the original. C/C++这时传的是地址
7. 
8. TBD
9. TBD
10. TBD
11. 
     


todo:
vs code debug时，如何输入stdin的内容？？？



