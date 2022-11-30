package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	fmt.Println("start")

	go func() {
		fmt.Println("running ", <-c1)
		c2 <- 1
	}()

	go func() {
		fmt.Println("loading")
		time.Sleep(2 * time.Second)
		c1 <- 1
	}()

	<-c2
	fmt.Println("stop")
}
