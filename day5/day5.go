package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const tempInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

const filename = "input.txt"

func readInput() string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

type dictionary struct {
	destination int64
	source      int64
	length      int64
}

type database struct {
	seeds                 []int64
	seedToSoil            []dictionary
	soilToFertilizer      []dictionary
	fertilizerToWater     []dictionary
	waterToLight          []dictionary
	lightToTemperature    []dictionary
	temperatureToHumidity []dictionary
	humidityToLocation    []dictionary
}

func (db *database) init() {
	input := readInput()

	category := strings.Split(input, "\n\n")

	seeds := strings.Split(category[0][strings.Index(category[0], ":")+2:], " ")

	db.seeds = make([]int64, len(seeds))
	for i, seed := range seeds {
		n, err := strconv.ParseInt(seed, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		db.seeds[i] = n
	}

	db.seedToSoil = fillDictionary(category[1])
	db.soilToFertilizer = fillDictionary(category[2])
	db.fertilizerToWater = fillDictionary(category[3])
	db.waterToLight = fillDictionary(category[4])
	db.lightToTemperature = fillDictionary(category[5])
	db.temperatureToHumidity = fillDictionary(category[6])
	db.humidityToLocation = fillDictionary(category[7])
}

func (db database) getLocationBySeed(seed int64) int64 {
	var soil, fertilizer, water, light, temperature, humidity, location int64
	soil = getDestination(seed, db.seedToSoil)
	fertilizer = getDestination(soil, db.soilToFertilizer)
	water = getDestination(fertilizer, db.fertilizerToWater)
	light = getDestination(water, db.waterToLight)
	temperature = getDestination(light, db.lightToTemperature)
	humidity = getDestination(temperature, db.temperatureToHumidity)
	location = getDestination(humidity, db.humidityToLocation)

	return location
}

func getDestination(value int64, dict []dictionary) int64 {
	for _, d := range dict {
		if value >= d.source && value < d.source+d.length {
			return d.destination + (value - d.source)
		}
	}
	return value
}

func fillDictionary(s string) []dictionary {
	out := []dictionary{}
	noHeader := s[strings.Index(s, ":")+2:]
	rows := strings.Split(noHeader, "\n")
	for _, row := range rows {
		nums := strings.Split(row, " ")
		destination, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		source, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		length, err := strconv.ParseInt(nums[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, dictionary{destination, source, length})
	}
	return out
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	db := database{}
	db.init()
	location := int64(0)
	for _, seed := range db.seeds {
		l := db.getLocationBySeed(seed)
		if location == 0 {
			location = l
		} else if l < location {
			location = l
		}
	}
	fmt.Println(location)
}

func partTwo() {
	db := database{}
	db.init()
	location := int64(0)
	for i, seed := range db.seeds {
		if i%2 == 0 {
			for j := int64(0); j < db.seeds[i+1]; j++ {
				l := db.getLocationBySeed(seed + j)
				if location == 0 {
					location = l
				} else if l < location {
					location = l
				}
			}
		}
	}
	fmt.Println(location)
}
