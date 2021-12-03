package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"

	"example.com/greetings"
)

func compute(myfunc func(float64, float64) float64) float64 {
	return myfunc(3, 4)
}

// Wrapper function
func adder() func(int) int {
	sum := 0
	wrapped := func(i int) int {
		sum += i
		return sum
	}
	return wrapped
}

// Returns a randomly initialized Point
func randomPoint() greetings.Point {
	return greetings.Point{X: rand.Float64(), Y: rand.Float64()}
}

func main() {
	// RNG seed
	rand.Seed(time.Now().UnixNano())

	// Print a simple message
	fmt.Printf("This is a simple sentence")

	// Call a function from the package
	message := "John"
	fmt.Println(greetings.SayHello(message))

	// Define variables
	x := 5
	fmt.Println(x, " + ", 1, " = ", greetings.AddOne(x))

	// Define sum struct
	p := greetings.Point{X: 2, Y: 3}
	fmt.Printf("The point is %v of type %T", p, p)
	fmt.Println("The type is ", reflect.TypeOf(p))

	// Initialize a vector
	vec := [5]int{}
	for i := 0; i < 5; i++ {
		vec[i] = 10 * i
	}
	fmt.Printf("The vector: %v\n", vec)

	// Compute a sum
	total := 0.
	var num float64
	for i := 0.; i < 10; i++ {
		num = rand.Float64()
		fmt.Println("The num is: ", num)
		total += num
	}
	fmt.Printf("The total is: %v\n", total)

	// Find the largest
	fmt.Printf("The maximum is %v\n", greetings.FindMax(vec))

	// Construct two random points
	p1, p2 := greetings.Point{X: rand.Float64(), Y: rand.Float64()},
		greetings.Point{X: rand.Float64(), Y: rand.Float64()}
	distance := greetings.Distance(p1, p2)
	fmt.Println("The distance between ", p1, " and ", p2, " is: ", distance)

	// A slice example
	v := [10]int{}
	s := v[1:3]
	fmt.Println("The slice length", len(s), " The slice capacity: ", cap(s))
	fmt.Println("The first element of slice: ", s[0])
	s2 := []int{0, 1, 2, 3, 4}
	fmt.Println("The slice length", len(s2), " The slice capacity: ", cap(s2))
	s3 := make([]int, 5)
	for i := 0; i < 5; i++ {
		s3[i] = 10 * i
	}
	fmt.Println("The slice content is: ", s3)
	fmt.Println("The slice length", len(s3), " The slice capacity: ", cap(s3))
	s4 := make([]float64, 10)
	for i := 0; i < len((s4)); i++ {
		s4[i] = rand.Float64()
	}
	fmt.Println("The randomly initialized array ", s4)
	s5 := []struct {
		int
		bool
	}{
		{1, true},
		{2, false},
		{3, false},
		{4, true},
	}
	fmt.Println("The slice s5 content: ", s5)
	fmt.Println("The first content: ", s[1])
	s6 := [][]int{
		make([]int, 5),
		make([]int, 5),
		make([]int, 5),
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			s6[i][j] = i * j
		}
	}
	fmt.Println("The content of s6: ", s6)

	s7 := make([]float64, 10)
	for i := 0; i < 10; i++ {
		s7[i] = rand.Float64()
	}
	for k, j := range s7 {
		fmt.Printf("%v => %.2v\n", k, j)
	}

	// Maps,
	d := make(map[int]greetings.Point)
	d[1] = greetings.Point{X: rand.Float64(), Y: rand.Float64()}
	d[2] = greetings.Point{X: rand.Float64(), Y: rand.Float64()}
	d[3] = greetings.Point{X: rand.Float64(), Y: rand.Float64()}
	fmt.Println("The content d is ", d)
	d2 := make(map[string]float64)
	d2["first"] = 1.
	d2["second"] = 2.
	fmt.Println("The content of d2: ", d2)

	// Define a function
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Printf("The function value: %v. The function type: %T\n", hypot, hypot)
	fmt.Println("Computing the composition: ", compute(hypot))

	// Call composition function
	original := adder() // wrapped function
	for i := 0; i < 10; i++ {
		fmt.Println("The sum is ", original(i))
	}

	// Construct a points
	p3 := randomPoint()
	fmt.Printf("Point magnitude: %v\n", p3.Magnitude())

	// Initialize a node
	node := new(greetings.Lorenz)
	node.Init(10., 8/3, 28.)
	fmt.Println("The Lorenz node ", node)
}
