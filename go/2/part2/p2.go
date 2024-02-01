package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalGames := 0

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")

		if err != nil {
			log.Fatal(err)
		}

		minBlue := 0
		minGreen := 0
		minRed := 0

		gameString := parts[1]

		gameParts := strings.Split(gameString, ";")
		for _, gamePart := range gameParts {
			gamePart = strings.TrimSpace(gamePart)
			gamePartSplit := strings.Split(gamePart, ",")

			for _, game := range gamePartSplit {
				curGame := strings.TrimSpace(game)

				splitCurGame := strings.Split(curGame, " ")

				color := splitCurGame[1]
				score := splitCurGame[0]

				scoreInt, err := strconv.Atoi(score)
				if color == "blue" {
					minBlue = max(minBlue, scoreInt)
				} else if color == "green" {
					minGreen = max(minGreen, scoreInt)
				} else if color == "red" {
					minRed = max(minRed, scoreInt)
				}

				if err != nil {
					log.Fatal(err)
				}

			}
		}

		totalGames += (minBlue * minGreen * minRed)

	}
	fmt.Println(totalGames)

}
