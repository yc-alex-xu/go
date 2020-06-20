When we cannot confidently say that one event happens before the other, then the events x and y are concurrent.

Go mantra (咒语) 
> Do not communicate by sharing memory; instead, share memory by communicating.

In Go, each concurrently executing activity is called a **goroutine**. When a program starts, its only goroutine is the one that calls the main function, so we call it the **main goroutine**. New goroutines are created by the **go statement**. There is no programmatic way for one goroutine to stop another, but as we will see later, there are ways to communicate with a goroutine to request that it stop itself.

[example code](https://github.com/yc-alex-xu/go/blob/master/src/practise/goroutine/)

# channels
If goroutines are the activities of a concurrent Go program, **channels** are the connections between them.

```go
ch = make(chan int) // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```
A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel.alex: 数据已经发到channel,就是没有程序读的话就不返回。 When a value is sent on an unbuffered channel, the receipt of the value happens **before** the reawakening of the sending goroutine.alex:并且被block的goroutine是放在一个队列里，被唤醒的顺序不会乱。

[direction of channel](../src/practise/channels/producerConsumer/main.go)
* chan<-  send only
* <-chan  read only
* chan    read/write
 
You needn’t close every channel when you’ve finished with it. It’s only necessary to close a channel when it is important to tell the receiving goroutines that all data have been sent. alex: GC会回收channel. 但是OS level的资源，如打开的文件，GC是无法自动close的。


## Buffered Channels
设置一个buffer size, (Alex, 所谓unbuffered Channels 的buffer size 应该是0???)

[code](https://github.com/yc-alex-xu/go/blob/master/src/practise/channels/buffered)
可以修改bufSZ


# Select
[example code](https://github.com/yc-alex-xu/go/tree/master/src/practise/select)
* main.go： original code
* default.go: 用default 改写
* timeout.go: 避免整个程序进入阻塞的情况
 

# race condition
Consider a function that works correctly in a sequential program. That function is **concurrency-safe** if it continues to work correctly even when called concurrently, that is, from two or more goroutines with no additional synchronization.

There are many reasons a function might not work when called concurrently, including  deadlock, livelock and resource starvation.
> In computer science, starvation is a multitasking-related problem, where a process is perpetually denied necessary resources. Without those resources, the program can never finish its task. Starvation is similar in effect to deadlock. Two or more programs become deadlocked together, when each of them wait for a resource occupied by another program in the same set. On the other hand, one or more programs are in starvation, when each of them is waiting for resources that are occupied by programs, that may or may not be in the same set that are starving. Moreover, in a deadlock, no program in the set changes its state. There may be a similar situation called livelock, where the processes changes their states continually, but cannot progress. Livelock is a special-case of starvation. In that sense, deadlock can also be said a special-case of starvation. Usually the problems in which programs are perpetually deprived of resources are referred as deadlock problems when none of them are changing their states and each of them is waiting for resources only occupied by programs in the same set. All other such problems are referred to as starvation.

A race condition is a situation in which the program does not give the correct result for some interleavings of the operations of multiple goroutines. 

A data race occurs whenever two goroutines access the same variable concurrently and at least one of the accesses is a write. It follows from this definition that there are three ways to avoid a data race.
* not to write the variable. 
* avoid accessing the variable from multiple goroutines
* mutual exclusion: 
 
# types of mutex:
* syn.mutex: By convention, the variables guarded by a mutex are declared immediately after the declaration of the mutex itself. The region of code between Lock and Unlock in which a goroutine is free  to read and modify the shared variables is called a critical section. Go’s mutexes are not re-entrant.
* sync.RWMutex: allows read-only operations to proceed in parallel with each other, but write operations to have fully exclusive access.
* sync.Once:a Once consists of a mutex and a boolean variable that records whether initialization has taken place;
  
# The Race Detector
Just add the **-race** flag to your ***go build, go run, or go test*** command. 

However, it can only detect race conditions that occur during a run.

# Goroutine Scheduling
goroutine use
* small and not fixed sized stack, so context switch is quicker
* multiplexes (or schedules) m goroutines on n (= GOMAXPROCS) OS threads.it is concerned **only** with the goroutines of a single Go program.
* goroutines schduling happen not due to timer
* Goroutines have no notion of identity to programmer.
* 

  
