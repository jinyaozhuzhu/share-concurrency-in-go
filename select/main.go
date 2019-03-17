package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	close(c1)
	c2 := make(chan string)
	close(c2)

	c := make(chan bool)
	go func() {
		var i, j int
		for {
			select {
			case <-c1:
				i++
			case <-c2:
				j++
			case t := <-time.After(1 * time.Microsecond):
				fmt.Println(t)
				fmt.Printf("i = %v\n", i)
				fmt.Printf("j = %v\n", j)
				close(c)
				return
			}
		}
	}()
	fmt.Printf("c = %v\n", <-c)
}
