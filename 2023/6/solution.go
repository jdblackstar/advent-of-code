package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func parseInput(filepath string) ([]Race, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var times []int
	var distances []int
	var races []Race

	// Read the first line for times
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for i := 1; i < len(parts); i++ {
			time, _ := strconv.Atoi(parts[i])
			times = append(times, time)
		}
	}

	// Read the second line for distances
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for i := 1; i < len(parts); i++ {
			distance, _ := strconv.Atoi(parts[i])
			distances = append(distances, distance)
		}
	}

	// Create races from times and distances
	for i := 0; i < len(times); i++ {
		races = append(races, Race{Time: times[i], Distance: distances[i]})
		fmt.Printf("Parsed race with time %d and distance %d\n", times[i], distances[i]) // Debug print
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return races, nil
}

func calculateWaysToWin(race Race) int {
	waysToWin := 0
	for i := 0; i <= race.Time; i++ {
		distance := i * (race.Time - i)
		if distance > race.Distance {
			waysToWin++
		}
	}
	fmt.Printf("Calculated %d ways to win for race with time %d and distance %d\n", waysToWin, race.Time, race.Distance) // Debug print
	return waysToWin
}

func partOne(races []Race) int {
	totalWays := 1
	for _, race := range races {
		ways := calculateWaysToWin(race)
		totalWays *= ways
	}
	return totalWays
}

func main() {
	races, err := parseInput("2023/6/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := partOne(races)
	fmt.Println(result)
}
