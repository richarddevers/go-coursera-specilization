package main

import "fmt"

func fezf() {
	animalMap := map[string]animal{
		"cow":   animal{"grass", "walk", "moo"},
		"bird":  animal{"worms", "fly", "peep"},
		"snake": animal{"mice", "slither", "hsss"},
	}
	for {
		fmt.Print(">")
		var name, method string
		fmt.Scan(&name, &method)
		target, ok := animalMap[name]
		if !ok {
			panic("invalid name")
		}
		switch method {
		case "eat":
			target.Eat()
		case "move":
			target.Move()
		case "speak":
			target.Speak()
		default:
			panic("invalid method")
		}
	}
}

type animal struct {
	food       string
	locomotion string
	spoken     string
}

func (a animal) Eat() {
	fmt.Println(a.food)
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
}

func (a animal) Speak() {
	fmt.Println(a.spoken)
}
