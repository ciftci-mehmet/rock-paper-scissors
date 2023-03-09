package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	ratingFile       = "rating.txt"
	pointsForWinning = 100
	pointsForDrawing = 50
	pointsForLosing  = 0
)

func main() {
	scoreboard, err := readScoreboard()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter your name:")
	name := takeInput()
	fmt.Println("Hello,", name)
	rating := scoreboard[name]

	for {
		input := takeInput()

		if input == "!exit" {
			fmt.Println("Bye!")
			return
		}

		if input == "!rating" {
			fmt.Println("Your rating:", rating)
			continue
		}

		options := []string{"rock", "paper", "scissors"}

		userOption := input
		if indexOf(userOption, options) == -1 {
			fmt.Println("Invalid input")
			continue
		}

		computerOption := getRandomOption(options)
		result, message := getResult(userOption, computerOption, options)
		fmt.Println(message)
		switch result {
		case "win":
			scoreboard[name] += pointsForWinning
		case "draw":
			scoreboard[name] += pointsForDrawing
		case "lose":
			scoreboard[name] += pointsForLosing
		}
		saveScoreboard(scoreboard)
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

func getResult(userOption, computerOption string, options []string) (string, string) {
	if userOption == computerOption {
		return "draw", "There is a draw (" + computerOption + ")"
	}
	length := len(options)
	half := length / 2
	userOptionIndex := indexOf(userOption, options)
	computerOptionIndex := indexOf(computerOption, options)
	if userOptionIndex+half < length {
		if userOptionIndex+half >= computerOptionIndex && userOptionIndex < computerOptionIndex {
			return "lose", "Sorry, but the computer chose " + computerOption
		}
		return "win", "Well done. The computer chose " + computerOption + " and failed"
	} else {
		if (userOptionIndex+half)%length >= computerOptionIndex || userOptionIndex < computerOptionIndex {
			return "lose", "Sorry, but the computer chose " + computerOption
		}
		return "win", "Well done. The computer chose " + computerOption + " and failed"
	}
}

func readScoreboard() (map[string]int, error) {
	data, err := os.ReadFile(ratingFile)
	if errors.Is(err, os.ErrNotExist) {
		err2 := createFile()
		if err2 != nil {
			return nil, err2
		}
	} else if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	scoreboard := make(map[string]int)
	for _, v := range lines {
		splitLine := strings.Split(v, " ")
		if len(splitLine) != 2 {
			break
		}
		name := splitLine[0]
		rating, err := strconv.Atoi(splitLine[1])
		if err != nil {
			return nil, err
		}
		scoreboard[name] = rating
	}
	return scoreboard, nil
}

func saveScoreboard(scoreboard map[string]int) error {
	var scoreboardString string
	for k, v := range scoreboard {
		scoreboardString += k + " " + strconv.Itoa(v) + "\n"
	}
	bs := []byte(scoreboardString)
	err := os.WriteFile(ratingFile, bs, 0777)
	return err
}

func createFile() error {
	bs := []byte("")
	err := os.WriteFile(ratingFile, bs, 0777)
	if err != nil {
		return err
	}
	return nil
}
