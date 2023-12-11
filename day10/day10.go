package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

const tempInput = `.....
.S-7.
.|.|.
.L-J.
.....`

const tempInput2 = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

const filename = "input.txt"

func readInput() string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

type movement struct {
	x, y    int
	exitDir string
}

const (
	// pipes
	NSPipe = "|"
	EWPipe = "-"
	NEBend = "L"
	NWBend = "J"
	SWBend = "7"
	SEBemd = "F"
	ground = "."
	start  = "S"

	// sides
	north = "N"
	south = "S"
	east  = "E"
	west  = "W"
)

type movementMap map[string]map[string]movement

type island struct {
	startX, startY     int
	currentX, currentY int
	currentDir         string
	visited            []string
	movement           movementMap
	grid               groundGrid
}

func (i *island) reset() {
	i.currentX = i.startX
	i.currentY = i.startY
	i.currentDir = ""
	i.visited = []string{}
}

func (i *island) move(d string) error {
	i.currentDir = d
	startMap := map[string][]int{north: {0, -1}, south: {0, 1}, east: {1, 0}, west: {-1, 0}}
	if i.currentX == i.startX && i.currentY == i.startY {
		i.visited = append(i.visited, start)
		i.currentX += startMap[i.currentDir][0]
		i.currentY += startMap[i.currentDir][1]
	}
	for {
		if i.grid[i.currentY][i.currentX] == ground {
			i.reset()
			return errors.New("hit ground")
		}
		if m, ok := i.movement[i.grid[i.currentY][i.currentX]][i.currentDir]; ok {
			i.currentX += m.x
			i.currentY += m.y
			i.currentDir = m.exitDir
			i.visited = append(i.visited, i.grid[i.currentY][i.currentX])
		} else {
			i.reset()
			return errors.New("hit dead end")
		}
		if i.currentX == i.startX && i.currentY == i.startY {
			return nil
		}
	}
}

// x and y start at the upper left corner
func makeMap() movementMap {
	m := make(map[string]map[string]movement)
	m[NSPipe] = make(map[string]movement)
	m[NSPipe][south] = movement{0, 1, south}
	m[NSPipe][north] = movement{0, -1, north}
	m[EWPipe] = make(map[string]movement)
	m[EWPipe][east] = movement{1, 0, east}
	m[EWPipe][west] = movement{-1, 0, west}
	m[NEBend] = make(map[string]movement)
	m[NEBend][south] = movement{1, 0, east}
	m[NEBend][west] = movement{0, -1, north}
	m[NWBend] = make(map[string]movement)
	m[NWBend][south] = movement{-1, 0, west}
	m[NWBend][east] = movement{0, -1, north}
	m[SWBend] = make(map[string]movement)
	m[SWBend][north] = movement{-1, 0, west}
	m[SWBend][east] = movement{0, 1, south}
	m[SEBemd] = make(map[string]movement)
	m[SEBemd][north] = movement{1, 0, east}
	m[SEBemd][west] = movement{0, 1, south}

	return m
}

type groundGrid [][]string

func parseInput(input string) groundGrid {
	lines := strings.Split(input, "\n")
	m := make(groundGrid, len(lines))
	for i, line := range lines {
		m[i] = strings.Split(line, "")
	}
	return m
}

func (g groundGrid) findStart() (int, int) {
	for y, row := range g {
		for x, col := range row {
			if col == start {
				return x, y
			}
		}
	}
	return -1, -1
}

func initIsland(input string) *island {
	m := makeMap()
	g := parseInput(input)
	x, y := g.findStart()
	return &island{
		startX:     x,
		startY:     y,
		currentX:   x,
		currentY:   y,
		currentDir: "",
		visited:    []string{},
		movement:   m,
		grid:       g,
	}
}

func main() {
	// Part 1
	st := time.Now()
	fmt.Println("Part 1: ", partOne(readInput()))
	fmt.Println("Part 1 Time:", time.Since(st))
	st = time.Now()
	// Part 2
	// fmt.Println("Part 2: ", partTwo(readInput()))
	// fmt.Println("Part 2 Time:", time.Since(st))
}

func partOne(input string) int {
	island := initIsland(input)
	d := []string{south, east, north, west}
	for _, v := range d {
		err := island.move(v)
		if err == nil {
			break
		}
		fmt.Println(v, ": ", err)
	}
	fmt.Println(island.visited, len(island.visited))
	return int(math.Ceil(float64(len(island.visited)) / float64(2)))
}
