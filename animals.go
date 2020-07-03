package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	invalid_prompt := "Invalid Command!"

	for true {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')

		if len(command) < 4 {
			fmt.Println(invalid_prompt)
			continue
		}

		if strings.TrimSpace(command) == "exit" {
			break
		}

		args := strings.Fields(command)

		switch args[0] {
		case "cow":
			switch args[1] {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			default:
				fmt.Println(invalid_prompt)
			}
		case "bird":
			switch args[1] {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			default:
				fmt.Println(invalid_prompt)
			}
		case "snake":
			switch args[1] {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			default:
				fmt.Println(invalid_prompt)
			}
		default:
			fmt.Println(invalid_prompt)
		}

	}
}
