
# go build
这个命令主要用于编译代码。在包的编译过程中，若有必要，会同时编译与之相关联的包。

可以说主要用于编译package main的程序。　对于我来说，大多练习代码都比较简单，直接　
```bash
go run xxx.go
```
就可以了。

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
go get -u 参数可以自动更新包，而且当go get的时候会自动获取该包依赖的其他第三方包

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

# godoc
```bash
$ sudo apt install golang-golang-x-tools

```

如何查看相应package的文档呢？ 例如builtin包，那么执行godoc builtin 如果是http包，那么执行godoc net/http 查看某一个包里面的函数，那么执行godoc fmt Printf 也可以查看相应的代码，执行godoc -src fmt Printf

通过命令在命令行执行 godoc -http=:端口号 比如godoc -http=:8080。然后在浏览器中打开127.0.0.1:8080，你将会看到一个golang.org的本地copy版本，通过它你可以查询pkg文档等其它内容。如果你设置了GOPATH，在pkg分类下，不但会列出标准包的文档，还会列出你本地GOPATH中所有项目的相关文档，这对于经常被墙的用户来说是一个不错的选择。    