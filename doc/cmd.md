
# go build
这个命令主要用于编译代码。在包的编译过程中，若有必要，会同时编译与之相关联的包。由于src/practise下代码都比较简单，直接　
```bash
go run xxx.go
```
运行就可以了。

其他工程化设计：
* go build会忽略目录下以“_”或“.”开头的go文件。
* 如果你的源代码针对不同的操作系统需要不同的处理，那么你可以根据不同的操作系统后缀来命名文件。例如有一个读取数组的程序，它对于不同的操作系统可能有如下几个源文件：
array_linux.go array_darwin.go array_windows.go array_freebsd.go
go build的时候会选择性地编译以系统名结尾的文件（Linux、Darwin、Windows、Freebsd）。例如Linux系统下面编译只会选择array_linux.go文件，其它系统命名后缀文件全部忽略。

# go clean
这个命令是用来移除当前源码包和关联源码包里面编译生成的文件。我一般都是利用这个命令清除编译文件，然后GitHub递交源码，在本机测试的时候这些编译文件都是和系统相关的，但是对于源码管理来说没必要。

# go fmt
go强制了代码格式（比如左大括号必须放在行尾），不按照此格式的代码将不能编译通过，为了减少浪费在排版上的时间，go工具集中提供了一个go fmt命令 它可以帮你格式化你写好的代码文件，使你写代码的时候不需要关心格式，你只需要在写完之后执行go fmt <文件名>.go，你的代码就被修改成了标准格式

# go get
go get -u -v 可以自动更新包，而且当go get的时候会自动获取该包依赖的其他第三方包

go get本质上可以理解为首先第一步是通过源码工具clone代码到src下面，然后执行go install

# go install
这个命令在内部实际上分成了两步操作：第一步是生成结果文件(可执行文件或者.a包)，第二步会把编译好的结果移到$GOPATH/pkg或者$GOPATH/bin。

# go test
执行这个命令，会自动读取源码目录下面名为*_test.go的文件，生成并运行测试用的可执行文件。

# go tool
go tool下面下载聚集了很多命令，这里我们只介绍两个，fix和vet

    go tool fix . 用来修复以前老版本的代码到新版本，例如go1之前老版本的代码转化到go1,例如API的变化
    go tool vet directory|files 用来分析当前目录的代码是否都是正确的代码,例如是不是调用fmt.Printf里面的参数不正确，例如函数里面提前return了然后出现了无用代码之类的。

# go generate
这个命令是从Go1.4开始才设计的，用于在编译前自动化生成某类代码。go generate和go build是完全不一样的命令，通过分析源码中特殊的注释，然后执行相应的命令。这些命令都是很明确的，没有任何的依赖在里面。而且大家在用这个之前心里面一定要有一个理念，这个go generate是给你用的，不是给使用你这个包的人用的，是方便你来生成一些代码的。

这里我们来举一个简单的例子，例如我们经常会使用yacc来生成代码，那么我们常用这样的命令：

    go tool yacc -o gopher.go -p parser gopher.y
-o 指定了输出的文件名， -p指定了package的名称，这是一个单独的命令，如果我们想让go generate来触发这个命令，那么就可以在当前目录的任意一个xxx.go文件里面的任意位置增加一行如下的注释：

    //go:generate go tool yacc -o gopher.go -p parser gopher.y
这里我们注意了，//go:generate是没有任何空格的，这其实就是一个固定的格式，在扫描源码文件的时候就是根据这个来判断的。

所以我们可以通过如下的命令来生成，编译，测试。如果gopher.y文件有修改，那么就重新执行go generate重新生成文件就好。

    $ go generate
    $ go build
    $ go test    

# go doc
想当与c语言开发用的man

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

$ go doc fmt println
package fmt // import "fmt"

func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.
$ go doc builtin string
package builtin // import "builtin"

type string string
    string is the set of all strings of 8-bit bytes, conventionally but not
    necessarily representing UTF-8-encoded text. A string may be empty, but not
    nil. Values of string type are immutable.


$ go doc image/color  Color
package color // import "image/color"

type Color interface {
	// RGBA returns the alpha-premultiplied red, green, blue and alpha values
	// for the color. Each value ranges within [0, 0xffff], but is represented
	// by a uint32 so that multiplying by a blend factor up to 0xffff will not
	// overflow.
	//
	// An alpha-premultiplied color component c has been scaled by alpha (a),
	// so has valid values 0 <= c <= a.
	RGBA() (r, g, b, a uint32)
}
    Color can convert itself to alpha-premultiplied 16-bits per channel RGBA.
    The conversion may be lossy.

