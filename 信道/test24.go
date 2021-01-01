package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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
waitGroup：用于等待一批协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。
假设：我们有三个并发执行的Go协程(由Go主协程生成)，Go主协程要等待这3个协程执行结束后，才会执行。这就可以使用WaitGroup来实现。
*/

//func process(i int, group *sync.WaitGroup) {
//	fmt.Println("started Goroutine ", i)
//	//time.Sleep(2 * time.Second)
//	fmt.Printf("Goroutine %d ended\n", i)
//	//使计数减去1
//	group.Done()
//}
//
//func main() {
//	a := 3
//	var b sync.WaitGroup
//	// 创建三个协程并发执行
//	for i := 0; i < a; i++ {
//		// 计数器累加。 到0的时候解除信道阻塞
//		b.Add(1)
//		// 需要传入地址
//		go process(i, &b)
//	}
//	//三个分协程若没有执行完，主协程阻塞。直到计数为0时候释放。
//	b.Wait()
//	fmt.Println("finished")
//}

/**
应用：缓冲信道的重要应用之一就是实现工作池。
工作池：一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。
实现：使用缓冲信道实现工作池。我们的工作池的任务是计算所输入数字的每一位的和。例如，如果输入234，结果会是9。像工作池输入的是一列伪随机数。
核心功能：
	1）；创建一个Go协程池，监听一个等待作业分配的输入型缓冲信道。
	2）：将作业添加到该输入型缓冲信道。
	3）：作业完成后，再将结果写入一个输出型缓冲信道。
	4）：从输出型缓冲信道读取并打印结果。
*/

// randnum:随机数
// computerTotalNum:数字之和
type Job struct {
	id      int
	randnum int
}
type Result struct {
	job              Job
	computerTotalNum int
}

// 创建用于接收结果和输出结果的缓冲信道
// 工作协程(Worker Goroutine)会监听缓冲信道Job里更新的作业.一旦工作协程完成了作业,其结果会写入缓冲信道results.
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

//计算每一位之和返回结果
func digits(a int) int {
	sum := 0
	for a != 0 {
		digit := a % 10
		sum += digit
		a /= 10
	}
	// 为了模仿计算过程花费了一段时间，则在函数中加了两秒休眠
	//time.Sleep(2 * time.Second)
	return sum
}

// 创建一个工作协程的函数
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randnum)}
		results <- output
	}
	wg.Done()
}

// 创建一个Go协程的工作池
func createWorkerPool(numOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		// 传地址
		// 创建协程
		go worker(&wg)
	}
	//等待所有的Go协程执行完毕.
	//等待过程堵塞
	wg.Wait()
	// 所有协程完成后，关闭result信道。
	// 因为所有协程都已经执行完毕，于是不再需要对result信道写入数据
	close(results)
}

// 把作业分配给工作者
func allocate(numOfJobs int) {
	for i := 0; i < numOfJobs; i++ {
		randnum := rand.Intn(999)
		job := Job{i, randnum}
		jobs <- job
	}
	close(jobs)
}

// 创建一个读取results信道和打印输出的函数
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randnum, result.computerTotalNum)
		time.Sleep(1 * time.Second)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	numOfJobs := 12
	go allocate(numOfJobs)
	done := make(chan bool)
	go result(done)
	numOfWorkers := 10
	createWorkerPool(numOfWorkers)
	// 如果不写会导致主线程提前结束。因为go result不会阻塞主线程
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
