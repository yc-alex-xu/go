# goroutine
goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。
```go
go hello(a, b, c)
```
通过关键字go就启动了一个goroutine。

[example code](https://github.com/yc-alex-xu/go/blob/master/src/practise/goroutine/main.go)
* main.go 两个goroutine 协调处理的情况
* goodbye.go  主goroutine先结束的情况
* mutex.go    在goodbye.go基础上加了mutex, 保证主goroutine等其他goroutine完成再退出
* sync.go     mutex.go 基础上用两个mutex实现了goroutine按计划顺序执行

goroutine 遇到下面的情况下可能会产生重新调度（大家判断哪些代码属于下面这些情况）：
* 阻塞 I/O
* select操作
* 阻塞在channel
* 等待锁
* 主动调用 runtime.Gosched()

goroutine 的抢占

如果一个 goroutine 不包含上面提到的几种情况，那么其他的 goroutine 就无法被调度到相应的 CPU 上面运行，这是不应该发生的。这时候就需要抢占机制来打断长时间占用 CPU 资源的 goroutine ，发起重新调度。Golang 运行时（runtime）中的系统监控线程 sysmon 可以找出“长时间占用”的 goroutine，从而“提醒”相应的 goroutine 该中断了。

note:
* sysmon 在独立的 M（线程）中运行，且不需要绑定 P。这意味着， runtime.GOMAXPROCS(1) 限制 P 的数量为 1 的情况下，即使一个 goroutine 一直占用这个 P 进行密集型计算（意味着 goroutine 一直占有唯一的 P），依然不影响 sysmon 的正常运行。

* sysmon 可以找到“长时间占用 P”的 goroutine，但也只是标记 goroutine 应该被抢占了，并无法强制进行 goroutine 的切换。因此本文的 “第二段代码” 在进行 for 循环时并不会被抢占，而是在 for 循环结束后执行 fmt.Printf("sum: %d\n", sum) 的时候才被抢占（因为 for 循环里没有被插入抢占检查点，也就是说抢占检查点是编译器预先插入的，在非内联的函数的前面，具体可以查看最后几篇参考文章）。

# channels
多个goroutine运行在同一个进程里面，共享内存数据，不过需要解决PV等同步问题，

Go提供了一个很好的通信机制channel。channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel。　

channel通过操作符 **<-** 来接收和发送数据

默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所谓阻塞，也就是如果读取（v := <-ch）它将会被阻塞，直到buffer不为空。其次，任何发送（ch<-v）将会被阻塞，直到buffer不为满。无缓冲channel是在多个goroutine之间同步很棒的工具。


## Buffered Channels
设置一个buffer size, (Alex, 所谓unbuffered Channels 的buffer size 应该是1 )

[code](https://github.com/yc-alex-xu/go/blob/master/src/practise/channels/)
* original.go:没有考虑同步问题，goroutine的运行次序是不确定的。
* main.go： 在original基础上做了同步，goroutine的运行次序是确定的。
* buffered.go:　在original基础上改为buffered channel, 并且用了range 来收channel消息。注意修改不慎重的话容易deadlock

# Select
我们上面介绍的都是只有一个channel的情况，那么如果存在多个channel的时候，我们该如何操作呢，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。

select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是**随机**的选择一个执行的。

[example code](https://github.com/yc-alex-xu/go/tree/master/src/practise/select)
* main.go： original code
* default.go: 用default 改写
* timeout.go: 避免整个程序进入阻塞的情况

  
# sumary
CSP模型是上个世纪七十年代提出的，用于描述两个独立的并发实体通过共享的通讯 channel(管道)进行通信的并发模型。 CSP中channel是第一类对象，它不关注发送消息的实体，而关注与发送消息时使用的channel。

Golang 就是借用CSP模型的一些概念为之实现并发进行理论支持，其实从实际上出发，go语言并没有，完全实现了CSP模型的所有理论，仅仅是借用了 process和channel这两个概念。process是在go语言上的表现就是 goroutine 是实际并发执行的实体，每个实体之间是通过channel通讯来实现数据共享。
 * channel 本质上是消息队列，缺省buffer size 为 1;　执行读或写都会因为buffer为空或者满而阻塞。
 * goroutine 是提供给开发人员的接口，底层还是要映射到OS识别的thread.

