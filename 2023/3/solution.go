package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func getPossibleSymbols(filepath string) []rune {
	file := openFile(filepath)
	scanner := bufio.NewScanner(file)
	var symbols []rune
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if !unicode.IsNumber(char) && char != '.' {
				symbols = append(symbols, char)
			}
		}
	}
	return symbols

}

func contains(runeSlice []rune, targetRune rune) bool {
	for _, currentRune := range runeSlice {
		if currentRune == targetRune {
			return true
		}
	}
	return false
}

func isLineAdjacentToSymbol(current string, previous string, next string, symbols []rune) []bool {
	length := len(current)
	result := make([]bool, length)

	for i := 0; i < length; i++ {
		if unicode.IsDigit(rune(current[i])) {
			// check previous line
			if previous != "" {
				if i > 0 && contains(symbols, rune(previous[i-1])) {
					result[i] = true
				}
				if contains(symbols, rune(previous[i])) {
					result[i] = true
					continue
				}
				if i < len(previous)-1 && contains(symbols, rune(previous[i+1])) {
					result[i] = true
					continue
				}
			}
			// check next line
			if next != "" {
				if i > 0 && contains(symbols, rune(next[i-1])) {
					result[i] = true
					continue
				}
				if contains(symbols, rune(next[i])) {
					result[i] = true
					continue
				}
				if i < len(next)-1 && contains(symbols, rune(next[i+1])) {
					result[i] = true
					continue
				}
			}
			// check current line
			if i > 0 && contains(symbols, rune(current[i-1])) {
				result[i] = true
				continue
			}
			if i < length-1 && contains(symbols, rune(current[i+1])) {
				result[i] = true
				continue
			}
		}
	}
	return result
}

func part_1() {
	file := openFile("2023/3/input.txt")
	defer file.Close()

	symbols := getPossibleSymbols("2023/3/input.txt")

	scanner := bufio.NewScanner(file)
	var previous, current, next string
	sum := 0
	num := 0
	if scanner.Scan() {
		current = scanner.Text()
	}
	if scanner.Scan() {
		next = scanner.Text()
	}

	for scanner.Scan() {
		previous, current, next = current, next, scanner.Text()
		for i, char := range current {
			if unicode.IsDigit(char) && isLineAdjacentToSymbol(current, previous, next, symbols)[i] {
				num = num*10 + int(char-'0')
			} else {
				sum += num
				num = 0
			}
		}
		sum += num
		num = 0
	}

	// process the last line
	previous, current = current, next
	for i, char := range current {
		if unicode.IsDigit(char) && isLineAdjacentToSymbol(current, previous, "", symbols)[i] {
			num = num*10 + int(char-'0')
		} else {
			sum += num
			num = 0
		}
	}
	sum += num

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func part_2() {

}

func main() {
	part_1()
	part_2()
}
