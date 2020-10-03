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
GOPATH这个目录用来存放Go源码，Go的可运行文件，以及相应的编译之后的包文件。所以这个目录下面有三个子目录：src、bin、pkg
* src 存放源代码（比如：.go .c .h .s等）
* pkg 编译后生成的文件（比如：.a）
* bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）

# gfw issue
golang 在 github 上建立了一个镜像库，如：https://golang.org/x/net 的镜像库，在GitHub上对应的是https://github.com/golang/net.

```bash
$ mkdir -p golang.org/x/t
$ cd golang.org/x
$ git clone https://github.com/golang/tools
$ git clone https://github.com/golang/net
$ git clone https://github.com/golang/lint
$ git clone https://github.com/golang/xerrors
$ git clone https://github.com/golang/mod
$ git clone https://github.com/golang/sync

$ go install golang.org/x/tools/cmd/godoc
$ go install golang.org/x/tools/cmd/guru
$ go install golang.org/x/tools/cmd/gorename
$ go install golang.org/x/tools/cmd/fiximports
$ go install golang.org/x/tools/cmd/godex
```
# install goreturns/debug
```bash
$ go get -v -u github.com/sqs/goreturns
$ go get -v github.com/ramya-rao-a/go-outline


$ go get github.com/go-delve/delve/cmd/dlv
$ dlv
Delve is a source level debugger for Go programs.
```

