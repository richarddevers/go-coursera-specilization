package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func scan_user_inputs() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text(), scanner.Err()
	}
	return "", nil
}

func main_dezdez() {

	var user_name string
	var user_address string
	user_data := make(map[string]string)

	fmt.Println("Enter your name")
	user_name, err_name := scan_user_inputs()
	if err_name != nil {
		fmt.Println("Error reading user input. Bye")
		return
	}
	user_data["name"] = user_name

	fmt.Println("Enter your address")
	user_address, err_address := scan_user_inputs()
	if err_address != nil {
		fmt.Println("Error reading user input. Bye")
		return
	}
	user_data["address"] = user_address

	user_data_json, err_json := json.Marshal(user_data)
	if err_json != nil {
		fmt.Println("Error reading marshalling user data. Bye")
		return
	}
	fmt.Printf("%s", user_data_json)
}
