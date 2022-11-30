package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)
	go func() {
		fmt.Println("running first time")
		time.Sleep(2 * time.Second)
		c <- true
	}()

	fmt.Println("received first data", <-c)
	fmt.Println("main routine")
}
