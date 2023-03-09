package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	for {
		input := takeInput()

		if input == "!exit" {
			fmt.Println("Bye!")
			return
		}

		options := []string{"rock", "paper", "scissors"}

		userOption := input
		if indexOf(userOption, options) == -1 {
			fmt.Println("Invalid input")
			continue
		}

		computerOption := getRandomOption(options)
		fmt.Println(getResult(userOption, computerOption, options))

	}
}

func takeInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func indexOf(needle string, haystack []string) int {
	for k, v := range haystack {
		if needle == v {
			return k
		}
	}
	return -1
}

func getRandomOption(options []string) string {
	randomIndex := rand.Intn(len(options))
	return options[randomIndex]
}

func getResult(userOption, computerOption string, options []string) string {
	if userOption == computerOption {
		return "There is a draw (" + computerOption + ")"
	}
	length := len(options)
	half := length / 2
	userOptionIndex := indexOf(userOption, options)
	computerOptionIndex := indexOf(computerOption, options)
	if userOptionIndex+half < length {
		if userOptionIndex+half >= computerOptionIndex && userOptionIndex < computerOptionIndex {
			return "Sorry, but the computer chose " + computerOption
		}
		return "Well done. The computer chose " + computerOption + " and failed"
	} else {
		if (userOptionIndex+half)%length >= computerOptionIndex || userOptionIndex < computerOptionIndex {
			return "Sorry, but the computer chose " + computerOption
		}
		return "Well done. The computer chose " + computerOption + " and failed"
	}
}
