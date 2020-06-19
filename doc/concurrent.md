# goroutine
In Go, each concurrently executing activity is called a **goroutine**. When a program starts, its only goroutine is the one that calls the main function, so we call it the main goroutine. New goroutines are created by the **go statement**. There is no programmatic way for one goroutine to stop another, but as we will see later, there are ways to communicate with a goroutine to request that it stop itself.


[example code](https://github.com/yc-alex-xu/go/blob/master/src/practise/goroutine/main.go)
* main.go 两个goroutine 协调处理的情况
* goodbye.go  主goroutine先结束的情况
* mutex.go    在goodbye.go基础上加了mutex, 保证主goroutine等副goroutine完成再退出
* sync.go     mutex.go 基础上用两个mutex实现了goroutine按计划顺序执行

# channels

If goroutines are the activities of a concurrent Go program, **channels** are the connections between them.

```go
ch = make(chan int) // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```
A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel.When a value is sent on an unbuffered channel, the receipt of the value happens before the reawakening of the sending goroutine.

channel通过操作符 **<-** 来接收和发送数据

默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所谓阻塞，也就是如果读取（v := <-ch）它将会被阻塞，直到buffer不为空。其次，任何发送（ch<-v）将会被阻塞，直到buffer不为满。无缓冲channel是在多个goroutine之间同步很棒的工具。

## Buffered Channels
设置一个buffer size, (Alex, 所谓unbuffered Channels 的buffer size 应该是0???)

[code](https://github.com/yc-alex-xu/go/blob/master/src/practise/channels/)
* original.go:没有考虑同步问题，goroutine的运行次序是不确定的。
* sync.go： 在original基础上做了同步，goroutine的运行次序是确定的。
* buffered.go:　在original基础上改为buffered channel, 并且用了range 来收channel消息。注意修改不慎重的话容易deadlock

# Select
[example code](https://github.com/yc-alex-xu/go/tree/master/src/practise/select)
* main.go： original code
* default.go: 用default 改写
* timeout.go: 避免整个程序进入阻塞的情况
 

# race condition