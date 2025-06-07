package main

func testOpenDefer(a int) {
	if a == 1 {
		defer println("Defer in if") // open-coded defer
		return
	}
	if a == 2 {
		defer println("Defer in if") // open-coded defer
		// If we uncomment these lines, all defer will be converted to stack-allocated defer because 8 defer x 2 return > 15
		// return
	}
	if a == 3 {
		defer println("Defer in if") // open-coded defer
	}
	if a == 4 {
		defer println("Defer in if") // open-coded defer
	}
	if a == 5 {
		defer println("Defer in if") // open-coded defer
	}
	if a == 6 {
		defer println("Defer in if") // open-coded defer
	}
	if a == 7 {
		defer println("Defer in if") // open-coded defer
	}
	if a == 8 {
		defer println("Defer in if") // open-coded defer
	}
	// If we uncomment these lines, all above defer will be converted to stack-allocated defer because more than 8
	// if a == 9 {
	// 	defer println("Defer in if")
	// }
}

func testDefer(a int) {
	if a == 4 {
		defer println("Defer in if") // stack-allocated defer
	}
	if a == 5 {
		defer println("Defer in if") // stack-allocated defer
	}

	for range a {
		defer println("Defer in for") // heap-allocated defer
	}
}

func main() {
	testOpenDefer(10)
	testDefer(10)
}

// runtime.deferprocStack --> mean stack-allocated defer
// runtime.deferproc --> mean heap-allocated defer

// ubuntu@vm ~/G/l/e/3 (master)> go build -gcflags="-S" ./defer_type.go 2>&1 | grep -E "opendefer|runtime.deferproc"
//         0x002a 00042 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:3)   FUNCDATA        $4, main.testOpenDefer.opendefer(SB)
//         0x0037 00055 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:36)  CALL    runtime.deferprocStack(SB)
//         0x0063 00099 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:39)  CALL    runtime.deferprocStack(SB)
//         0x00a5 00165 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:43)  CALL    runtime.deferproc(SB)
//         rel 56+4 t=R_CALL runtime.deferprocStack+0
//         rel 100+4 t=R_CALL runtime.deferprocStack+0
//         rel 166+4 t=R_CALL runtime.deferproc+0
// main.testOpenDefer.opendefer SRODATA dupok size=2

// Khi có thêm `if a == 9`
// ubuntu@vm ~/G/l/e/3 (master)> go build -gcflags="-S" ./defer_type.go 2>&1 | grep -E "opendefer|runtime.deferproc"
//         0x0042 00066 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:5)   CALL    runtime.deferprocStack(SB)
//         0x0074 00116 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:8)   CALL    runtime.deferprocStack(SB)
//         0x00ad 00173 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:11)  CALL    runtime.deferprocStack(SB)
//         0x00e0 00224 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:14)  CALL    runtime.deferprocStack(SB)
//         0x0112 00274 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:17)  CALL    runtime.deferprocStack(SB)
//         0x0144 00324 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:20)  CALL    runtime.deferprocStack(SB)
//         0x0173 00371 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:23)  CALL    runtime.deferprocStack(SB)
//         0x01a7 00423 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:26)  CALL    runtime.deferprocStack(SB)
//         0x01cf 00463 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:29)  CALL    runtime.deferprocStack(SB)
//         rel 67+4 t=R_CALL runtime.deferprocStack+0
//         rel 117+4 t=R_CALL runtime.deferprocStack+0
//         rel 174+4 t=R_CALL runtime.deferprocStack+0
//         rel 225+4 t=R_CALL runtime.deferprocStack+0
//         rel 275+4 t=R_CALL runtime.deferprocStack+0
//         rel 325+4 t=R_CALL runtime.deferprocStack+0
//         rel 372+4 t=R_CALL runtime.deferprocStack+0
//         rel 424+4 t=R_CALL runtime.deferprocStack+0
//         rel 464+4 t=R_CALL runtime.deferprocStack+0
//         0x0037 00055 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:35)  CALL    runtime.deferprocStack(SB)
//         0x0063 00099 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:38)  CALL    runtime.deferprocStack(SB)
//         0x00a5 00165 (/home/ubuntu/GITHUB/lets-go/examples/3/defer_type.go:42)  CALL    runtime.deferproc(SB)
//         rel 56+4 t=R_CALL runtime.deferprocStack+0
//         rel 100+4 t=R_CALL runtime.deferprocStack+0
//         rel 166+4 t=R_CALL runtime.deferproc+0
