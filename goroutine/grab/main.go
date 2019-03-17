package main

import (
	"fmt"
)

func main() {
	var a [4]int64
	for i := 0; i < 4; i++ {
		go func(i int) {
			for {
				a[i]++
			}
		}(i)
	}
	//time.Sleep(time.Second)
	fmt.Println("main goroutine exit")
}
