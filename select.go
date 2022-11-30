package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)
	abort := make(chan bool)

	go func() {
		var data string
		for {
			fmt.Println("Please enter one input")
			fmt.Scan(&data)
			if data == "X" {
				abort <- true
			} else {
				c1 <- data
			}
		}
	}()

	for {
		select {
		case a := <-c1:
			fmt.Println("c1 arrived: ", a)
		case <-abort:
			fmt.Println("bye")
			return
		}
	}
}
