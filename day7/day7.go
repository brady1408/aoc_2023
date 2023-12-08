package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

type hands struct {
	hand  string
	score int
	rank  int
}

const (
	highCard int = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

const filename = "input.txt"

func readInput() string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	// Part 1
	fmt.Println("Part 1: ", part1(readInput()))

	// Part 2
	fmt.Println("Part 2: ", part2(readInput()))
}

func parseInput(input string) []hands {
	out := []hands{}
	ss := strings.Split(input, "\n")
	for _, s := range ss {
		h := strings.Split(s, " ")
		n, err := strconv.Atoi(h[1])
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, hands{hand: h[0], score: n})
	}
	return out
}

func getRank(countCards map[string]int) int {
	if len(countCards) == 1 {
		return fiveOfAKind
	}
	if len(countCards) == 2 {
		for _, c := range countCards {
			if c == 4 {
				return fourOfAKind
			}
			if c == 3 {
				return fullHouse
			}
		}
	}
	if len(countCards) == 3 {
		for _, c := range countCards {
			if c == 3 {
				return threeOfAKind
			}
			if c == 2 {
				return twoPair
			}
		}
	}
	if len(countCards) == 4 {
		return onePair
	}
	if len(countCards) == 5 {
		return highCard
	}
	fmt.Println(countCards)
	return -1
}

func rankHands(h []hands) []hands {
	out := []hands{}
	for _, v := range h {
		countCards := make(map[string]int)
		for _, c := range v.hand {
			countCards[string(c)]++
		}
		out = append(out, hands{hand: v.hand, score: v.score, rank: getRank(countCards)})
	}
	return out
}

func part1(input string) int {
	h := parseInput(input)
	r := rankHands(h)
	sort.Slice(r, func(i, j int) bool {
		m := map[rune]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
		if r[i].rank == r[j].rank {
			for k := 0; k < len(r[i].hand); k++ {
				if m[rune(r[i].hand[k])] == m[rune(r[j].hand[k])] {
					continue
				}
				return m[rune(r[i].hand[k])] < m[rune(r[j].hand[k])]
			}
		}
		return r[i].rank < r[j].rank
	})
	sum := 0
	for k, v := range r {
		sum += v.score * (k + 1)
	}
	return sum
}

func rankJokerHands(h []hands) []hands {
	out := []hands{}
	for _, v := range h {
		jokerCount := 0
		countCards := make(map[string]int)
		for _, c := range v.hand {
			if string(c) == "J" {
				jokerCount++
				continue
			}
			countCards[string(c)]++
		}
		highCard := ""
		highScore := 0
		for k, v := range countCards {
			if v > highScore {
				highCard = k
				highScore = v
			}
		}
		countCards[highCard] += jokerCount
		out = append(out, hands{hand: v.hand, score: v.score, rank: getRank(countCards)})
	}
	return out
}

func part2(input string) int {
	h := parseInput(input)
	r := rankJokerHands(h)
	sort.Slice(r, func(i, j int) bool {
		m := map[rune]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 1, 'Q': 12, 'K': 13, 'A': 14}
		if r[i].rank == r[j].rank {
			for k := 0; k < len(r[i].hand); k++ {
				if m[rune(r[i].hand[k])] == m[rune(r[j].hand[k])] {
					continue
				}
				return m[rune(r[i].hand[k])] < m[rune(r[j].hand[k])]
			}
		}
		return r[i].rank < r[j].rank
	})
	sum := 0
	for k, v := range r {
		sum += v.score * (k + 1)
	}
	return sum
}
