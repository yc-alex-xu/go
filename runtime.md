# gc
runtime，支持垃圾回收，这属于动态语言的特性之一吧，虽然目前来说GC不算完美，但是足以应付我们所能遇到的大多数情况，特别是Go1.1之后的GC。

# goroutine 相关函数

* Goexit 退出当前执行的goroutine，但是defer函数还会继续调用
* Gosched 让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
* NumCPU 返回 CPU 核数量
* NumGoroutine 返回正在执行和排队的任务总数
* GOMAXPROCS 用来设置可以并行计算的CPU核数的最大值，并返回之前的值。