// Write a program to sort an array of integers.
// The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
// Each partition should be of approximately equal size.
// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers.
// Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
// When sorting is complete, the main goroutine should print the entire sorted list.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const CHUNK_SIZE = 4

type userInputs struct {
	array1 []int
	array2 []int
	array3 []int
	array4 []int
}

func (u userInputs) merged() []int {
	var result []int
	result = append(u.array1, u.array2...)
	result = append(result, u.array3...)
	result = append(result, u.array4...)
	return result
}

func splitArray(arr []int) [][]int {
	var result [][]int

	for i := 0; i < CHUNK_SIZE; i++ {

		min := (i * len(arr) / CHUNK_SIZE)
		max := ((i + 1) * len(arr)) / CHUNK_SIZE

		result = append(result, arr[min:max])

	}
	fmt.Println("Splitted inputs: ", result)
	return result
}

func scanUserInput() string {
	fmt.Println("Please enter a list of integers (eg: '43 -2 17 14 2 29 37 53 90 -20 200 199 15'):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Println("Error scanning user inputs")
		os.Exit(1)
	}

	return scanner.Text()
}

func convertStrSliceToIntSlice(string_slice []string) []int {
	var result []int
	for _, s := range string_slice {
		converted, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error converting input to int bye")
			os.Exit(1)
		}
		result = append(result, converted)
	}
	return result
}

func GetUserInput() userInputs {
	raw_input := scanUserInput()                                          // return a single string
	raw_input_slice_str := strings.Fields(raw_input)                      // return slice of str
	raw_input_slice_int := convertStrSliceToIntSlice(raw_input_slice_str) // return slice of int

	array_of_int := splitArray(raw_input_slice_int)
	return userInputs{
		array1: array_of_int[0],
		array2: array_of_int[1],
		array3: array_of_int[2],
		array4: array_of_int[3],
	}
}

func sorter(arr []int, c chan []int) {
	fmt.Println("array to sort: ", arr)
	sort.Ints(arr)
	c <- arr
}

func main() {
	user_input := GetUserInput()
	c := make(chan []int, CHUNK_SIZE)

	go sorter(user_input.array1, c)
	go sorter(user_input.array2, c)
	go sorter(user_input.array3, c)
	go sorter(user_input.array4, c)
	fmt.Println("sorted array: ", <-c)
	fmt.Println("sorted array: ", <-c)
	fmt.Println("sorted array: ", <-c)
	fmt.Println("sorted array: ", <-c)

	merged_slice := user_input.merged()
	go sorter(merged_slice, c)
	fmt.Println("result:", <-c)
}
