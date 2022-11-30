package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool, 1)
	go func() {
		fmt.Println("new routine")
		c <- true
		fmt.Println("after first send")
		c <- false
		fmt.Println("after second send")
		c <- true
		fmt.Println("after third send")
	}()

	fmt.Println("main routine 1", <-c)
	time.Sleep(2 * time.Second)
	fmt.Println("main routine 2", <-c)
	time.Sleep(2 * time.Second)
	fmt.Println("main routine 3", <-c)
}
