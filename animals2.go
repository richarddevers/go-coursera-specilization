package main

import (
	"fmt"
)

type Animal interface {
	Speak()
	Move()
	Eat()
}

type Cow struct{ name string }

func (c Cow) Speak() {
	fmt.Println("moo")
}
func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}

type Bird struct{ name string }

func (b Bird) Speak() {
	fmt.Println("peep")
}
func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}

type Snake struct{ name string }

func (s Snake) Speak() {
	fmt.Println("hsss")
}
func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}

type CustomError struct {
	description string
}

func (c CustomError) Error() string {
	return c.description
}

func checkIfPossible(s string, arr [3]string) error {
	for _, val := range arr {
		if val == s {
			return nil
		}
	}
	custom_error := CustomError{description: fmt.Sprintf("Value %s not available. Possibility are: %s", s, arr)}
	return custom_error
}

func ScanUserInput() (string, string, string, error) {
	COMMANDS_POSSIBILITY := [3]string{"newanimal", "query"}
	ANIMALS_POSSIBILITY := [3]string{"cow", "bird", "snake"}
	ACTIONS_POSSIBILITY := [3]string{"speak", "eat", "move"}
	var first_arg, second_arg, third_arg string

	fmt.Print(">")
	_, err := fmt.Scan(&first_arg, &second_arg, &third_arg)
	if err != nil {
		fmt.Println("Error scanning user input ", err)
		return "", "", "", err
	}

	err_command := checkIfPossible(first_arg, COMMANDS_POSSIBILITY)
	if err_command != nil {
		return "", "", "", err_command
	}

	if first_arg == "newanimal" {
		err_animal := checkIfPossible(third_arg, ANIMALS_POSSIBILITY)
		if err_animal != nil {
			return "", "", "", err_animal
		}
	}

	if first_arg == "query" {
		err_action := checkIfPossible(third_arg, ACTIONS_POSSIBILITY)
		if err_action != nil {
			return "", "", "", err_action
		}
	}

	return first_arg, second_arg, third_arg, nil
}

func CreateAnimal(animal_name, animal_type string, animals_map map[string]Animal) {
	fmt.Println("animaltype ", animal_type)
	switch animal_type {
	case "cow":
		animals_map[animal_name] = Cow{name: animal_name}
	case "bird":
		animals_map[animal_name] = Bird{name: animal_name}
	case "snake":
		animals_map[animal_name] = Snake{name: animal_name}
	}

	fmt.Println("Created it!")

}

func QueryAnimal(name, action string, animals_map map[string]Animal) error {
	animal, ok := animals_map[name]
	fmt.Println("ok ", ok)
	if !ok {
		custom_error := CustomError{description: fmt.Sprintf("animal %s do not exist", name)}
		return custom_error
	}

	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
	return nil

}

func main() {
	animals_map := make(map[string]Animal)

	for {
		first_arg, second_arg, third_arg, err := ScanUserInput()

		if err != nil {
			fmt.Println(err)
			continue
		}

		switch first_arg {
		case "newanimal":
			CreateAnimal(second_arg, third_arg, animals_map)
		case "query":
			err_query := QueryAnimal(second_arg, third_arg, animals_map)
			if err_query != nil {
				fmt.Println(err_query)
				continue
			}
		}

	}

}
