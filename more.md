# start from godoc
由于https://golang.org/国内没法访问，只能启动godoc
```bash
$ godoc -http=:8080 &
[1] 25294
```
然后在浏览器中可以看到
![godoc in browser](images/godoc_http.png)

这可以说是godoc 非常有趣的一点，自己写的程序跟builtin同等地位，godoc 就是将　GOPATH，GOROOT下的所有exported 的注释显示出来。

可能go设计时有代码、文档在一起的思想。

```bash
$ go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/alex/.cache/go-build"
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/alex/base/go"
GORACE=""
GOROOT="/usr/lib/go-1.10"

$ ls /usr/lib/go-1.10/src
all.bash          bootstrap.bash  bytes       cmp.bash   crypto    errors  go     index         log        math           net     race.bash  run.bash  strconv  testing  unsafe
all.bat           bufio           clean.bash  compress   database  expvar  hash   internal      make.bash  mime           os      race.bat   run.bat   strings  text     vendor
androidtest.bash  buildall.bash   clean.bat   container  debug     flag    html   io            make.bat   naclmake.bash  path    reflect    runtime   sync     time
archive           builtin         cmd         context    encoding  fmt     image  iostest.bash  Make.dist  nacltest.bash  plugin  regexp     sort      syscall  unicode
```
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

// bool is the set of boolean values, true and false.
type bool bool

// true and false are the two untyped boolean values.
const (
	true  = 0 == 0 // Untyped bool.
	false = 0 != 0 // Untyped bool.
)

// uint8 is the set of all unsigned 8-bit integers.
// Range: 0 through 255.
type uint8 uint8

// uint16 is the set of all unsigned 16-bit integers.
// Range: 0 through 65535.
type uint16 uint16

```
有些则提供了真实的go代码。
```bash
alex@minipc:/usr/lib/go/src/net/http$ ls
```


# Package sync 
provides basic synchronization primitives such as mutual exclusion locks. Other than the Once and WaitGroup types, most are intended for use by low-level library routines. Higher-level synchronization is better done via channels and communication.

# Package net 
provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.

Although the package provides access to low-level networking primitives, most clients will need only the basic interface provided by the Dial, Listen, and Accept functions and the associated Conn and Listener interfaces. The crypto/tls package uses the same interfaces and similar Dial and Listen functions.

# Package html 
provides functions for escaping and unescaping HTML text.

# Package encoding 
defines interfaces shared by other packages that convert data to and from byte-level and textual representations. Packages that check for these interfaces include encoding/gob, encoding/json, and encoding/xml. As a result, implementing an interface once can make a type useful in multiple encodings. Standard types that implement these interfaces include time.Time and net.IP. The interfaces come in pairs that produce and consume encoded data.

# Package os 
provides a platform-independent interface to operating system functionality. The design is Unix-like, although the error handling is Go-like; failing calls return values of type error rather than error numbers. Often, more information is available within the error. For example, if a call that takes a file name fails, such as Open or Stat, the error will include the failing file name when printed and will be of type *PathError, which may be unpacked for more information.

The os interface is intended to be uniform across all operating systems. Features not generally available appear in the system-specific package syscall.

# Package sort 
provides primitives for sorting slices and user-defined collections.