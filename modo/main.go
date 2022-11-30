package main

import (
	"mymodule/point"
)

func main() {
	toto := point.CreatePoint(2, 3)

	toto.ShowPoint()
	toto.IncreasePoint()
	toto.ShowPoint()
}
