package main

import "fmt"

type Duck interface {
	// Every duck need to say quackquack
	sayQuackQuack()
}

// 2 duck in 2 countries
type EnglishDuck struct {
}
type VietnamDuck struct {
}

// Receiver with method
func (d EnglishDuck) sayQuackQuack() {
	fmt.Println("English-Quack say quack quack!")
}

func (d VietnamDuck) sayQuackQuack() {
	fmt.Println("Vietnam-Quack nói quạc quạc!")
}

type Dog struct {
}

func (d Dog) Bark() {
	fmt.Println("go go")
}

// This function accept any duck, not dog
func sayQuack(duck Duck) {
	duck.sayQuackQuack()
}

func main() {
	// Create 2 ducks
	en := EnglishDuck{}
	vn := VietnamDuck{}

	sayQuack(en)
	sayQuack(vn)

	// dog := Dog{}
	// sayQuack(dog) // compile error - cannot use dog (type Dog) as type Duck
}
