package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

/*
	将整个文件读取到内存
	ioutil.ReadFile ＋　绝对路径
*/
/*
func main() {
	// 使用绝对路径
	data, err := ioutil.ReadFile("D:\\go_workplace\\src\\file\\test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
*/

/*
	按字节读取 file.read()
*/
/*
func main() {
	filePath := "D:\\go_workplace\\src\\file\\test.txt"
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("has some error")
		return
	}
	defer f.Close()
	a := make([]byte,0)
	b := make([]byte, 3)
	// TODO 用缓冲区读取
	for {
		//从file中读取到byte中,按照自己读
		f1, err := f.Read(b)
		fmt.Println(f1)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
		}
		if f1 == 0 {
			break
		}
		a = append(a, b[:f1]...)
	}
	fmt.Println(string(a))
}
*/

/*
	bufio 按行读取
	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
*/
/*
func main() {
	filePath := "D:\\go_workplace\\src\\file\\test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("has some error")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		// io.EOF文字流的结尾
		// 考虑最后一行，如果是前面的按照\n读取，到最后一行了考虑EOF
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("file read over")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err", err)
			return
		}
		fmt.Print(line)
	}
}
*/

/*
	文件写入操作：os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
	func OpenFile(name string, flag int, perm FileMode) (*File, error) {}
	name:要打开的文件名     flag:打开文件的模式
	os.O_WRONLY 只写
	os.O_CREATE 创建文件
	os.O_RDONLY readOnly
	os.O_RDWR readwrite
	os.O_TRUNC 清空
	os.O_APPEND 追加
	perm:文件权限，一个八进制树，读 4 ， 写2 执行 1
*/
/*
	将字符串写入文件
*/
/*
func main() {
	// 如果txt空则写入，不空则覆盖
	f, err := os.Create("D:\\go_workplace\\src\\file\\test.txt")
	if err != nil {
		fmt.Println("has some error")
		return
	}
	f1, err := f.WriteString("hello i am firstly write something with go")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(f1, "bytes written successfully")
	err = f.Close()
	if err != nil {
		panic(err)
		return
	}
}
*/

/*
	将字节写入到文件，使用write的方法
*/
/*
func main() {
	file, err := os.Create("D:\\go_workplace\\src\\file\\test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	file1, err := file.Write(d2)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println(file1, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
*/

/*
	将字符串一行一行的写入文件
*/
/*
func main() {
	file, err := os.Create("D:\\go_workplace\\src\\file\\file1")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	b := []string{"welcome to China", "China is a good place that attracts you not to go home", "believe me"}

	for _, v := range b {
		fmt.Fprintln(file, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file written successfully")
}
*/

/*
	追加到文件：我们将Hehaiyang追加到file1这个文件中
	这个文件以追加和写的方式打开，这些标志将通过Open方法实现
*/
/*
func main() {
	file, err := os.OpenFile("D:\\go_workplace\\src\\file\\file1", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "hehaiyang"
	_, err = fmt.Fprintln(file, newLine)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}
*/

/*
	并发写入文件：多个goroutine同时写文件时，可能会遇到竞态条件。因此当发生同步写的时候需要一个channel作为一致写入的条件
	应用：创建100个goroutine。每个goroutine将产生一个随机数，届时将有100个随机数产生，这些随机数将被写入到文件里面。
	步骤：
		1. 创建一个channel用来读和写这个随机数
		2. 创建100个生产者goroutine。每个goroutine将产生随机数并将随机数写入到channel里
		3. 创建一个消费者goroutine用来从channel读取随机数并将他写入文件。这样的话我们就只有一个goroutine向文件中写数据，从而避免竞态条件，
		4. 一旦完成则关闭文件
*/
// create random number
func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n //写入
	wg.Done() //调用方法表示完成
}

func consume(data chan int, done chan bool) {
	file, err := os.Create("D:\\go_workplace\\src\\file\\concurrent")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	for d := range data {
		_, err := fmt.Fprintln(file, d)
		if err != nil {
			fmt.Println(err)
			file.Close()
			done <- false
			return
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}
func main() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