$ go doc io.Copy
package io // import "io"

func Copy(dst Writer, src Reader) (written int64, err error)
    Copy copies from src to dst until either EOF is reached on src or an error
    occurs. It returns the number of bytes copied and the first error
    encountered while copying, if any.

    A successful Copy returns err == nil, not err == EOF. Because Copy is
    defined to read from src until EOF, it does not treat an EOF from Read as an
    error to be reported.

    If src implements the WriterTo interface, the copy is implemented by calling
    src.WriteTo(dst). Otherwise, if dst implements the ReaderFrom interface, the
    copy is implemented by calling dst.ReadFrom(src).



```
go doc 工具会从 Go 程序和包文件中提取顶级声明的首行注释以及每个对象的相关注释，并生成相关文档。它也可以作为一个提供在线文档浏览的 web 服务器，golang.org 就是通过这种形式实现的。

# go help
这是对go 命令的帮助信息

```bash
$ go help get
usage: go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]

Get downloads the packages named by the import paths, along with their
dependencies. It then installs the named packages, like 'go install'.

$ go help packages
Many commands apply to a set of packages:

	go action [packages]

To make common patterns more convenient, there are two special cases.
First, /... at the end of the pattern can match an empty string,
so that net/... matches both net and packages in its subdirectories, like net/http.
Second, any slash-separated pattern element containing a wildcard never
participates in a match of the "vendor" element in the path of a vendored
package, so that ./... does not match packages in subdirectories of
./vendor or ./mycode/vendor, but ./vendor/... and ./mycode/vendor/... do.
Note, however, that a directory named vendor that itself contains code
is not a vendored package: cmd/vendor would be a command named vendor,
and the pattern cmd/... matches it.
See golang.org/s/go15vendor for more about vendoring.

```
# install after gfw
```bash
go/src/golang.org/x$ git clone https://github.com/golang/tools.git
$ go install golang.org/x/tools/cmd/godoc
$ go install golang.org/x/tools/cmd/guru
$ go install golang.org/x/tools/cmd/gorename
$ go install golang.org/x/tools/cmd/fiximports
$ go install golang.org/x/tools/cmd/godex
```

# godoc
由于https://golang.org/国内没法访问，只能启动godoc

```bash
$ godoc -http=:8080 &
[1] 25294
```
然后在浏览器中可以看到

![godoc in browser](images/godoc_http.png)

这可以说是godoc 非常有趣的一点，自己写的程序跟builtin同等地位，godoc 就是将　GOPATH，GOROOT下的所有exported 的注释显示出来。

当然,有的包只是提供注释
```bash
alex@minipc:/usr/lib/go/src/builtin$ cat builtin.go 
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package builtin provides documentation for Go's predeclared identifiers.
	The items documented here are not actually in package builtin
	but their descriptions here allow godoc to present documentation
	for the language's special identifiers.
*/
package builtin
```
# gopl
```bash
go/src/golang.org/x$ git clone https://github.com/golang/sync
$ go get  -v golang.org/x/tools/gopls
```
![setup in vs code](images/vscode_gopls.png)

# objdum
```bash
 $ go build main.go 
 $ go tool objdump main > obj.s
 $ cat obj.s
```

# go test
Go语言中自带有一个轻量级的测试框架testing和自带的go test命令会自动读取源码目录下面名为*_test.go的文件，生成并运行测试用的可执行文件。默认的情况下，不需要任何的参数，它会自动把你源码包下面所有test文件测试完毕，当然你也可以带上参数，详情请参考go help testflag

[rbt & its testing](https://github.com/yc-alex-xu/go/tree/master/src/practise/rbtree)

```bash
$ go test
PASS
ok  	practise/rbtree	0.002s
```
## 3rd party: gotest
gotests插件自动生成测试代码:

```bash
go get -u -v github.com/cweill/gotests/...
```
安装后，VS code 就可以选定一个 function 然后让他自动生成 unit test case 和benchmark test；test case 也可以在VS code 中选择执行。

# profiler
Go 提供了一个低级的分析 API runtime/pprof ，但如果你在开发一个长期运行的服务，使用更高级的 net/http/pprof 包会更加便利。你只需要在代码中加入 import _ "net/http/pprof" ，它就会自动注册所需的 HTTP 处理器（Handler） 。
