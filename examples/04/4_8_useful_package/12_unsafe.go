/*
Unsafe package is used to perform low-level programming in Go.
Normally, we dont need to use it.

| Function / Type             | Description                                                  |
| --------------------------- | ------------------------------------------------------------ |
| `unsafe.Pointer`            | A generic pointer type, can be converted to/from any pointer |
| `unsafe.Sizeof(x)`          | Returns the number of bytes used by `x`                      |
| `unsafe.Alignof(x)`         | Returns the required alignment of `x`                        |
| `unsafe.Offsetof(struct.f)` | Returns the offset (in bytes) of field `f` in struct         |
*/

package main

import (
	"fmt"
	"unsafe"
)

func pointer() {
	var x int = 42
	var p *int = &x
	fmt.Println("Value of x:", *p)  // 42
	fmt.Println("Address of x:", p) // 0xc00000a0b8 (example)
}

func readFromAddress() {
	addr := uintptr(0xc0000120e0) // raw address from pointer(), normally it is same between 2 times running
	ptr := (*int)(unsafe.Pointer(addr))
	fmt.Println(*ptr)
}

func main() {
	var x string = "Hello, World!"
	fmt.Println("Size of x:", unsafe.Sizeof(x))   // 8 (pointer) + 8 (length) = 16 bytes
	fmt.Println("Align of x:", unsafe.Alignof(x)) // 8	bytes
	fmt.Println("\n")
	pointer()
	readFromAddress()
}
