package main

import (
	"bufio"
	"fmt"
	helpers "github.com/jdblackstar/advent-of-code"
	"unicode"
)

type Coordinate struct {
	X int
	Y int
}

func identifyNumbers(filepath string) map[Coordinate]int {
	file := helpers.OpenFile(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := make(map[Coordinate]int)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		number := 0
		startX := -1
		for x, char := range line {
			if unicode.IsDigit(char) {
				if number == 0 {
					startX = x
				}
				number = number*10 + int(char-'0')
			} else {
				if number != 0 {
					numLength := x - startX
					for i := 0; i < numLength; i++ {
						coord := Coordinate{X: startX + i, Y: y}
						numbers[coord] = number
					}
					fmt.Printf("Identified number %d at (%d, %d) to (%d, %d)\n", number, startX, y, startX+numLength-1, y) // print the identified number
					number = 0
				}
			}
		}
		if number != 0 {
			numLength := len(line) - startX
			for i := 0; i < numLength; i++ {
				coord := Coordinate{X: startX + i, Y: y}
				numbers[coord] = number
			}
			fmt.Printf("Identified number %d at (%d, %d) to (%d, %d)\n", number, startX, y, startX+numLength-1, y) // print the identified number
		}
		y++
	}
	return numbers
}

func checkAroundNumber(coord Coordinate, symbols map[Coordinate]rune) bool {
	directions := []Coordinate{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for _, direction := range directions {
		adjacentCoord := Coordinate{X: coord.X + direction.X, Y: coord.Y + direction.Y}
		if _, ok := symbols[adjacentCoord]; ok {
			return true
		}
	}

	return false
}

func findNumberPairsNearGears(gears map[Coordinate]rune, numbers map[Coordinate]int) [][]int {
	numberPairs := [][]int{}

	directions := []Coordinate{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for gear := range gears {
		adjacentNumbers := make(map[int]bool)
		for _, direction := range directions {
			adjacentCoord := Coordinate{X: gear.X + direction.X, Y: gear.Y + direction.Y}
			if num, ok := numbers[adjacentCoord]; ok {
				adjacentNumbers[num] = true
			}
		}
		if len(adjacentNumbers) == 2 {
			pair := make([]int, 0, 2)
			for num := range adjacentNumbers {
				pair = append(pair, num)
			}
			numberPairs = append(numberPairs, pair)
			gearRatio := pair[0] * pair[1]
			fmt.Printf("Identified number pair %d and %d with gear ratio %d\n", pair[0], pair[1], gearRatio) // print the number pair and gear ratio
		}
	}

	return numberPairs
}

func getPossibleSymbols(filepath string) map[Coordinate]rune {
	file := helpers.OpenFile(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	symbols := make(map[Coordinate]rune)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if !unicode.IsNumber(char) && char != '.' {
				coord := Coordinate{X: x, Y: y}
				symbols[coord] = char
				fmt.Printf("Identified symbol %c at (%d, %d)\n", char, x, y)
			}
		}
		y++
	}
	return symbols
}

func findGears(filepath string) map[Coordinate]rune {
	file := helpers.OpenFile(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gears := make(map[Coordinate]rune)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if char == '*' {
				coord := Coordinate{X: x, Y: y}
				gears[coord] = char
				fmt.Printf("Identified gear at (%d, %d)\n", x, y) // print the coordinates of the gear
			}
		}
		y++
	}
	return gears
}

func part_1() {
	filepath := "3/input.txt"
	symbols := getPossibleSymbols(filepath)
	numbers := identifyNumbers(filepath)

	sum := 0
	addedNumbers := make(map[int]bool)
	for coord, num := range numbers {
		if _, alreadyAdded := addedNumbers[num]; !alreadyAdded && checkAroundNumber(coord, symbols) {
			sum += num
			addedNumbers[num] = true
		}
	}

	fmt.Println(sum)
}

func part_2() {
	filepath := "3/input.txt"
	gears := findGears(filepath)
	numbers := identifyNumbers(filepath)

	numberPairs := findNumberPairsNearGears(gears, numbers)

	sum := 0
	for _, pair := range numberPairs {
		gearRatio := pair[0] * pair[1]
		sum += gearRatio
	}

	fmt.Println(sum)
}

func main() {
	part_1()
	part_2()
}
