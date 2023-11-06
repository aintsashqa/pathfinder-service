package types

import (
	"fmt"
	"math"
)

type Point [2]float64

func (p *Point) X() float64 { return p[0] }
func (p *Point) Y() float64 { return p[1] }

func (p *Point) Add(p2 Point) Point {
	return Point{p.X() + p2.X(), p.Y() + p2.Y()}
}

func (p *Point) Distance(p2 Point) float64 {
	var (
		x = math.Pow(p.X()-p2.X(), 2)
		y = math.Pow(p.Y()-p2.Y(), 2)
	)

	return math.Sqrt(x + y)
}

func (p *Point) Eq(p2 Point) bool {
	return p.X() == p2.X() && p.Y() == p2.Y()
}

// String implements fmt.Stringer
func (p *Point) String() string {
	return fmt.Sprintf("(X:%f; Y:%f)", p.X(), p.Y())
}
