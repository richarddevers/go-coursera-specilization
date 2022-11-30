package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func BubbleSort(sli []int) {
	// Takes a slice of integers as an argument and returns nothing
	// The BubbleSort() function should modify the slice so that the elements are in sorted order.
	sli_backup := make([]int, len(sli))
	for {
		copy(sli_backup, sli)
		for i := range sli {
			if i < len(sli)-1 {
				Swap(sli, i)
			}

		}
		equal := reflect.DeepEqual(sli_backup, sli)
		if equal {
			break
		}
	}

}

func Swap(sli []int, i int) {
	// Should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
	// Swap the contents of the slice in position i with the contents in position i+1.
	current_val := sli[i]
	next_val := sli[i+1]
	if current_val > next_val {
		sli[i] = next_val
		sli[i+1] = current_val
	}

}

func GetUserInput() []int {
	var user_input string
	var result []int
	for user_input != "X" {
		fmt.Println("Enter a integer or press X to finish the slice.")
		fmt.Scan(&user_input)

		user_input_as_int, err := strconv.Atoi(user_input)

		if err == nil {
			result = append(result, user_input_as_int)
		} else {
			fmt.Println("Input is not convertible to an int, please retry")
		}

		if len(result) == 10 {
			break
		}
	}
	return result
}

func maindzadazddazd() {
	sli := GetUserInput()

	fmt.Println("Initial slice:", sli)

	BubbleSort(sli)

	fmt.Println("Sorted slice:", sli)

}
