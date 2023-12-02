package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

		cubes := strings.TrimSpace(lineSplit[1])

		// Split each block group by , and ;
		cubeQuantities := strings.FieldsFunc(cubes, func(r rune) bool {
			return r == ',' || r == ';'
		})

		// The maximum colour constants
		redLimit, greenLimit, blueLimit := 12, 13, 14

		lineValid := true

		for _, cubeQuantity := range cubeQuantities {
			cubeQuantity = strings.TrimSpace(cubeQuantity)

			pair := strings.SplitN(cubeQuantity, " ", 2)

			quantity, err := strconv.Atoi(pair[0])
			if err != nil {
				fmt.Println("Error parsing quantity:", err)
				continue
			}

			colour := strings.TrimSpace(pair[1])

			switch colour {
			case "red":
				if quantity > redLimit {
					fmt.Println(colour, quantity)
					lineValid = false
				}
			case "green":
				if quantity > greenLimit {
					fmt.Println(colour, quantity)
					lineValid = false
				}
			case "blue":
				if quantity > blueLimit {
					fmt.Println(colour, quantity)
					lineValid = false
				}
			}

			if !lineValid {
				break
			}
		}

		if lineValid {
			game := lineSplit[0]
			if len(game) > 0 {
				// Get only the number of the game by left trimming non digits
				gameNumberStr := strings.TrimLeftFunc(game, func(r rune) bool {
					return !unicode.IsDigit(r)
				})

				gameNumber, err := strconv.Atoi(gameNumberStr)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					total += gameNumber
				}
			}
		}
	}

	fmt.Println(total)
}
