package main

import "fmt"

func maindazda2() {
	var user_input float32
	var truncated_input int

	fmt.Scan(&user_input)

	truncated_input = int(user_input)
	fmt.Println(truncated_input)
}
