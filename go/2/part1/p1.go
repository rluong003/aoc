package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalGames := 0

	cubeMap := make(map[string]int)

	cubeMap["blue"] = 14
	cubeMap["green"] = 13
	cubeMap["red"] = 12

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")

		gameNumber := parts[0]
		gameNumberSplit := strings.Split(gameNumber, " ")

		gameId, err := strconv.Atoi(gameNumberSplit[1])
		impossible := false

		if err != nil {
			log.Fatal(err)
		}

		gameString := parts[1]
		totalGames += gameId

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

				if err != nil {
					log.Fatal(err)
				}

				if scoreInt > cubeMap[color] {
					impossible = true
					break
				}
			}
		}
		if impossible {
			totalGames -= gameId
		}
	}
	fmt.Println(totalGames)

}
