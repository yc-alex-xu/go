# pkg testing
Go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，testing框架和其他语言中的测试框架类似，你可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例，那么接下来让我们一一来看一下怎么写。

[rbt & its testing](https://github.com/yc-alex-xu/go/tree/master/src/practise/rbtree)

```bash
$ go test
PASS
ok  	practise/rbtree	0.002s
```
# gotest
gotests插件自动生成测试代码:

```bash
go get -u -v github.com/cweill/gotests/...
```
安装后，VS code 就可以选定一个 function 然后让他自动生成 unit test case 和benchmark test；test case 也可以在VS code 中选择执行。


