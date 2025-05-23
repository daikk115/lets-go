package mathutil

import "fmt"

var Globalname = "Plus"
var name = "plus"

func PlusTwoNumber(a, b int) int {
	return privatePlusTwoNumber(a, b)
}

func privatePlusTwoNumber(a, b int) int {
	fmt.Println("- Private name", name)
	return a + b
}
