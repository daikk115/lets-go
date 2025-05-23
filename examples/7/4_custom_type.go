package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func c2F(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) // explicit conversion
}

func main() {
	c := Celsius(36.5)
	f := c2F(c)

	fmt.Println("Celsius:", c)
	fmt.Println("Fahrenheit:", f)
}
