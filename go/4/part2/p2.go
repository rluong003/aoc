package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	totalTickets := make(map[string]int)
	game := 1
	for scanner.Scan() {
		line := scanner.Text()

		//convert curGame to string
		curGame := strconv.Itoa(game)

		if _, ok := totalTickets[curGame]; ok {
			totalTickets[curGame]++
		} else {
			totalTickets[curGame] = 1
		}
		total += totalTickets[curGame]

		splitNumbers := strings.Split(line, "|")
		winningNumbers := make(map[string]bool)

		cardString := splitNumbers[0]

		cardStringSplit := strings.Split(cardString, ":")
		winningNums := strings.Fields(cardStringSplit[1])

		givenNums := strings.Fields(splitNumbers[1])

		numWins := 0
		for _, num := range winningNums {
			winningNumbers[num] = true
		}

		for _, num := range givenNums {
			if _, ok := winningNumbers[num]; ok {
				numWins++
			}
		}

		for tickets := 0; tickets < totalTickets[curGame]; tickets++ {
			for i := game + 1; i <= numWins+game; i++ {
				if _, ok := totalTickets[strconv.Itoa(i)]; ok {
					totalTickets[strconv.Itoa(i)]++
				} else {
					totalTickets[strconv.Itoa(i)] = 1
				}
			}
		}
		game++
	}
	fmt.Println(total)
}
