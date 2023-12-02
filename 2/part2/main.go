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
	file, err := os.Open("../input.txt")
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

		// Need to get the largest colour value for each colour
		redMax, greenMax, blueMax := 0, 0, 0

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
				if quantity > redMax {
					redMax = quantity
				}
			case "green":
				if quantity > greenMax {
					greenMax = quantity
				}
			case "blue":
				if quantity > blueMax {
					blueMax = quantity
				}
			}
		}

		power := redMax * greenMax * blueMax

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			total += power
		}
	}

	fmt.Println(total)
}
