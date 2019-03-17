package main

import "fmt"

func main() {

	/**
		1. channel 可以实现sync.WaitGroup的功能
		2. 非缓存channel
		3. channel读写
	    4. 关闭channel
	 */
	//c := say()
	//fmt.Printf("from chan of int, the is value = %d\n",<-c )
	c := say()
	i, flag := <-c
	fmt.Printf("from chan of int, the is i = %v flag = %v\n", i, flag)

}

func say() <-chan int {
	c := make(chan int)
	go func() {
		fmt.Println("hello world !!!")
		c <- 1
	}()
	return c
}

func say2() <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		fmt.Println("hello world !!!")
		c <- 1
	}()
	return c
}
