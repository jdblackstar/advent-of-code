package main

import (
	"bufio"
	"fmt"
	helpers "github.com/jdblackstar/advent-of-code"
	"log"
	"strconv"
)

func part_1() {
	file := helpers.OpenFile("1/input.txt")
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := -1
		lastDigit := -1
		for _, r := range line {
			if '0' <= r && r <= '9' {
				if firstDigit == -1 {
					firstDigit = int(r - '0')
				}
				lastDigit = int(r - '0')
			}
		}
		if firstDigit != -1 && lastDigit != -1 {
			sum += firstDigit*10 + lastDigit
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var wordToNumber = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func part_2() {
	file := helpers.OpenFile("1/input.txt")
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := "", ""
		found := false

		for i := 0; i < len(line); i++ {
			if '0' <= line[i] && line[i] <= '9' {
				if !found {
					firstDigit = string(line[i])
					found = true
				}
				lastDigit = string(line[i])
			}

			for j := i + 1; j <= len(line); j++ {
				substr := line[i:j]
				for word := range wordToNumber {
					if substr == word {
						if !found {
							firstDigit = wordToNumber[word]
							found = true
						}
						lastDigit = wordToNumber[word]
					}
				}
			}
		}

		if firstDigit != "" && lastDigit != "" {
			value, err := strconv.Atoi(firstDigit + lastDigit)
			if err != nil {
				log.Fatal(err)
			}
			sum += value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func main() {
	fmt.Print("Solution for part 1: ")
	part_1()
	fmt.Print("Solution for part 2: ")
	part_2()
}
