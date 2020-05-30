# goroutine
goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。

[example code](https://github.com/yc-alex-xu/go/blob/master/src/practise/goroutine/main.go)

# channels
多个goroutine运行在同一个进程里面，共享内存数据，不过设计上我们要遵循：不要通过共享来通信，而要通过通信来共享。Go提供了一个很好的通信机制channel。channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：

默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。


## Buffered Channels
设置一个buffer size, 省的sender 因为buffer不为空而block.

[code](https://github.com/yc-alex-xu/go/blob/master/src/practise/channels/)
* original.go:没有考虑同步问题  
* main.go： 在original基础上做了同步
* buffered.go:　在original基础上改为buffered channel, 并且用了range 来收channel消息。注意修改不慎重的话容易deadlock
```bash
$ go run original.go 
```

# Select
我们上面介绍的都是只有一个channel的情况，那么如果存在多个channel的时候，我们该如何操作呢，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。

select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是**随机**的选择一个执行的。

[example code](https://github.com/yc-alex-xu/go/tree/master/src/practise/select)
* main.go： original code
* default.go: 用default 改写
* timeout.go: 避免整个程序进入阻塞的情况

# runtime goroutine
runtime包中有几个处理goroutine的函数：

* Goexit 退出当前执行的goroutine，但是defer函数还会继续调用
* Gosched 让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
* NumCPU 返回 CPU 核数量
* NumGoroutine 返回正在执行和排队的任务总数
* GOMAXPROCS 用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
  