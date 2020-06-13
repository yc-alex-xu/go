对于Go的版本管理主要用过 glide，下面介绍 Go 1.11 之后官方支持的版本管理工具 mod。

http://www.matthiassommer.it/programming/go-please-visual-studio-code-go-mod-go-language-server/

使用 Go Modules 后，理论上：
* 不用再定义 GOPATH 
* 代码可以随意放置,不再需要 src 目录下
* 不再需要 vendor 机制以及其他第 3 方 dep 工具
* 执行 Go 命令，不再需要指定 GOPATH

# setting
```bash
$ go env |grep GO11
GO111MODULE=""
```
可以用环境变量 GO111MODULE 开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是 auto。
* GO111MODULE=off 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
* GO111MODULE=on 模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。
* GO111MODULE=auto 在 $GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持。

# gfw issue
对于全世界绝大多数Gophers来说，Go module的引入带来的都是满满的幸福感，但是对于位于中国大陆地区的Gopher来说，在这种幸福感袭来的同时，也夹带了一丝“无奈”。其原因在于module-aware mode下，go tool默认不再使用传统GOPATH下或top vendor下面的包了，而是在GOPATH/pkg/mod(go 1.11中是这个位置，也许以后版本这个位置会变动)下面寻找Go module的local cache。

由于众所周知的原因，在大陆地区我们无法直接通过go get命令或git clone获取到一些第三方包，这其中最常见的就是golang.org/x下面的各种优秀的包。但是在传统的GOPATH mode下，我们可以先从golang.org/x/xxx的mirror站点github.com/golang/xxx上git clone这些包，然后将其重命名为golang.org/x/xxx。这样也能勉强通过开发者本地的编译。又或将这些包放入vendor目录并提交到repo中，也能实现正确的构建。

但是go module引入后，一旦工作在module-aware mode下，go build将不care GOPATH下或是vendor下的包，而是到GOPATH/pkg/mod查询是否有module的cache，如果没有，则会去下载某个版本的module，而对于golang.org/x/xxx下面的module，在大陆地区往往会get失败。

所以目前还是先
```bash
$ export GO111MODULE=off
$ go env |grep GO11
GO111MODULE="off"
```
