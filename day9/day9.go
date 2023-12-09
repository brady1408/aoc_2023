package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	fmt.Println(partOne(readInput()))
	fmt.Println(partTwo(tempInput))
}

func predictNext(in []int) int {
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
	num := 0
	for i := len(predictMatrix) - 1; i >= 0; i-- {
		num += predictMatrix[i][len(predictMatrix[i])-1]
	}
	return num
}

func partOne(input string) int {
	puzzle := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		rowString := strings.Fields(line)
		row := []int{}
		for _, numString := range rowString {
			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		puzzle = append(puzzle, row)
	}
	sum := 0
	for _, v := range puzzle {
		sum += predictNext(v)
	}
	return sum
}

func predictPrevious(in []int) int {
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
	fmt.Println(predictMatrix)
	num := 0
	for i := len(predictMatrix) - 1; i >= 0; i-- {
		if i == 0 {
			num += predictMatrix[i][0]
		} else {
			num -= predictMatrix[i][0]
		}
	}
	return num
}

func partTwo(input string) int {
	puzzle := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		rowString := strings.Fields(line)
		row := []int{}
		for _, numString := range rowString {
			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		puzzle = append(puzzle, row)
	}
	sum := 0
	for _, v := range puzzle {
		fmt.Println(predictPrevious(v))
		sum += predictPrevious(v)
	}
	return sum
}
