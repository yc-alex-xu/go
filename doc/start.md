# install go

```bash
$ sudo apt-get install python-software-properties
$ sudo add-apt-repository ppa:gophers/go
$ sudo apt-get update
$ sudo apt-get install golang git-core mercurial

$ cat /etc/profile
export GOPATH=/home/alex/go
export PATH=$PATH:/home/alex/go/bin
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN

$ source /etc/profile 
```
GOPATH这个目录用来存放Go源码，Go的可运行文件，以及相应的编译之后的包文件。所以这个目录下面有三个子目录：src、bin、pkg
* src 存放源代码（比如：.go .c .h .s等）
* pkg 编译后生成的文件（比如：.a）
* bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）






