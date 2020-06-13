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