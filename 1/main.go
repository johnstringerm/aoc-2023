package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// Open the file or use os.Stdin for standard input
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	total := 0

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()

		// Get the first number from the line
		firstNum, err := getFirstNumber(line)
		// Get the last number from the line
		lastNum, err := getLastNumber(line)

		combinedNum, err := strconv.Atoi(fmt.Sprintf("%d%d", firstNum, lastNum))
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			total = total + combinedNum
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(total)
}

func getFirstNumber(line string) (int, error) {
	// Iterate through each character in the line
	for _, char := range line {
		// Check if the character is a digit
		if unicode.IsDigit(char) {
			// Convert the digit to an integer
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, err
			}
			return num, nil
		}
	}

	// If no digit is found, return an error or a default value
	return 0, fmt.Errorf("no digit found in the line")
}

func getLastNumber(line string) (int, error) {
	// Iterate backward through each character in the line
	for i := len(line) - 1; i >= 0; i-- {
		char := rune(line[i])

		// Check if the character is a digit
		if unicode.IsDigit(char) {
			// Convert the digit to an integer
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, err
			}
			return num, nil
		}
	}

	// If no digit is found, return an error or a default value
	return 0, fmt.Errorf("no digit found in the line")
}
