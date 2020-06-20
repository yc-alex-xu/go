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

# predeclared 
```go
Constants: true false iota nil
Types: int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex128 complex64  bool byte rune string error
Functions: make len cap new append copy close delete complex real imag panic recover

    p := new(int)
    delete(p)

```
# bitwise binary operators
```go
&   AND
|   OR
^   XOR
&^  bit clear (AND NOT)
<<  left shift
>>  right shift
```

# Printf format

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
```go
	var r rune = '国'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
```   


# type
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
# const
The value of a constant must be basic type: a number, string , or boolean. e.g.
```go
true false iota nil
```
**Untyped Constants**: The compiler represents these uncommitted constants with much greater numeric precision than values of basic types, and arithmetic on them is more precise than machine arithmetic; you may assume at least 256 bits of precision. alex:也就是说，它们的精度太大，无法用basic type的空间来存储，所以compiler就给它们一个特例,告诉programmer你就不用操心它的data type了.

[code sample](../src/practise/const)


# array:
array as fuction parameter: When a function is called, a copy of each argument value is assigned to the corresponding parameter variable, so the function receives a copy, not the original. C/C++这时传的是地址

# slice:
If you need to test whether a slice is empty, use len(s) == 0, not s == nil,because 
```go
var s []int // len(s) == 0, s == nil
s = nil  // len(s) == 0, s == nil
s = []int(nil) // len(s) == 0, s == nil
s = []int{}  // len(s) == 0, s != nil
```

# pointer  
we can read or update the value of a variable **indirectly** via a pointer, without using or even knowing the name of the variable, if indeed it has a name. pointer of 像是综合了**pointer of c 的写法和reference in C++的用法**. 

C++ References are often confused with pointers, but three major differences between references and pointers are 
   - You cannot have NULL references. You must always be able to assume that a reference is connected to a legitimate piece of storage.
   - Once a reference is initialized to an object, it cannot be changed to refer to another object.  Pointers can be pointed to another object at any time.
   - A reference must be initialized when it is created. Pointers can be initialized at any time.
  

# function 
define function type:
```go
package http // import "net/http"

type HandlerFunc func(ResponseWriter, *Request)
    The HandlerFunc type is an adapter to allow the use of ordinary functions as
    HTTP handlers. If f is a function with the appropriate signature,
    HandlerFunc(f) is a Handler that calls f.

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

```


# local variables 
* have dynamic lifetimes。 A compiler may choose to allocate local variables on the heap or on the stack but, perhaps surprisingly, this choice is not detemined by whether var or new was used to declare the variable. 也就是说，程序员不用操心，一个func可以返回其locall variable,放心用，gc不会愚蠢的recycle其占用空间。
* Garbage collection is a tremendous help in writing correct programs, but it does not relieve you of the burden of thinking about memory. You don’t need to explicitly allocate and free memory, but to write efficient programs you still need to be aware of the lifet ime of variables.For example, keeping unnecessary pointers to short-lived objects within long-lived objects,especially global variables, will prevent the garbage collector from reclaiming the short-lived objects.

its scope:

![a good exmpale of scope](images/scope.png)
