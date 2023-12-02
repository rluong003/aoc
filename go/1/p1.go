package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	digits := make(map[string]string)
	digits["1"] = "1"
	digits["2"] = "2"
	digits["3"] = "3"
	digits["4"] = "4"
	digits["5"] = "5"
	digits["6"] = "6"
	digits["7"] = "7"
	digits["8"] = "8"
	digits["9"] = "9"
	sum := 0

	for scan.Scan() {
		calibration := scan.Text()
		l := 0
		r := len(calibration) - 1
		leftDigit := ""
		rightDigit := ""

		for leftDigit == "" || rightDigit == "" {
			leftChar := string(calibration[l])
			rightChar := string(calibration[r])

			if leftDigit == "" {
				val, ok := digits[leftChar]
				if ok {
					leftDigit = val
				}
				l += 1
			}

			if rightDigit == "" {
				val, ok := digits[rightChar]
				if ok {
					rightDigit = val
				}

				r -= 1
			}
		}

		newNumString := leftDigit + rightDigit

		newNum, _ := strconv.Atoi(newNumString)

		sum += newNum
	}
	fmt.Print(sum)
}
