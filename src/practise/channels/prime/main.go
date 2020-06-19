package main

import "fmt"

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

/*
* generate * goroutine 发出从 2 开始的每个整数，每个新的 goroutine 仅过滤特定的质数倍数 - 2、3、5、7 …，将第一个找到的质数发送给 * main *。
如果旋转它从顶部看，您会看到从 goroutine 发送到 main 的所有数字都是质数。

————————————————
原文作者：Summer
转自链接：https://learnku.com/go/t/39490
版权声明：著作权归作者所有。商业转载请联系作者获得授权，非商业转载请保留以上作者信息和原文链接。
*/
func main() {
	ch := make(chan int)
	go generate(ch) // Launch generate goroutine.
	const (
		numPrimes = 10
	)
	for i := 0; i < numPrimes; i++ {
		//每一级filter(包括第一级的generator)都只产生一个prime,也就是第一个
		prime := <-ch
		fmt.Println(prime)
		chNext := make(chan int)
		//启动一个filter goroutine,如果它收到的数字不能被prime整除则输出到下一级
		go filter(ch, chNext, prime)
		ch = chNext
	}
	close(ch)
}
