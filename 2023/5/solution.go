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

func main() {
	filepath := "5/input.txt"
	seeds, _ := getSeeds(filepath)

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

	mappingOrder := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

	newSeeds := make([]int, len(seeds))
	minSeed := int(^uint(0) >> 1) // set to max int value
	for i, seed := range seeds {
		fmt.Printf("Original seed: %d\n", seed)
		for _, mapName := range mappingOrder {
			for _, currentMap := range maps {
				for _, mapping := range currentMap[mapName] {
					if mapping[0] <= seed && seed <= mapping[0]+mapping[2] {
						seed = seed - (mapping[0] - mapping[1])
						fmt.Printf("Mapping through %s, seed becomes %d\n", mapName, seed)
						break
					}
				}
			}
		}
		newSeeds[i] = seed
		fmt.Printf("Final value: %d\n", seed)
		if seed < minSeed {
			minSeed = seed
		}
	}
	fmt.Println("New seeds:", newSeeds)
	fmt.Println("Minimum seed:", minSeed)
}
