package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const tempInput = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

const filename = "input.txt"

type specialGrid struct {
	grid map[int]map[int]bool
}

func readInput() string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	values := []int{}
	input := readInput()
	cards := strings.Split(input, "\n")
	for _, card := range cards {
		card = strings.ReplaceAll(card, "  ", " ")
		card = strings.TrimSpace(card[strings.Index(card, ":")+1:])
		winningStringNumbers := strings.TrimSpace(card[strings.Index(card, "|")+1:])
		playingStringNumbers := strings.TrimSpace(card[:strings.Index(card, "|")])
		winningNumbers := []int{}
		playingNumbers := []int{}
		for _, n := range strings.Split(winningStringNumbers, " ") {
			digit, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			winningNumbers = append(winningNumbers, digit)
		}
		for _, n := range strings.Split(playingStringNumbers, " ") {
			digit, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			playingNumbers = append(playingNumbers, digit)
		}
		found := 0
		for _, n := range playingNumbers {
			for _, w := range winningNumbers {
				if n == w {
					found++
				}
			}
		}
		if found > 0 {
			values = append(values, 1<<(found-1))
		}
	}
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println(sum)
}

func addCards(cardID, found int, cardCountMap map[int]int) {
	cards := cardCountMap[cardID]
	for i := 1; i <= found; i++ {
		cardCountMap[cardID+i] += cards
	}
}

func partTwo() {
	cardCountMap := make(map[int]int)
	input := readInput()
	for k := range strings.Split(input, "\n") {
		cardCountMap[k+1]++
	}
	for i, card := range strings.Split(input, "\n") {
		cardID := i + 1
		card = strings.ReplaceAll(card, "  ", " ")
		card = strings.TrimSpace(card[strings.Index(card, ":")+1:])
		winningStringNumbers := strings.TrimSpace(card[strings.Index(card, "|")+1:])
		playingStringNumbers := strings.TrimSpace(card[:strings.Index(card, "|")])
		winningNumbers := []int{}
		playingNumbers := []int{}
		for _, n := range strings.Split(winningStringNumbers, " ") {
			digit, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			winningNumbers = append(winningNumbers, digit)
		}
		for _, n := range strings.Split(playingStringNumbers, " ") {
			digit, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			playingNumbers = append(playingNumbers, digit)
		}
		found := 0
		for _, n := range playingNumbers {
			for _, w := range winningNumbers {
				if n == w {
					found++
				}
			}
		}
		if found > 0 {
			addCards(cardID, found, cardCountMap)
		}
	}
	sum := 0
	for _, v := range cardCountMap {
		sum += v
	}
	fmt.Println(sum)
}
