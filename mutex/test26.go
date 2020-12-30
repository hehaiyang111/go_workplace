package main

import (
	"fmt"
	"sync"
)

/**
临界区：当程序并发的运行时，多个Go协程不应该同时访问修改那些共享资源的代码。这些修改共享资源的代码称为临界区。
举例：x = x + 1
case1： 假设 x 的初始值为 0。而协程 1 获取 x 的初始值，并计算 x + 1。而在协程 1 将计算值赋值给 x 之前，系统上下文切换到了协程 2。于是，协程 2 获取了 x 的初始值（依然为 0），并计算 x + 1。接着系统上下文又切换回了协程 1。现在，协程 1 将计算值 1 赋值给 x，因此 x 等于 1。然后，协程 2 继续开始执行，把计算值（依然是 1）复制给了 x，因此在所有协程执行完毕之后，x 都等于 1。
case2:  在上面的情形里，协程 1 开始执行，完成了三个步骤后结束，因此 x 的值等于 1。接着，开始执行协程 2。目前 x 的值等于 1。而当协程 2 执行完毕时，x 的值等于 2。
综上，从这两个例子可以发现，根据上下文切换的不同情形，x的最终值是1或者2.这种不太理想的情况成为竞态条件(Race Condition).其程序的输出是由协程的执行顺序决定的。
方法:在这一系列的情况下，如果在任意时刻只允许一个Go协程访问临界区就欸避免竞态条件。使用mutex可以达到这个目的
mutex:提供一种加锁机制，可确保在某时刻只有一个协程在临界区运行，以防止出现竞态条件。
mutex：在sync包可以找到。mutex定义了两个方法：lock和unlock。所有在lock和unlock之间的代码。都只能由一个Go协程执行，于是就可以避免竞态条件。
*/

/**
Example：含有竞态条件的程序
result:多运行几次会有不通的结果
*/
//var x = 0
//func increment(wg *sync.WaitGroup){
//	x = x + 1
//	wg.Done()
//}
//func main() {
//	var w sync.WaitGroup
//	for i := 0; i < 1000; i++ {
//		w.Add(1)
//		go increment(&w)
//	}
//	w.Wait()
//	fmt.Println("final value of x", x)
//}
/**
Example:使用mutex解决
result:结果固定为1000
*/
//var x = 0
//func increment(wg *sync.WaitGroup,m *sync.Mutex){
//	m.Lock()
//	x = x + 1
//	m.Unlock()
//	wg.Done()
//}
//func main() {
//	var w sync.WaitGroup
//	var m sync.Mutex
//	for i := 0; i < 1000; i++ {
//		w.Add(1)
//		go increment(&w,&m)
//	}
//	w.Wait()
//	fmt.Println("final value of x", x)
//}

/**
Example：除了使用mutex，我们还可以使用信道来处理竞态条件
处理方法：创建一个缓冲信道容量为1，在可能发生竞态条件的语句之前传入值进入缓冲信道，由于容量为1，所以任何其他协程试图写入该信道时，都会发生阻塞，直到信道的值被读取
*/
var x = 0

func increase(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	a := make(chan bool, 1)
	for i := 0; i < 2000; i++ {
		w.Add(1)
		go increase(&w, a)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
