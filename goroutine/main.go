package main

import (
	"fmt"
)

func main() {

	//定义函数
	say := func() {
		fmt.Println("hello world !!!")
	}
	//启动goroutine
	go say()
	//go sayHello()

	//main goroutine 等待goroutine执行
	//time.Sleep(1 * time.Second)
	fmt.Println("main goroutine execute.....")
}

func sayHello()  {
	fmt.Println("hello world !!!")
}
