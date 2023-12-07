package main

import (
	"bufio"
	"fmt"
	helpers "github.com/jdblackstar/advent-of-code"
	"unicode"
)

/*
--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/

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
