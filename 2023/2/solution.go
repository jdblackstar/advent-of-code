package main

import (
	"bufio"
	"fmt"
	helpers "github.com/jdblackstar/advent-of-code"
	"log"
	"strconv"
	"strings"
)

type GameMaximum struct {
	ID           int
	GameMaxRed   int
	GameMaxGreen int
	GameMaxBlue  int
}

func parseGameMaximum(gameStr string) GameMaximum {
	game := GameMaximum{}
	// split the string into game number and rounds on colon
	parts := strings.Split(gameStr, ": ")
	game.ID, _ = strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))
	// split remaining string on semi-colons, these represent rounds
	rounds := strings.Split(parts[1], "; ")
	maxRed, maxGreen, maxBlue := 0, 0, 0

	// iterate over the rounds
	for _, round := range rounds {
		// split the rounds into color pairs
		// [2 blue], [3 red], [3 green]
		// strings.Split() removes the characters that are passed in
		pairs := strings.Split(round, ", ")

		// iterate over the pairs
		for _, pair := range pairs {
			// split the pair into color and count
			colorCount := strings.Split(pair, " ")
			count, _ := strconv.Atoi(colorCount[0])
			color := colorCount[1]

			// update the max count for the color
			switch color {
			case "red":
				if count > maxRed {
					maxRed = count
				}
			case "green":
				if count > maxGreen {
					maxGreen = count
				}
			case "blue":
				if count > maxBlue {
					maxBlue = count
				}
			}
		}
	}

	game.GameMaxRed = maxRed
	game.GameMaxGreen = maxGreen
	game.GameMaxBlue = maxBlue

	return game
}

func processGames(games []GameMaximum) int {
	var sumOfValidGameIDs int
	for _, game := range games {
		if game.GameMaxRed <= 12 && game.GameMaxGreen <= 13 && game.GameMaxBlue <= 14 {
			sumOfValidGameIDs += game.ID
		}
	}
	return sumOfValidGameIDs
}

func part_1() {
	file := helpers.OpenFile("2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var games []GameMaximum
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGameMaximum(line)
		games = append(games, game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(processGames(games))
}

type Round struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	ID     int
	Rounds []Round
}

func processMinimumCubes(games []Game) int {
	var totalPower int
	for _, game := range games {
		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, round := range game.Rounds {
			if round.Red > maxRed {
				maxRed = round.Red
			}
			if round.Green > maxGreen {
				maxGreen = round.Green
			}
			if round.Blue > maxBlue {
				maxBlue = round.Blue
			}
		}
		power := maxRed * maxGreen * maxBlue
		totalPower += power
	}
	return totalPower
}

func parseGame(gameStr string) Game {
	game := Game{}
	// split the string into game number and rounds on colon
	parts := strings.Split(gameStr, ": ")
	game.ID, _ = strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))
	// split remaining string on semi-colons, these represent rounds
	rounds := strings.Split(parts[1], "; ")

	for _, round := range rounds {
		// split the rounds into color pairs
		// [2 blue], [3 red], [3 green]
		// strings.Split() removes the characters that are passed in
		pairs := strings.Split(round, ", ")
		round := Round{}
		for _, pair := range pairs {
			// split the pair into color and count
			colorCount := strings.Split(pair, " ")
			count, _ := strconv.Atoi(colorCount[0])
			color := colorCount[1]

			// update the count for the color in the round
			switch color {
			case "red":
				round.Red = count
			case "green":
				round.Green = count
			case "blue":
				round.Blue = count
			}
		}
		game.Rounds = append(game.Rounds, round)
	}

	return game
}

func part_2() {
	file := helpers.OpenFile("2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var games []Game
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		games = append(games, game)
	}

	fmt.Println(processMinimumCubes(games))
}

func main() {
	part_1()
	part_2()
}
