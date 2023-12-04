package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input text file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")
		numbers := strings.Split(strings.TrimSpace(lineSplit[1]), "|")

		winningNumbers := stringToNumbersArray(strings.TrimSpace(numbers[0]))
		gameNumbers := stringToNumbersArray(strings.TrimSpace(numbers[1]))

		lineTotal := 0

		for _, gameNumber := range gameNumbers {
			for _, winningNumber := range winningNumbers {
				if gameNumber == winningNumber {
					if lineTotal == 0 {
						lineTotal += 1
					} else {
						lineTotal = lineTotal * 2
					}
				}
			}
		}

		total += lineTotal
	}

	fmt.Println(total)
}

func stringToNumbersArray(s string) []int {
	var numbers []int
	fields := strings.Fields(s)
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}
