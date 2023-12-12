package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

const tempInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

const filename = "input.txt"

type coords struct {
	row, col int
}

type universe struct {
	layout                        [][]string
	expandedRows, expandedColumns []int
	galaxies                      []coords
	expansion                     int
}

func readInput() string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	// Part 1
	st := time.Now()
	fmt.Println("Part 1: ", partOne(readInput(), 2))
	fmt.Println("Part 1 Time:", time.Since(st))
	st = time.Now()
	// Part 2
	fmt.Println("Part 2: ", partTwo(readInput(), 1000000))
	fmt.Println("Part 2 Time:", time.Since(st))
}

func partTwo(input string, galaxyExpansionRate int) int {
	return partOne(input, galaxyExpansionRate)
}

func partOne(input string, galaxyExpansionRate int) int {
	universe := parseInput(input, galaxyExpansionRate)
	universe.findGalaxies(universe.layout)
	distances := universe.mapDistances()
	sum := 0
	for _, distance := range distances {
		sum += distance
	}
	return sum
}

func parseInput(input string, galaxyExpansionRate int) universe {
	matrix := make([][]string, 0)
	emptyRows := make([]int, 0)
	emptyColumns := make([]int, 0)
	for k, line := range strings.Split(input, "\n") {
		emptyLine := true
		row := make([]string, 0)
		for _, char := range line {
			if char == '#' {
				emptyLine = false
			}
			row = append(row, string(char))
		}
		if emptyLine {
			emptyRows = append(emptyRows, k)
		}
		matrix = append(matrix, row)
	}
	for i := 0; i < len(matrix[0]); i++ {
		emptyColumn := true
		for _, row := range matrix {
			if row[i] == "#" {
				emptyColumn = false
			}
		}
		if emptyColumn {
			emptyColumns = append(emptyColumns, i)
		}
	}
	return universe{matrix, emptyRows, emptyColumns, []coords{}, galaxyExpansionRate - 1}
}

func (u *universe) findGalaxies(matrix [][]string) {
	galaxies := make([]coords, 0)
	for k, row := range matrix {
		for l, char := range row {
			if char == "#" {
				galaxies = append(galaxies, u.expandCoords(coords{k, l}))
			}
		}
	}
	u.galaxies = galaxies
}

func (u universe) expandCoords(c coords) coords {
	newRow, newCol := c.row, c.col
	for _, e := range u.expandedRows {
		if c.row >= e {
			newRow += u.expansion
		}
	}
	for _, e := range u.expandedColumns {
		if c.col >= e {
			newCol += u.expansion
		}
	}
	return coords{newRow, newCol}
}

func (u universe) mapDistances() []int {
	distances := []int{}
	for k, galaxy := range u.galaxies {
		for l, otherGalaxy := range u.galaxies {
			if l <= k {
				continue
			}
			distances = append(distances, u.distance(galaxy, otherGalaxy))
		}
	}
	return distances
}

func (u universe) distance(galaxy, otherGalaxy coords) int {
	return int(math.Abs(float64(galaxy.row-otherGalaxy.row)) + math.Abs(float64(galaxy.col-otherGalaxy.col)))
}
