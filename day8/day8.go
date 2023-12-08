package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const input1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const input2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const input3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

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

type puzzle struct {
	directions string
	coords     map[string][]string
}

func parsePuzzle(in string) puzzle {
	p := puzzle{coords: map[string][]string{}}
	for k, s := range strings.Split(in, "\n") {
		if k == 0 {
			p.directions = s
			continue
		}
		if s == "" {
			continue
		}
		p.coords[s[:3]] = strings.Split(s[7:len(s)-1], ", ")
	}
	return p
}

func part1(input string) int {
	p := parsePuzzle(input)
	m := map[string]int{"L": 0, "R": 1}
	next := "AAA"
	for i := 0; true; i++ {
		if next == "ZZZ" {
			return i
		}
		next = p.coords[next][m[string(p.directions[i%len(p.directions)])]]
	}
	return -1
}

// Ok this is going to be a complete shot in the dark
func part2(input string) int {
	p := parsePuzzle(input)
	m := map[string]int{"L": 0, "R": 1}
	startingPoints := map[string]int{}
	for k := range p.coords {
		if k[2] == 'A' {
			startingPoints[k] = 0
		}
	}
	for k := range startingPoints {
		next := k
		for i := 0; true; i++ {
			if next[2] == 'Z' {
				startingPoints[k] = i
				break
			}
			next = p.coords[next][m[string(p.directions[i%len(p.directions)])]]
		}
	}
	nums := []int{}
	for _, v := range startingPoints {
		nums = append(nums, v)
	}
	return LCM(nums[0], nums[1], nums[2:]...)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
