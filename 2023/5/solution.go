package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func getRangeOfSeeds(filepath string) ([][]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip the first line
	seedsLine := scanner.Text()

	seedsStr := strings.Split(seedsLine, " ")[1:] // skip the "seeds:" part
	ranges := make([][]int, 0)
	totalCount := 0
	for i := 0; i < len(seedsStr); i += 2 {
		start, _ := strconv.Atoi(seedsStr[i])
		length, _ := strconv.Atoi(seedsStr[i+1])
		ranges = append(ranges, []int{start, start + length - 1})
		totalCount += length
	}

	fmt.Printf("Total count of seeds to be examined: %d\n", totalCount)
	return ranges, nil
}

func parseMapping(fileContent string) []map[string][][3]int {
	maps := make([]map[string][][3]int, 0)
	lines := strings.Split(fileContent, "\n")[2:] // start from line 3
	mapName := ""
	currentMap := make(map[string][][3]int)
	for _, line := range lines {
		if line == "" {
			maps = append(maps, currentMap)
			currentMap = make(map[string][][3]int)
			continue
		}
		if strings.Contains(line, " map:") {
			mapName = strings.Split(line, " map:")[0]
		} else {
			numsStr := strings.Split(line, " ")
			destination, _ := strconv.Atoi(numsStr[0])
			source, _ := strconv.Atoi(numsStr[1])
			length, _ := strconv.Atoi(numsStr[2])
			currentMap[mapName] = append(currentMap[mapName], [3]int{source, destination, length})
		}
	}
	maps = append(maps, currentMap)
	return maps
}

/*
example is to place fertilizer (f) 10 into a corresponding water
this is how we avoid the noob trap of creating every possible combination

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

If 53 <= f <= (53 + 8), then w = f - (53 - 49)
If 11 <= f <= (11 + 42), then w = f - (11 - 0)
If 0 <= f <= (0 + 7), then w = f - (0 - 42)
If 7 <= f <= (7 + 4), then w = f - (7 - 57)

The pattern is essentially:
if map[1] <= input <= (map[1] + map[2]), then output = input - (map[1] - map[0])

In this scenario, the last line executes properly, so if you enter 10, the output would be 10 - (7 - 57) or 60
*/

func findLowestSeed(filepath string, getSeedsFunc interface{}) int {
	file, _ := os.Open(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	fileContent := strings.Join(txtlines, "\n")

	maps := parseMapping(fileContent)

	mappingOrder := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	minSeed := int(^uint(0) >> 1) // set to max int value

	switch getSeeds := getSeedsFunc.(type) {
	case func(string) ([]int, error):
		seeds, _ := getSeeds(filepath)
		for _, seed := range seeds {
			fmt.Printf("Original seed: %d\n", seed)
			seed = mapSeed(seed, maps, mappingOrder)
			fmt.Printf("Final value: %d\n", seed)
			if seed < minSeed {
				minSeed = seed
			}
		}
	case func(string) ([][]int, error):
		ranges, _ := getSeeds(filepath)
		for _, rng := range ranges {
			for seed := rng[0]; seed <= rng[1]; seed++ {
				fmt.Printf("Original seed: %d\n", seed)
				seed = mapSeed(seed, maps, mappingOrder)
				fmt.Printf("Final value: %d\n", seed)
				if seed < minSeed {
					minSeed = seed
				}
			}
		}
	}

	return minSeed
}

func mapSeed(seed int, maps []map[string][][3]int, mappingOrder []string) int {
	for _, mapName := range mappingOrder {
		matched := false
		for _, currentMap := range maps {
			for _, mapping := range currentMap[mapName] {
				if mapping[0] <= seed && seed < mapping[0]+mapping[2] {
					seed = seed - (mapping[0] - mapping[1])
					fmt.Printf("Mapping through %s, seed becomes %d\n", mapName, seed)
					matched = true
				}
			}
		}
		if !matched {
			fmt.Printf("No match in %s map, seed stays %d\n", mapName, seed)
		}
	}
	return seed
}

func main() {
	filepath := "5/small.txt"
	fmt.Println("Solution for part one: ", findLowestSeed(filepath, getSeeds))
	fmt.Println("Solution for part two: ", findLowestSeed(filepath, getRangeOfSeeds))
}
