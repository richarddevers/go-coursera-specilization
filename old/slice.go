package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main_dfzef() {
	var user_input string
	exit_char := "X"
	sli := make([]int, 3)

	for {
		fmt.Println("Enter an integer or enter 'X' to exit")
		fmt.Scan(&user_input)

		if user_input == exit_char {
			fmt.Println("'X' have been entered, bye.")
			return
		}

		val, err := strconv.Atoi(string(user_input))
		if err != nil {
			fmt.Println("Invalid value")
			continue
		}
		sli = append(sli, val)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
