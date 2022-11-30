package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		fmt.Println("new routine")
		wg.Done()
	}(&wg)

	wg.Wait()
	fmt.Println("main routine")
}
