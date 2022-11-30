package point

import "fmt"

type Point struct {
	x int
	y int
}

func CreatePoint(x, y int) Point {
	return Point{x: x, y: y}
}

func (p Point) ShowPoint() {
	fmt.Println("x:", p.x)
	fmt.Println("y:", p.x)
}

func (p *Point) IncreasePoint() {
	p.x += 1
	p.y += 1
}
