package main

import (
	"fmt"
	"sync"
)

/**
之前讨论的是无缓冲信道，无缓冲信道的发送和接受是阻塞的
缓冲信道：只在缓冲满的时候，才会阻塞向缓冲信道发送数据
*/

/**
Example1：缓冲区的定义
*/
//func main() {
//	//使用make定义一个capacity为2的信道
//	a := make(chan string, 2)
//	a <- "s"
//	a <- "d"
//	fmt.Println(<-a)  		//读取完以后缓冲区就空了
//	fmt.Println(<-a)
//
//
//}

/**
Example：理解缓冲区
*/
//func readNumber(a chan int){
//	for i:=0; i<5; i++ {
//		a <- i
//		fmt.Println("successfully wrote", i, "to number")
//	}
//	close(a)
//}
//func main() {
//	a := make(chan int, 2)
//	go readNumber(a)
//	time.Sleep(2 * time.Second)
//	for b:= range a {
//		fmt.Println("read value", b,"from a")
//		// 读一个释放1个， go下称就会赋值一下，执行一下33行
//		time.Sleep(2 * time.Second)
//	}
//}

/**
Example：死锁
*/
//func main() {
//	a := make(chan int, 2)
//	a <- 2
//	a <- 3
//	a <- 4 // 这个时候发生阻塞，需要其他协程进行释放。可惜没有，死锁了
//	fmt.Println(<-a)
//	fmt.Println(<-a)
//}

/**
Example:容量与长度
		容量:指信道可以存储值的数量
		长度：缓冲信道的长度是指信道中当前排队的元素的个数
*/
//func main()  {
//	a := make(chan string, 3)
//	a <- "Hehaiyang"
//	a <- "liuxinyi"
//	fmt.Println(cap(a))
//	fmt.Println(len(a))
//	fmt.Println(<-a)
//	fmt.Println(cap(a))
//	fmt.Println(len(a))
//}

/**
waitGroup：用于实现工作池（Worker Pools）
工作池：用于等待一批协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。
假设：我们有三个并发执行的Go协程(由Go主协程生成)，Go主协程要等待这3个协程执行结束后，才会种植。这就可以使用WaitGroup来实现。
*/

func process(i int, group *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	//time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	//使计数减去1
	group.Done()
}

func main() {
	a := 3
	var b sync.WaitGroup
	// 创建三个协程并发执行
	for i := 0; i < a; i++ {
		// 计数器累加。 到0的时候解除信道阻塞
		b.Add(1)
		go process(i, &b)
	}
	//三个分协程若没有执行完，主协程阻塞。直到计数为0时候释放。
	b.Wait()
	fmt.Println("finished")
}
