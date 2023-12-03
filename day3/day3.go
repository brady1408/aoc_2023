package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

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

func isDigit(r rune) bool {
	var digitCheck = regexp.MustCompile(`^[0-9]+$`)
	return digitCheck.MatchString(string(r))
}

func newSpecialGrid(input string) *specialGrid {
	var isNumberOrDot = regexp.MustCompile(`^[0-9\.]+$`)
	g := &specialGrid{
		grid: make(map[int]map[int]bool),
	}

	data := readInput()
	for y, line := range strings.Split(data, "\n") {
		for x, char := range line {
			if !isNumberOrDot.MatchString(string(char)) {
				if g.grid[y] == nil {
					g.grid[y] = make(map[int]bool)
				}
				g.grid[y][x] = true
			}
		}
	}
	return g
}

func (s specialGrid) touchedBySpecial(x, y int) bool {
	startX := x - 1
	startY := y - 1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s.grid[startY+i][startX+j] {
				return true
			}
		}
	}

	return false
}

func main() {
	partOne()
}

func partOne() {
	g := newSpecialGrid(input)
	allTouched := []int{}
	data := readInput()
	for y, line := range strings.Split(data, "\n") {
		number := ""
		touched := false
		for x, char := range line {
			if isDigit(char) {
				number += string(char)
				if g.touchedBySpecial(x, y) {
					touched = true
				}
			}
			if x == len(line)-1 || !isDigit(char) {
				if touched && number != "" {
					n, err := strconv.Atoi(number)
					if err != nil {
						log.Panic(err)
					}
					allTouched = append(allTouched, n)
				}
				touched = false
				number = ""
			}
		}

	}
	sum := 0
	for _, n := range allTouched {
		sum += n
	}
	fmt.Println(allTouched)
	fmt.Println(sum)
}
