package main

import (
	"fmt"
	"math"
)

// Function that generates the displacement calculation function
func GenDisplaceFn(acc, initialVel, initialDisp float64) func(time float64) float64 {

	// The displacement calculation function
	timeCalcFn := func(time float64) float64 {
		return (0.5 * acc * math.Pow(time, 2)) + (initialVel * time) + initialDisp
	}
	return timeCalcFn
}

func main() {
	var acc, initialVel, initialDisp, time float64
	fmt.Println("Enter a value for Accleration: ")
	fmt.Scan(&acc)
	fmt.Println("Enter a value for initial Velocity: ")
	fmt.Scan(&initialVel)
	fmt.Println("Enter a value for initial Displacement: ")
	fmt.Scan(&initialDisp)

	displaceFunc := GenDisplaceFn(acc, initialVel, initialDisp)

	fmt.Println("Enter a value for time: ")
	fmt.Scan(&time)

	fmt.Println("The final displacement will be:", displaceFunc(time))
}
