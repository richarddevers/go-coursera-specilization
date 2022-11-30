package main

import (
	"fmt"
	"os"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func checkIfPossible(s string, possible [3]string) {
	found := false
	for _, val := range possible {
		if s == val {
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("%v requested is not managed:", s)
		os.Exit(1)
	}

}

func maifzefezn() {
	possible_animals := [3]string{"cow", "bird", "snake"}
	possible_actions := [3]string{"eat", "move", "speak"}

	var user_animal string
	var user_action string
	var created_animal Animal

	fmt.Print(">")
	fmt.Scan(&user_animal, &user_action)
	checkIfPossible(user_animal, possible_animals)
	checkIfPossible(user_action, possible_actions)

	if user_animal == "cow" {
		created_animal = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	} else if user_animal == "bird" {
		created_animal = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	} else {
		created_animal = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	}

	if user_action == "eat" {
		created_animal.Eat()
	} else if user_action == "move" {
		created_animal.Move()
	} else {
		created_animal.Speak()
	}

}
