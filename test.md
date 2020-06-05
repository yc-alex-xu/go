# pkg testing
Go语言中自带有一个轻量级的测试框架testing和自带的go test命令会自动读取源码目录下面名为*_test.go的文件，生成并运行测试用的可执行文件。默认的情况下，不需要任何的参数，它会自动把你源码包下面所有test文件测试完毕，当然你也可以带上参数，详情请参考go help testflag

[rbt & its testing](https://github.com/yc-alex-xu/go/tree/master/src/practise/rbtree)

```bash
$ go test
PASS
ok  	practise/rbtree	0.002s
```
# 3rd party: gotest
gotests插件自动生成测试代码:

```bash
go get -u -v github.com/cweill/gotests/...
```
安装后，VS code 就可以选定一个 function 然后让他自动生成 unit test case 和benchmark test；test case 也可以在VS code 中选择执行。


