No time waste on env issue, so as simple as possible

# install
https://studygolang.com/dl

# configure env
```bash
$ cat /etc/profile
export GOPATH=/home/alex/go
export PATH=$PATH:/home/alex/go/bin
$ source /etc/profile
$ env |grep GO
GOPATH=/home/alex/go
$ go env
GOARCH="amd64"
```
now all go code will put under folder ~/go




# The test program 
```bash
src$ vi hello.go
src$ go build hello.go 
src$ ./hello 
hello, world
src$ 
```
GOPATH这个目录用来存放Go源码，Go的可运行文件，以及相应的编译之后的包文件。所以这个目录下面有三个子目录：src、bin、pkg
* src 存放源代码（比如：.go .c .h .s等）
* pkg 编译后生成的文件（比如：.a）
* bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）

# gfw issue
golang 在 github 上建立了一个镜像库，如：https://golang.org/x/net 的镜像库，在GitHub上对应的是https://github.com/golang/net.

```bash
src/golang.org/x$ git clone https://github.com/golang/tools
src/golang.org/x$ git clone https://github.com/golang/lint
```
## glide
用于mirror
```bash
alex@minipc:~$ go get -u github.com/xkeyideal/glide
alex@minipc:~$ glide
```
## github.com/sqs/goreturns failure
```bash
$ go get -v -u github.com/sqs/goreturns
github.com/sqs/goreturns (download)
Fetching https://golang.org/x/tools/imports?go-get=1
https fetch failed: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: i/o timeout

$ code //restart the vs code via terminal
Installing github.com/sqs/goreturns SUCCEEDED
```

# vs code
* https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code
* https://github.com/golang/vscode-go

![install the tools from　menu "command palette" ](images/vscode_install_tools.png)

# install debug
```bash
alex@minipc:~$ go get github.com/go-delve/delve/cmd/dlv
alex@minipc:~$ dlv
Delve is a source level debugger for Go programs.
```

