package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := takeInput()

	userOption := input

	options := []string{"rock", "paper", "scissors"}

	fmt.Printf("Sorry, but the computer chose %s\n", getWinningOption(userOption, options))

}

func takeInput() string {
	scanner := bufio.NewScanner(os.Stdin)

	// fmt.Println()
	// fmt.Printf("Enter a request: ")
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

func getWinningOption(userOption string, options []string) string {
	length := len(options)
	half := length / 2
	inputIndex := indexOf(userOption, options)
	if half+inputIndex >= length {
		return options[inputIndex+half-length]
	}
	return options[inputIndex+half]
}
