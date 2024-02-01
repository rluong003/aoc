package main

import (
	"bufio"
	"fmt"
	"os"
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
	for scanner.Scan() {
		line := scanner.Text()
		splitNumbers := strings.Split(line, "|")
		winningNumbers := make(map[string]bool)

		cardString := splitNumbers[0]

		cardStringSplit := strings.Split(cardString, ":")
		winningNums := strings.Fields(cardStringSplit[1])

		givenNums := strings.Fields(splitNumbers[1])

		curSum := 0
		for _, num := range winningNums {
			winningNumbers[num] = true
		}

		for _, num := range givenNums {
			if _, ok := winningNumbers[num]; ok {
				if curSum == 0 {
					curSum = 1
				} else {
					curSum *= 2
				}
			}
		}
		total += curSum
	}
	fmt.Println(total)
}
