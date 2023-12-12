package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	helpers "github.com/jdblackstar/advent-of-code"
)

type Card struct {
	ID             int
	WinningNumbers []int
	OurNumbers     []int
}

func splitCardAttributes(s string) (int, []int, []int) {
	// Split the string on the colon first to separate the card ID from the rest
	parts := strings.Split(s, ":")
	cardIDStr := strings.TrimSpace(parts[0])
	cardID, _ := strconv.Atoi(strings.Split(cardIDStr, " ")[1])

	// Split the remaining string on the pipe to separate the winning numbers from our numbers
	numbersParts := strings.Split(parts[1], "|")
	winningNumbersStr := strings.Fields(strings.TrimSpace(numbersParts[0]))
	ourNumbersStr := strings.Fields(strings.TrimSpace(numbersParts[1]))

	// Convert the number strings to integers
	winningNumbers := make([]int, len(winningNumbersStr))
	for i, numStr := range winningNumbersStr {
		winningNumbers[i], _ = strconv.Atoi(numStr)
	}

	ourNumbers := make([]int, len(ourNumbersStr))
	for i, numStr := range ourNumbersStr {
		ourNumbers[i], _ = strconv.Atoi(numStr)
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

func countMatches(winningNumbers []int, ourNumbers []int) int {
	matches := 0

	for _, ourNumber := range ourNumbers {
		for _, winningNumber := range winningNumbers {
			if ourNumber == winningNumber {
				matches++
				break
			}
		}
	}

	return matches
}

func part_1(filepath string) int {
	file := helpers.OpenFile(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0

	for scanner.Scan() {
		cardStr := scanner.Text()
		cardID, winningNumbers, ourNumbers := splitCardAttributes(cardStr)
		card := Card{
			ID:             cardID,
			WinningNumbers: winningNumbers,
			OurNumbers:     ourNumbers,
		}
		totalPoints += totalPointsPerCard(card.WinningNumbers, card.OurNumbers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalPoints
}

type Deck struct {
	Cards map[int]Card
	Count int
}

func NewDeck(filepath string) *Deck {
	file := helpers.OpenFile(filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	deck := &Deck{
		Cards: make(map[int]Card),
	}

	for scanner.Scan() {
		cardStr := scanner.Text()
		cardID, winningNumbers, ourNumbers := splitCardAttributes(cardStr)
		card := Card{
			ID:             cardID,
			WinningNumbers: winningNumbers,
			OurNumbers:     ourNumbers,
		}
		deck.AddCard(card)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return deck
}

func (d *Deck) AddCard(card Card) {
	if _, exists := d.Cards[card.ID]; exists {
		d.Count++
	} else {
		log.Fatalf("Card with ID %d does not exist", card.ID)
	}
}

func part_2(filepath string) int {
	deck := NewDeck(filepath)

	maxID := len(deck.Cards)

	for i := 1; i <= maxID; i++ {
		card, exists := deck.Cards[i]
		if !exists {
			log.Fatalf("Card with ID %d does not exist", i)
		}
		matches := countMatches(card.WinningNumbers, card.OurNumbers)

		for j := 0; j < matches; j++ {
			if i+j+1 <= maxID {
				nextCard := deck.Cards[i+j+1]
				deck.AddCard(nextCard)
			} else {
				break
			}
		}
	}

	return deck.Count
}

func main() {
	filepath := "4/small.txt"
	fmt.Println("Solution for part 1: ", part_1(filepath))
	fmt.Println("Solution for part 2: ", part_2(filepath))
}
