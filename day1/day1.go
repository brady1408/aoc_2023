package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const filename = "input.txt"

// const inputPart1 = `1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet`

// const inputPart2 = `two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen`

func main() {
	partOne()
	partTwo()
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

func isNumber(s string) string {
	var numbers = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}

	for k, v := range numbers {
		if p := strings.Index(s, k); p == 0 {
			return v
		}
	}
	return ""
}

func partTwo() {
	calibrationValues := make([]int, 0)
	data := readInput()
	stringSlice := strings.Split(data, "\n")
	for _, s := range stringSlice {
		calibrationStringValues := make([]string, 0)
		for i, v := range s {
			if isDigit(v) {
				calibrationStringValues = append(calibrationStringValues, string(v))
			} else if n := isNumber(s[i:]); n != "" {
				calibrationStringValues = append(calibrationStringValues, n)
			}
		}
		calibrationStringValue := fmt.Sprintf("%s%s", calibrationStringValues[0], calibrationStringValues[len(calibrationStringValues)-1])
		calibrationValue, err := strconv.Atoi(calibrationStringValue)
		if err != nil {
			log.Fatal(err)
		}
		calibrationValues = append(calibrationValues, calibrationValue)
	}

	var sum int
	for _, v := range calibrationValues {
		sum += v
	}
	fmt.Println(sum)
}

func partOne() {
	calibrationValues := make([]int, 0)
	data := readInput()
	stringSlice := strings.Split(data, "\n")
	for _, s := range stringSlice {
		calibrationStringValues := make([]string, 0)
		for _, r := range s {
			if isDigit(r) {
				calibrationStringValues = append(calibrationStringValues, string(r))
			}
		}
		calibrationStringValue := fmt.Sprintf("%s%s", calibrationStringValues[0], calibrationStringValues[len(calibrationStringValues)-1])
		calibrationValue, err := strconv.Atoi(calibrationStringValue)
		if err != nil {
			log.Fatal(err)
		}
		calibrationValues = append(calibrationValues, calibrationValue)
	}

	var sum int
	for _, v := range calibrationValues {
		sum += v
	}
	fmt.Println(sum)
}
