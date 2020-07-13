package main

import (
	"fmt"
)

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func makeCow() Cow {
	c := Cow{}
	c.food = "grass"
	c.locomotion = "walk"
	c.noise = "moo"
	return c
}

func main() {
	a := makeCow()
	fmt.Println(a.food)
}
