package main

import (
	"fmt"
	"sync"
)

func main() {

	 var wg  sync.WaitGroup
	//定义函数
	say := func() {
		defer wg.Done()
		fmt.Println("hello world !!!")
	}
	//启动goroutine

	wg.Add(1)
	go say()
	//go sayHello()

	//main goroutine 等待goroutine执行
	wg.Wait()
	//time.Sleep(1 * time.Second)
	fmt.Println("main goroutine execute.....")
}

