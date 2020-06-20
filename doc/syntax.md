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


# const
The value of a constant must be a number, string , or boolean.   
```go
true false iota nil
```
[code sample](../src/practise/const)

# predeclared 
```go
Constants: true false iota nil
Types: int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex128 complex64  bool byte rune string error
Functions: make len cap new append copy close delete complex real imag panic recover

    p := new(int)
    delete(p)

```

# typo
```go
type name underlying-type
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
# 