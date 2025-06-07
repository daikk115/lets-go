package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// math.Sqrt() takes a float64, so we need to convert x and y to float64
	var f float64 = math.Sqrt(float64(x*x + y*y))

	// math.Sqrt() returns a float64, so we need to convert it back to an int
	var z uint = uint(f)
	fmt.Println("Squared of x^2 + y^2 =", z)

}
