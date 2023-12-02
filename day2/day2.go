package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const filename = "input.txt"

// const input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

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
	var sum int
	possibleGameIds := make([]int, 0)
	data := readInput()
	games := strings.Split(data, "\n")
	for _, game := range games {
		var impossible bool
		id, err := strconv.Atoi(game[strings.Index(game, " ")+1 : strings.Index(game, ":")])
		if err != nil {
			log.Fatal(err)
		}
		collection := strings.Split(game[strings.Index(game, ":")+2:], ";")
		for _, item := range collection {
			cubes := strings.Split(strings.TrimSpace(item), ",")
			for _, cube := range cubes {
				cube := strings.TrimSpace(cube)
				number := cube[:strings.Index(cube, " ")]
				i, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}
				color := cube[strings.Index(cube, " ")+1:]
				if isImpossible(color, i) {
					impossible = true
				}
			}
		}
		if impossible {
			continue
		}
		possibleGameIds = append(possibleGameIds, id)
	}
	for _, v := range possibleGameIds {
		sum += v
	}
	fmt.Println(sum)
}

func partTwo() {
	var sum int
	powers := make([]int, 0)
	data := readInput()
	games := strings.Split(data, "\n")
	for _, game := range games {
		neededCubes := make(map[string]int)
		collection := strings.Split(game[strings.Index(game, ":")+2:], ";")
		for _, item := range collection {
			cubes := strings.Split(strings.TrimSpace(item), ",")
			for _, cube := range cubes {
				cube := strings.TrimSpace(cube)
				number := cube[:strings.Index(cube, " ")]
				i, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}
				color := cube[strings.Index(cube, " ")+1:]
				if neededCubes[color] < i {
					neededCubes[color] = i
				}
			}
		}
		var power int
		for _, v := range neededCubes {
			if power == 0 {
				power = v
			} else {
				power *= v
			}
		}
		powers = append(powers, power)
	}
	for _, v := range powers {
		sum += v
	}
	fmt.Println(sum)
}

func isImpossible(color string, number int) bool {
	maxCubes := map[string]int{"red": 12, "green": 13, "blue": 14}
	if number > maxCubes[color] {
		return true
	}
	return false
}
