package runGoroutine

import (
	"fmt"
	"runGoroutine/oneWayChannel"
	"sync"
)

var waitA sync.WaitGroup

/**
 *	使用defer
 */
func sayHello(i int) {
	defer waitA.Done()
	fmt.Println("hello", i)
}

func runWait() {
	fmt.Println("runWait=========================")
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	for i := 0; i < 10; i++ {
		waitA.Add(1) // 启动一个goroutine就登记+1
		go sayHello(i)
	}
	fmt.Println("Hi")
	waitA.Wait()
	fmt.Println("runWait=========================", "\n")
}

func recv(c chan int) {
	go recv1(c)
}
func recv1(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func checkZeroChannel() {
	fmt.Println("checkZeroChannel=========================")
	ch := make(chan int)
	go recv(ch) // 创建一个 goroutine 从通道接收值
	ch <- 10
	fmt.Println("checkZeroChannel=========================发送成功", "\n")
}

func checkChannel() {
	oneWayChannel.RunOneGoRouting()
	ch := make(chan bool, 3)
	ch <- true
	ch <- false
	ch <- false
	close(ch)
	i := 0
	for i < 1 {
		fmt.Println("------")
		v, ok := <-ch
		if !ok {
			fmt.Println("通道只有在取完通道剩余值的情况下才会关闭关闭", v, ok)
			break
		}
		i = 1
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭=====", v, ok)
			break
		}
		fmt.Printf("v1:%#v ok1:%#v\n", v, ok)
	}
	for v := range ch {
		fmt.Println("通道range", v)
	}
}

func Run() {
	runWait()
	checkZeroChannel()
	checkChannel()
}
