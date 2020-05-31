# 借鉴其他语言
* array、slice、map
* make、new操作
* defer
* Panic和Recover
* init函数
* smart pointer ????
* 函数： 保留c语言的不定数量的参数；可以直接返回多个值，不用定义tuple
* channel 既是goroutine 之间的的message　channel,又是Semaphore

# go 相比c/c++,自带工程属性，如
1. 取消了头文件。 alex:既然可以import,应该是用.a文件的信息来替代头文件的函数原型检查功能
2. 似乎没有动态链接库概念
3. 1,2 的本源应该，go 不想像C++一样解决所有的问题，而是把某一应用领域的问题简化开发，就像google推多种语言一样。
4. 目录组织和package 来解决源码组织问题， 物理文件可以多个，但是只要package 名相同，他们就是同一逻辑文件
5. 自带工具： go fmt 解决code style；  godoc 解决 help 问题
6. TDD：go test, 会自动读取源码目录下面名为*_test.go的文件，生成并运行测试用的可执行文件。默认的情况下，不需要任何的参数，它会自动把你源码包下面所有test文件测试完毕，当然你也可以带上参数，详情请参考go help testflag


缺点：
* 从 0.9 版开始学习go,变化颇大，并且不兼容，以前写的程序现在已经不能运行。

# 语法检查严格
* 变量定义了未使用，包引入了未使用　　　　这些在c++中warning 都不算，但在go 就是error了。
* Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错,所以只能使用**_**来丢弃不需要的返回值
* exported type Xyyyy should have comment or be unexported

## 例外： 指针的自动转换
* 如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
* 如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method
所以，你不用担心你是调用的指针的method还是不是指针的method，Go知道你要做的一切，这对于有多年C/C++编程经验的同学来说，真是解决了一个很大的痛苦。
alex： 这解释成C++ 中的“引用”是不是更好？ go似乎没有指针概念,至少没有指针地址概念，设计者应该也不鼓励研究内部实现所以不暴露地址。编程效率高的语言总是强调思路，把实现留给语言开发人员。

# 谨慎
* new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值. 而C/C++是不浪费时间做零值填充的，除非你定义了构造函数。
*
 

# Object Oriented
* 不用“class”,就用“struct” 即可。
* 没有任何的私有、公有关键字，通过大小写来实现(大写开头的为公有，小写开头的为私有)，方法也同样适用这个原则。
* 继承,C++ 中继承是用父类之类的概念，而go 中直接的多，是用struct B 包含匿名的 struct B,C的方法，表示struct B 包含了struct B,C的所有field和方法。

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
* interface: 可以说它的概念与java没有啥区别，但用法完全不同，go是通过它给各种数据类型分类。打破了其他语言中按base class的归类的限制。另外，感觉它就起了指针的作用。
  
# reflect
体现的是go 结余动态和非动态语言之间的特点，估计是通过添加metadata实现的。

# code style
* 命名还是很短，如
```go
func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		rw, e := l.Accept()
		//省略出错处理
		tempDelay = 0
		c, err := srv.newConn(rw)
		if err != nil {
			continue
		}
		go c.serve()		
	}	
}
```	
* "message": "don't use underscores in Go names; var token_saved should be tokenSaved",
