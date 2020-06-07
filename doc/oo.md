# struct

[例子](https://github.com/yc-alex-xu/go/tree/master/src/practise/struct)
* main.go　　定义type person struct
* anonym.go  嵌套了匿名field，解决了OO继承问题
* override.go   Go规定外层的优先访问，也就是当你通过student.phone访问的时候，是访问student里面的字段，而不是human里面的字段，而human里面的字段可以通过非匿名的方式访问。


# method
method是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在func后面增加了一个receiver(也就是method所依从的主体)。
```go
func (r ReceiverType) funcName(parameters) (results)
```
用Rob Pike的话来说就是：
> "A method is a function with an implicit first argument, called a receiver."  
跟C++中method 隐含带this 指针一致。

在使用method的时候重要注意几点
* Receiver可以是以值传递，还可以是指针, 两者的差别在于, 指针作为Receiver会对实例对象的内容发生操作,而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作。
* method里面可以通过 **.** 访问接收者的字段
* Polymorphism: 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样

[code](https://github.com/yc-alex-xu/go/tree/master/src/practise/method)
* main.go：　任何的自定义类型中定义任意多的method
* inherit.go 　如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。如果想要实现自己的method,怎么办？可以override.

# interface
[example code](https://github.com/yc-alex-xu/go/tree/master/src/practise/interface)
* main.go
* empty.go：  空interface 的用法
* stringer.go:　fmt.Println 用到的Stringer interface 类型参数

可以说interface概念与java没有啥区别，但用法完全不同，go是通过它给各种数据类型分类。打破了其他语言中同一base class的才能归为一类的限制。interface可以被任意的对象实现。我们看到上面的Men interface被Human、Student和Employee实现。同理，一个对象可以实现任意多个interface，例如上面的Student实现了Men和YoungChap两个interface。

**空interface**不包含任何的method，正因为如此，所有的类型都实现了空interface。它有点类似于C语言的void*类型 (alex:但它可以代表值类型和引用类型)。

interface作为参数
```bash
$ godoc fmt Println
use 'godoc cmd/fmt' for documentation on the fmt command 

func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
    
```
**interface的实现**：作为一门编程语言，对方法的处理一般分为两种类型：一是将所有方法组织在一个表格里，静态地调用（C++, java）；二是调用时动态查找方法(python, smalltalk, js)。
而go语言是两者的结合：虽然有table，但是是需要在运行时计算的table。

e.g. Binary类实现了两个方法，String()和Get(), 由于它实现了String()，按照golang的隐式方法实现来看，Binary satisfied了Stringer接口。因此它可以赋值： s:=Stringer(b)。

s 由两个指针组成，
* tab：指向一个interface table，叫 itable。itable开头是一些描述类型的元字段，后面是一串method。注意这个方法是interface本身的方法，并非其dynamic value的方法。但这个方法的实现肯定是具体类的方法，这里就是Binary的方法。当这个interface无method时，itable可以省略，直接指向一个type即可。s.(type) 相当于 s.tab->type
* data： 指向dynamic value的一个拷贝，这里则是b的一份拷贝。也就是，给interface赋值时，会在堆上分配内存，用于存放拷贝的值。同样，当值本身只有一个字长时，这个指针也可以省略。
 
![interface internal](images/interface_internal.png) 


## interface的嵌入
e.g. io包下面的 io.ReadWriter ，它包含了io包下面的Reader和Writer两个interface,也就将它们所有method给隐式的包含进来了

```go
// io.ReadWriter
type ReadWriter interface {
	Reader
	Writer
}
```

# reflect
所谓反射就是能检查程序在运行时的状态。我们一般用到的包是reflect包。

[example of reflect](https://github.com/yc-alex-xu/go/blob/master/src/practise/reflect/main.go)