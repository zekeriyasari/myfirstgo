package greetings

import (
	"fmt"
	"math"
)

// Returns a hello message embeding the name
func SayHello(name string) string {
	return fmt.Sprintf("Hello %v", name)
}

// Adds 1 to `x``
func AddOne(x int) int {
	return x + 1
}

// Point data in two dimensional space
type Point struct {
	X float64 // x-coordinate
	Y float64 // y-coordinate
}

// Add method to Point bound to Point type
func (p Point) Magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Returns the dimension of `p`.
func dimension(p Point) (d int) {
	d = 2
	return
}

// Adds the points `p1` and `p2`
func AddPoints(p1, p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

// Returns the distance between from `p1` and `p2`
func Distance(p1, p2 Point) (dist float64) {
	if d := dimension(p1); d != 2 {
		panic("Invalid dimension")
	}

	dist = math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
	return
}

// Returns the maximum element of `arr`
func FindMax(arr [5]int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// Lorenz dynamics
type Lorenz struct {
	sigma float64
	beta  float64
	rho   float64
}

// Constructor with default initializer
func (node *Lorenz) Init(args ...float64) {
	n := len(args)
	if n == 0 {
		node.sigma = 10.
		node.beta = 8 / 3
		node.rho = 38.
	} else if n == 1 {
		node.sigma = args[0]
		node.beta = 8 / 3
		node.rho = 38.
	} else if n == 2 {
		node.sigma = args[0]
		node.beta = args[1]
		node.rho = 38.
	} else if n == 3 {
		node.sigma = args[0]
		node.beta = args[1]
		node.rho = args[2]
	} else {
		panic("Incorrect initialization")
	}
}
