package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	digits := make(map[string]string)
	digits["one"] = "1"
	digits["two"] = "2"
	digits["three"] = "3"
	digits["four"] = "4"
	digits["five"] = "5"
	digits["six"] = "6"
	digits["seven"] = "7"
	digits["eight"] = "8"
	digits["nine"] = "9"
	sum := 0

	for scan.Scan() {
		calibration := scan.Text()
		leftDigit := ""
		rightDigit := ""

		for i := 0; i < len(calibration); i++ {
			maxWordLength := 5
			if leftDigit != "" {
				break
			}
			if isDigit(string(calibration[i])) {
				leftDigit = string(calibration[i])
				break
			}
			for i+maxWordLength >= len(calibration) {
				maxWordLength--
			}
			for j := i + maxWordLength; j > i; j-- {
				if val, ok := digits[calibration[i:j]]; ok {
					if leftDigit != "" {
						break
					}

					leftDigit = val
					break
				}
			}

		}

		for i := len(calibration) - 1; i >= 0; i-- {
			maxWordLength := 4
			if rightDigit != "" {
				break
			}
			if isDigit(string(calibration[i])) {
				rightDigit = string(calibration[i])
				break
			}
			for i-maxWordLength < 0 {
				maxWordLength--
			}

			for j := i - maxWordLength; j < i; j++ {

				if val, ok := digits[calibration[j:i+1]]; ok {
					if rightDigit != "" {
						break
					}
					rightDigit = val

					break
				}

			}
		}

		newNumString := leftDigit + rightDigit

		newNum, _ := strconv.Atoi(newNumString)

		sum += newNum
	}
	fmt.Print(sum)
}
