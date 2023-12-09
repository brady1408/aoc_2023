package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const tempInput = `Time:      7  15   30
Distance:  9  40  200`

const filename = "input.txt"

func readInput() string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	st := time.Now()
	partOne()
	fmt.Println("nano", time.Since(st)*time.Nanosecond)
	st = time.Now()
	partTwo()
	fmt.Println("nano", time.Since(st)*time.Nanosecond)
}

type records struct {
	time           []int
	distance       []int
	singleTime     int
	singleDistance int
}

func newRecords() *records {
	out := &records{}
	ss := strings.Split(readInput(), "\n")
	// ss := strings.Split(tempInput, "\n")
	timesString := strings.Fields(ss[0][strings.Index(ss[0], ":")+1:])
	for _, timeString := range timesString {
		time, err := strconv.Atoi(timeString)
		if err != nil {
			log.Fatal(err)
		}
		out.time = append(out.time, time)
	}
	destsString := strings.Fields(ss[1][strings.Index(ss[1], ":")+1:])
	for _, destString := range destsString {
		dest, err := strconv.Atoi(destString)
		if err != nil {
			log.Fatal(err)
		}
		out.distance = append(out.distance, dest)
	}
	singleTimeString := strings.ReplaceAll(ss[0][strings.Index(ss[0], ":")+1:], " ", "")
	singleDistanceString := strings.ReplaceAll(ss[1][strings.Index(ss[1], ":")+1:], " ", "")
	singleTime, err := strconv.Atoi(singleTimeString)
	if err != nil {
		log.Fatal(err)
	}
	singleDistance, err := strconv.Atoi(singleDistanceString)
	if err != nil {
		log.Fatal(err)
	}
	out.singleTime = singleTime
	out.singleDistance = singleDistance
	return out
}

func partOne() {
	r := newRecords()
	total := 0
	for i, v := range r.time {
		first := 0
		last := 0
		// find first
		for j := 0; j < v; j++ {
			if (v-j)*j > r.distance[i] {
				first = j
				break
			}
		}
		// find last
		for j := v; j > 0; j-- {
			if (v-j)*j > r.distance[i] {
				last = j
				break
			}
		}
		ways := last - first + 1
		if total == 0 {
			total = ways
		} else {
			total *= ways
		}
	}
	fmt.Println(total)
}

func partTwo() {
	r := newRecords()
	first := 0
	last := 0
	calc := 0
	// find first
	for j := 0; j < r.singleTime; j++ {
		calc++
		if (r.singleTime-j)*j > r.singleDistance {
			first = j
			break
		}
	}
	// find last
	for j := r.singleTime; j > 0; j-- {
		calc++
		if (r.singleTime-j)*j > r.singleDistance {
			last = j
			break
		}
	}
	ways := last - first + 1
	fmt.Println(ways)
	fmt.Println("calculations made:", calc)
}
