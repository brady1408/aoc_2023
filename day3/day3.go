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

func newSpecialGrid(gridInput string) *specialGrid {
	var isNumberOrDot = regexp.MustCompile(`^[0-9\.]+$`)
	g := &specialGrid{
		grid: make(map[int]map[int]bool),
	}

	for y, line := range strings.Split(gridInput, "\n") {
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
	partTwo()
}

func partOne() {
	allTouched := []int{}
	data := readInput()
	g := newSpecialGrid(data)
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
	fmt.Println(sum)
}

type gearGrid struct {
	grid map[int]map[int]bool
}

func (g gearGrid) getTouchedLocation(x, y int) []string {
	startX := x - 1
	startY := y - 1
	touches := []string{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.grid[startY+i][startX+j] {
				touches = append(touches, fmt.Sprintf("%d:%d", startX+j, startY+i))
			}
		}
	}

	return touches
}

func newGearGrid(input string) *gearGrid {
	g := &gearGrid{
		grid: make(map[int]map[int]bool),
	}

	// data := readInput()
	data := input
	for y, line := range strings.Split(data, "\n") {
		for x, char := range line {
			if char == '*' {
				if g.grid[y] == nil {
					g.grid[y] = make(map[int]bool)
				}
				g.grid[y][x] = true
			}
		}
	}
	return g
}

func partTwo() {
	data := readInput()
	allTouched := make(map[string]map[int]bool)
	g := newGearGrid(data)
	for y, line := range strings.Split(data, "\n") {
		number := ""
		touchedGears := make(map[string]bool)
		for x, char := range line {
			if isDigit(char) {
				number += string(char)
				if touches := g.getTouchedLocation(x, y); len(touches) > 0 {
					for _, s := range touches {
						touchedGears[s] = true
					}
				}
			}
			if x == len(line)-1 || !isDigit(char) {
				if len(touchedGears) > 0 && number != "" {
					n, err := strconv.Atoi(number)
					if err != nil {
						log.Panic(err)
					}
					for k := range touchedGears {
						if allTouched[k] == nil {
							allTouched[k] = make(map[int]bool)
						}
						allTouched[k][n] = true
					}
				}
				touchedGears = make(map[string]bool)
				number = ""
			}
		}
	}
	sum := 0
	ratios := []int{}
	for _, v := range allTouched {
		if len(v) == 2 {
			ratio := 0
			for k := range v {
				if ratio == 0 {
					ratio = k
				} else {
					ratio *= k
				}
			}
			ratios = append(ratios, ratio)
		}
	}
	for _, n := range ratios {
		sum += n
	}
	fmt.Println(sum)
}
