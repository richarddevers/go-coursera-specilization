package main

import "fmt"

func main() {
	c := make(chan string)
	var data string
	go func() {
		for {
			fmt.Scan(&data)
			if data == "X" {
				close(c)
				break
			}
			c <- data
		}
	}()
	for i := range c {
		fmt.Println("hi ", i)
	}
}
