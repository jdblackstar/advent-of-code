package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

/*
attributes:
{
	"seed",
	"soil",
	"fertilizer",
	"water",
	"light",
	"temperature",
	"humidity",
	"location",
}

example mapping:
destination, source, length

seed-to-soil map:
50 98 2
52 50 48

parse what this means:
98 -> 50
99 -> 51

50 -> 52
51 -> 53
...
96 -> 98
97 -> 99

in english:
the second number in the map (seed) maps to the first number in the map (soil)
do this [the third number in the map] times, always stepping by 1

in pseudo-code:
map = [50, 98, 2]
for step in map[2]:
	map[1] + (step - 1) maps_to map[0] + (step - 1)

98 -> 50
99 -> 51

then repeat for next map until we get to location
*/



func getSeeds(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seedsLine := scanner.Text()

	seedsStr := strings.Split(strings.TrimPrefix(seedsLine, "seeds: "), " ")
	seeds := make([]int, len(seedsStr))
	for i, seedStr := range seedsStr {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			return nil, err
		}
		seeds[i] = seed
	}

	return seeds, nil
}


func main() {
	filepath := "5/small.txt"
	seeds, _ := getSeeds(filepath)

	fmt.Println(seeds)
}