package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const tempInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

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
	st := time.Now()
	fmt.Println("Part 1: ", partOne(readInput()))
	fmt.Println("Part 1 Time:", time.Since(st))
	st = time.Now()
	// Part 2
	fmt.Println("Part 2: ", partTwo(readInput()))
	fmt.Println("Part 2 Time:", time.Since(st))
}

func partOne(input string) int {
	puzzle := parsePuzzle(input)
	sum := 0
	for _, v := range puzzle {
		sum += predictNext(v)
	}
	return sum
}

func partTwo(input string) int {
	puzzle := parsePuzzle(input)
	sum := 0
	for _, v := range puzzle {
		sum += predictPrevious(v)
	}
	return sum
}

func predictNext(in []int) int {
	predictMatrix := createPredictMatrix(in)
	num := 0
	for i := len(predictMatrix) - 1; i >= 0; i-- {
		num += predictMatrix[i][len(predictMatrix[i])-1]
	}
	return num
}

func predictPrevious(in []int) int {
	predictMatrix := createPredictMatrix(in)
	num := 0
	for i := len(predictMatrix) - 1; i >= 0; i-- {
		num = (num * -1) + predictMatrix[i][0]
	}
	return num
}

func createPredictMatrix(in []int) [][]int {
	predictMatrix := [][]int{}
	predictMatrix = append(predictMatrix, in)
	for i := 0; true; i++ {
		step := []int{}
		for k, v := range predictMatrix[i] {
			if k == 0 {
				continue
			}
			step = append(step, v-predictMatrix[i][k-1])
		}
		predictMatrix = append(predictMatrix, step)
		allZero := true
		for _, v := range step {
			if v != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			break
		}
	}
	return predictMatrix
}

func parsePuzzle(in string) [][]int {
	p := [][]int{}
	for _, s := range strings.Split(in, "\n") {
		rowString := strings.Fields(s)
		row := []int{}
		for _, numString := range rowString {
			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		p = append(p, row)
	}
	return p
}
