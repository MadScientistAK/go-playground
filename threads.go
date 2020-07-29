package main

import "fmt"

func main() {
	// gobroutine 1

	x := 1
	x = x + 1
	fmt.Println(x)

	// gobroutine 2

	y := &x
	*y = *y + 1
	fmt.Println(x)
}

/*  These routines, when run concurrently, will have a race condition if the order of execution changes.
Since y is changing the value held by the address of x, it can fail if x hasn't been declared before y tries to change its value.
Similarly, the value of x printed may differ according to the order in which the statements have been executed.
*/
