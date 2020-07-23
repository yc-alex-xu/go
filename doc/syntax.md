# Names
25 个关键字或保留字

36 个预定义标识符  //可以用go doc builtin 检查 
```go
Constants: true false iota nil
Types: int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex128 complex64  bool byte rune string error
Functions: make len cap new append copy close delete complex real imag panic recover

p := new(int)
delete(p)
```

_ 被称为空白标识符，被用于抛弃值。

# Declarations

# const
The value of a constant must be basic type: a number, string , or boolean. e.g.
```go
true false iota nil
```
**Untyped Constants**: The compiler represents these uncommitted constants with much greater numeric precision than values of basic types, and arithmetic on them is more precise than machine arithmetic; you may assume at least 256 bits of precision. alex:也就是说，它们的精度太大，无法用basic type的空间来存储，所以compiler就给它们一个特例,告诉programmer你就不用操心它的data type了.

[code sample](../src/practise/const)

# Variables 
* have dynamic lifetimes。 A compiler may choose to allocate local variables on the heap or on the stack but, perhaps surprisingly, this choice is not detemined by whether var or new was used to declare the variable. 也就是说，程序员不用操心，一个func可以返回其locall variable,放心用，gc不会愚蠢的recycle其占用空间。
* Garbage collection is a tremendous help in writing correct programs, but it does not relieve you of the burden of thinking about memory. You don’t need to explicitly allocate and free memory, but to write efficient programs you still need to be aware of the lifet ime of variables.For example, keeping unnecessary pointers to short-lived objects within long-lived objects,especially global variables, will prevent the garbage collector from reclaiming the short-lived objects.

its scope:

![a good exmpale of scope](images/scope.png)

# Type 
```go
type name underlying-type
```

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

# control flow
## if
```go
    if f, err := os.Open(fname); err != nil {
        return err
    } else {
        // f and err are visible here too
        f.ReadByte()
        f.Close()
    }
```

## switch
A switch does not need an operand; it can just list the cases, each of which is a boolean expression. This form is called a **tagless switch**; it’s equivalent to switch true. 

Like the for and if st atements, a switch may include an optional simple statement—a short variable declaration, an increment or assignment statement, or a function call—that can be used to set a value before it is tested.

## for
The for loop is the only loop statement in Go.
```go
for initialization; condition; post {
// zero or more statements
}
```
其中initialization; condition; post 都是optional的。

Another form of the for loop iterates over a range of values from a data type like a string or a slice.

## break and continue
statements modify the flow of control .


# Basic data types & operators
* bool
* int
* float
* complex
* string

There is no char * in go. only string like C++
* Go 中的字符串是根据长度限定，而非特殊字符 \0
* string literal：双引号字面量不能跨行，如果想要多行字符串时，就必须使用`反引号
* 可以比较
* 可以用下标取值-->rune
* 有对应的strings 和 strconv  package

string & char 

go 中没有char 的概念，只有ascii的byte(uint8)和 UTF-8 的rune(int32),用‘’括起. e.g
```go
var ch1 byte = 65
var ch2 byte = '\x41'  
var ch3 int = '\u0041'  
fmt.Printf("%c - %c - %c\n", ch, ch2, ch3)
```
## bitwise binary operators
* 二元运算符:  &,|,^(XOR),&^ (bit clear AND NOT)
* 一元运算符:  ^ (取补数),<<(left shift), >>(right shift) 
* [code sample](../src/practise/bits) 

## Printf format
```go
    %d decimal integer
    %x, %o, %b integer in hexadecimal, octal, binary
    %f, %g, %e floating-point number: 3.141593 3.141592653589793 3.141593e+00
    %t boolean: true or false
    %c rune (Unicode code point)
    %s string
    %q quoted string "abc" or rune 'c'
    %v any value in a natural format
    %T type of any value
    %% literal percent sign (no operand)

	var r rune = '国'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
```   

# aggreagte type
## array:
array as fuction parameter: When a function is called, a copy of each argument value is assigned to the corresponding parameter variable, so the function **receives a copy**, not the original. C/C++这时传的是地址

不过可以直接传地址,但无论如何，arrays are still inherently inflexible because of their fixed size， 不适合做函数参数
```go
months := [...]string{1: "January", /* ... */, 12: "December"}//中括号里为空的话，compiler就识别为slice
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}//这种赋值方式类似struct
}
```
## struct
[code sample](../src/practise/struct)



# reference types
## slice:
A slice has three components: **a pointer, a length, and a capacity** (of its underlying array)  

If you need to test whether a slice is empty, use len(s) == 0, not s == nil,because 
```go
var s []int // len(s) == 0, s == nil
s = nil  // len(s) == 0, s == nil
s = []int(nil) // len(s) == 0, s == nil
s = []int{}  // len(s) == 0, s != nil

var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, x...) // append the slice x
fmt.Println(x)  // "[1 2 3 1 2 3]"


z = make([]int, zlen, zcap)
copy(z, x)
copy(z[len(x):], y)
```
## map
In Go, a map is a reference to a hash table, and a map type is written map[K]V.

One reason that we can’t take the address of a map element is that growing a map might cause rehashing of existing elements into new storage locations, 

The zero value for a map type is nil, that is, a reference to no hashtable at all.

```go
/*对应的key不存在的话，map 也会返回一个zero value, 如果有歧义的话
可以用下面的方法来判断 */
if age, ok := ages["bob"]; !ok {
}

seen := make(map[string]bool) //曲线实现set类型
```    
## pointer  
we can read or update the value of a variable **indirectly** via a pointer, without using or even knowing the name of the variable, if indeed it has a name. pointer of 像是综合了**pointer of c 的写法和reference in C++的用法**. 

C++ References are often confused with pointers, but three major differences between references and pointers are 
   - You cannot have NULL references. You must always be able to assume that a reference is connected to a legitimate piece of storage.
   - Once a reference is initialized to an object, it cannot be changed to refer to another object.  Pointers can be pointed to another object at any time.
   - A reference must be initialized when it is created. Pointers can be initialized at any time.
  
## function 
define function type:
```go
package http // import "net/http"

type HandlerFunc func(ResponseWriter, *Request)
    The HandlerFunc type is an adapter to allow the use of ordinary functions as
    HTTP handlers. If f is a function with the appropriate signature,
    HandlerFunc(f) is a Handler that calls f.

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

```
[code sample](../src/practise/func)

[closure sample](../src/practise/closure)

# method
指定了reciever的function，跟c++的this 指针有点类似。 在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个package的。

# interface types

[code sample](../src/practise/interface)

# reflection
[code sample](../src/practise/reflect)

# unsafe
