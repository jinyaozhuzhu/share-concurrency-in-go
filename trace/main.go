package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//runtime.GOMAXPROCS()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	for i := 0; i < 10; i++ {
	    go func(i int) {
	    	fmt.Printf("hello world %d \n",i)
		}(i)
	}
	time.Sleep(2*time.Second)
}
