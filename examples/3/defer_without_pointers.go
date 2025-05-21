package main

import "fmt"

type Car struct {
	model string
}

func (c Car) PrintModel() {
	fmt.Println("Car model by PrintModel:", c.model)
}

func main() {
	c := Car{model: "DeLorean DMC-12"}
	defer c.PrintModel()
	c.model = "Chevrolet Impala"
	fmt.Println("Car model changed to:", c.model)
}
