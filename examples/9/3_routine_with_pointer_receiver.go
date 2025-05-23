package main

import (
	"fmt"
	"time"
)

type Car struct {
	model string
}

func (c *Car) PrintModel() {
	time.Sleep(1 * time.Second)
	fmt.Println("Car model by PrintModel:", c.model)
}

func main() {
	c := Car{model: "DeLorean DMC-12"}
	go c.PrintModel()
	c.model = "Chevrolet Impala"
	fmt.Println("Car model changed to:", c.model)
	time.Sleep(3 * time.Second)
}
