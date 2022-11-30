package main

import "fmt"

type zgeg struct {
	pouet int
}

func (z zgeg) add() {
	z.pouet += 1
}

func main() {
	toto := zgeg{pouet: 0}

	toto.add()
	fmt.Println(toto)
	toto.add()
	fmt.Println(toto)
	toto.add()
	fmt.Println(toto)
}
