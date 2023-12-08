package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	helpers "github.com/jdblackstar/advent-of-code"
)

func splitNumbers(input string) (string, []int, []int) {
	splitInput := strings.Split(input, "|")
	cardIDAndWinningNumbersStr := strings.Fields(splitInput[0])
	ourNumbersStr := strings.Fields(splitInput[1])

	cardID := cardIDAndWinningNumbersStr[0]
	winningNumbersStr := cardIDAndWinningNumbersStr[1:]

	winningNumbers := make([]int, len(winningNumbersStr))
	ourNumbers := make([]int, len(ourNumbersStr))

	for i, numStr := range winningNumbersStr {
		num, _ := strconv.Atoi(numStr)
		winningNumbers[i] = num
	}

	for i, numStr := range ourNumbersStr {
		num, _ := strconv.Atoi(numStr)
		ourNumbers[i] = num
	}

	return cardID, winningNumbers, ourNumbers
}

func totalPointsPerCard(winningNumbers []int, ourNumbers []int) int {
	points := 0
	matches := 0

	for _, ourNumber := range ourNumbers {
		for _, winningNumber := range winningNumbers {
			if ourNumber == winningNumber {
				matches++
				break
			}
		}
	}

	if matches > 0 {
		points = 1 << (matches - 1)
	}

	return points
}

func part_1(filepath string) int {
	file := helpers.OpenFile(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0

	for scanner.Scan() {
		card := scanner.Text()
		// use the blank identifier to ignore the cardID returned from this function
		_, winningNumbers, ourNumbers := splitNumbers(card)
		totalPoints += totalPointsPerCard(winningNumbers, ourNumbers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalPoints
}

func main() {
	fmt.Println("Solution for part 1: ", part_1("4/input.txt"))
}
