// main.go
package main

import (
	"fmt"
	simplemath "mymath/mathutil"
)

func main() {
	fmt.Println("- Public name:", simplemath.Globalname)
	fmt.Println(">>> Plus:", simplemath.PlusTwoNumber(3, 5))
	fmt.Println(">>> Minus:", simplemath.MinusTwoNumber(3, 5))
}
