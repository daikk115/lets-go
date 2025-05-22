package main

import "fmt"

func gototestfunc() {
	i := 0

Continue:
	fmt.Println(i)
	i++
	if i > 10 {
		goto Stop
	} else {
		goto Continue
	}
Stop:
	// Can not put this before Continue, because it will immediately return
	fmt.Println("Reach the limit")
	return
}

func main() {
	gototestfunc()
}
