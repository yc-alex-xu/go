# go get
Get downloads the packages named by the import paths, along with their dependencies. It then installs the named packages, like 'go install'.e.g. go get -u -v 

Without that flag -u, packages that already exist locally will not be updated.

# go build
用于编译package main,其他package 只是检查一下编译错误。

其他工程化设计：
* go build会忽略目录下以“_”或“.”开头的go文件。
* 如果你的源代码针对不同的操作系统需要不同的处理，那么你可以根据不同的操作系统后缀来命名文件。例如有一个读取数组的程序，它对于不同的操作系统可能有如下几个源文件：
array_linux.go array_darwin.go array_windows.go array_freebsd.go
go build的时候会选择性地编译以系统名结尾的文件（Linux、Darwin、Windows、Freebsd）。例如Linux系统下面编译只会选择array_linux.go文件，其它系统命名后缀文件全部忽略。

# go install
适用于main以外的pacakge,
The go install command is very similar to go build, except that it saves the compiled code for each package and command instead of throwing it away. Compiled packages are saved beneath the $GOPATH/pkg directory corresponding to the src directory in which the source resides, and command executables are saved in the $GOPATH/bin directory. 

to cross-compile a Go program,Just set the GOOS or GOARCH env variables during the build.

#  go list
reports information about available packages


# go doc
类似 linux/c用的man

```bash
$ go doc builtin
$ go doc fmt println
$ go doc image/color  Color
$ go doc io.Copy

```
go doc 工具会从 Go 程序和包文件中提取顶级声明的首行注释以及每个对象的相关注释，并生成相关文档。golang.org 就是通过这种形式实现的。

# go test
Go语言中自带有一个轻量级的测试框架testing和自带的go test命令会自动读取源码目录下面名为*_test.go的文件，生成并运行测试用的可执行文件。默认的情况下，不需要任何的参数，它会自动把你源码包下面所有test文件测试完毕，当然你也可以带上参数，详情请参考go help testflag

Within *_test.go files, three kinds of functions are treated specially : 
* tests, 
* benchmarks, and
* examples. 

[rbt & its testing](https://github.com/yc-alex-xu/go/tree/master/src/practise/rbtree)

```bash
$ go test
PASS
ok  	practise/rbtree	0.002s
```
The go test tool has built-in support for several kinds of profiling.
```bash
$ go test -cpuprofile=cpu.out
$ go test -blockprofile=block.out
$ go test -memprofile=mem.out
$ go tool pprof -text -nodecount=10 ./http.test cpu.log
```

## 3rd party: gotest
gotests插件自动生成测试代码:

```bash
go get -u -v github.com/cweill/gotests/...
```
安装后，VS code 就可以选定一个 function 然后让他自动生成 unit test case 和benchmark test；test case 也可以在VS code 中选择执行。

# go clean
这个命令是用来移除当前源码包和关联源码包里面编译生成的文件。我一般都是利用这个命令清除编译文件，然后GitHub递交源码，在本机测试的时候这些编译文件都是和系统相关的，但是对于源码管理来说没必要。

# go fmt
go强制了代码格式（比如左大括号必须放在行尾），不按照此格式的代码将不能编译通过，为了减少浪费在排版上的时间，go工具集中提供了一个go fmt命令 它可以帮你格式化你写好的代码文件，使你写代码的时候不需要关心格式，你只需要在写完之后执行go fmt <文件名>.go，你的代码就被修改成了标准格式

# go tool
go tool下面下载聚集了很多命令，这里我们只介绍两个，fix和vet

    go tool fix . 用来修复以前老版本的代码到新版本，例如go1之前老版本的代码转化到go1,例如API的变化
    go tool vet directory|files 用来分析当前目录的代码是否都是正确的代码,例如是不是调用fmt.Printf里面的参数不正确，例如函数里面提前return了然后出现了无用代码之类的。

# go generate
这个命令是从Go1.4开始才设计的，用于在编译前自动化生成某类代码。go generate和go build是完全不一样的命令，通过分析源码中特殊的注释，然后执行相应的命令。这些命令都是很明确的，没有任何的依赖在里面。而且大家在用这个之前心里面一定要有一个理念，这个go generate是给你用的，不是给使用你这个包的人用的，是方便你来生成一些代码的。


# go help
这是对go 命令的帮助信息

```bash
$ go help get
usage: go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]

$ go help packages
Many commands apply to a set of packages:

	go action [packages]

```
# install some cmd after gfw
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
