package main

import (
	"fmt"
	"os"
)

// s = ½ a t2 + vot + so

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*(t*t) + (vo * t) + so
	}
}

func GetOneUserInput() float64 {
	var user_input float64
	_, err := fmt.Scan(&user_input)
	if err != nil {
		fmt.Println("Error scanning user input")
		os.Exit(1)
	}

	return user_input
}

func GetInitialUserInputs() (a, vo, so float64) {
	fmt.Println("Enter acceleration:")
	_a := GetOneUserInput()
	fmt.Println("Enter initial velocity:")
	_vo := GetOneUserInput()
	fmt.Println("Enter initial displacement:")
	_so := GetOneUserInput()
	return _a, _vo, _so
}

func maifafzan() {
	a, vo, so := GetInitialUserInputs()
	fnGen := GenDisplaceFn(a, vo, so)

	fmt.Println("Enter time:")
	t := GetOneUserInput()

	fmt.Println("Displacement:")
	fmt.Println(fnGen(t))
}
