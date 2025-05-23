package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "-c", "3", "google.com")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}
