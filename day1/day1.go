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

// const input = `1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet`

func main() {
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
