package main

import (
	"fmt"
	"os"
)

func osEnv() {
	fmt.Println("\n------------------ osEnv ------------------")
	os.Setenv("APP_NAME", "MyGoApp")

	envVars := os.Environ()
	fmt.Println("Environment Variables:", len(envVars))

	value := os.Getenv("APP_NAME")
	pwd := os.Getenv("PWD")

	fmt.Println("APP_NAME:", value)
	fmt.Println("PWD:", pwd)
}

func main() {
	osEnv()
}
