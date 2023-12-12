package main

import (
	"fmt"
	"math"
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

func main() {
	// Part 1
	st := time.Now()
	fmt.Println("Part 1: ", partOne(tempInput))
	fmt.Println("Part 1 Time:", time.Since(st))
	st = time.Now()
	// Part 2
	// fmt.Println("Part 2: ", partTwo(readInput()))
	// fmt.Println("Part 2 Time:", time.Since(st))
}

func partOne(input string) int {
	matrix := parseInput(input)
	galaxies := findGalaxies(matrix)
	fmt.Println(galaxies)
	return 0
}

func parseInput(input string) [][]string {
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
	expandedMatrix := expandMatrix(matrix, emptyRows, emptyColumns)
	return expandedMatrix
}

func expandMatrix(matrix [][]string, expandRow, expandColumn []int) [][]string {
	newMatrix := make([][]string, 0)
	for k, row := range matrix {
		newRow := make([]string, 0)
		for k, char := range row {
			newRow = append(newRow, char)
			for _, extra := range expandColumn {
				if k == extra {
					newRow = append(newRow, char)
				}
			}
		}
		newMatrix = append(newMatrix, newRow)
		for _, extra := range expandRow {
			if k == extra {
				newMatrix = append(newMatrix, newRow)
			}
		}
	}
	return newMatrix
}

type coords struct {
	row, col int
}

func findGalaxies(matrix [][]string) []coords {
	galaxies := make([]coords, 0)
	for k, row := range matrix {
		for l, char := range row {
			if char == "#" {
				galaxies = append(galaxies, coords{k, l})
			}
		}
	}
	return galaxies
}

func mapDistances(galaxies []coords) map[coords]map[coords]int {
	distances := make(map[coords]map[coords]int)
	for k, galaxy := range galaxies {
		distances[galaxy] = make(map[coords]int)
		for l, otherGalaxy := range galaxies {
			if k != l {
				distances[galaxy][otherGalaxy] = distance(galaxy, otherGalaxy)
			}
		}
	}
	return distances
}

func distance(a, b coords) int {
	return int(math.Abs(float64(a.row-b.row)) + math.Abs(float64(a.col-b.col)))
}
