package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func maindzadezdez() {
	var filename string

	fmt.Println("Enter filename:")
	fmt.Scan(&filename)

	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file,bye")
		os.Exit(1)
	}

	// Create a slice of string, one for each line
	var file_lines []string
	var tmp []byte
	for _, v := range data {
		if string(v) != "\n" {
			tmp = append(tmp, v)
		} else {
			file_lines = append(file_lines, string(tmp))
			tmp = nil
		}
	}

	// Create a slice of person
	var person_slice []Person
	for _, v := range file_lines {
		split := strings.Split(v, " ")

		if len(split) != 2 {
			fmt.Println("Error creating Person")
			os.Exit(1)
		}

		person_slice = append(person_slice, Person{fname: split[0], lname: split[1]})

	}

	// Iter over the persone slice and print
	for _, p := range person_slice {
		fmt.Printf("%v %v \n", p.fname, p.lname)
	}
}
