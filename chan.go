package main

import (
	"fmt"
	"time"
)

func calc2(v1, v2 int, c chan int) {
	time.Sleep(2 * time.Second)
	c <- v1 * v2
	fmt.Println("calc2")
}
func calc(v1, v2 int, c chan int) {
	time.Sleep(4 * time.Second)
	c <- v1 * v2
	fmt.Println("calc1")
}

func main() {
	c := make(chan int)
	go calc(2, 3, c)
	go calc2(4, 5, c)

	a := <-c
	fmt.Println(a)
	b := <-c
	fmt.Println(b)
	fmt.Println(a * b)

}
