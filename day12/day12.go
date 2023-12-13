package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const tempInput = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

type records struct {
	conditions string
	groups     []int
}

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
	// st := time.Now()
	fmt.Println("Part 1: ", partOne(readInput()))
	// fmt.Println("Part 1 Time:", time.Since(st))
	// st = time.Now()
	// Part 2
	fmt.Println("Part 2: ", partTwo(tempInput))
	// fmt.Println("Part 2 Time:", time.Since(st))
	// fmt.Println(math.Pow(2, 5))
	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("%d, %b\n", i, i)
	// }
}

func parseInput(input string, folds int) []records {
	r := []records{}
	for _, line := range strings.Split(input, "\n") {
		ss := strings.Split(line, " ")
		gs := strings.Split(ss[1], ",")
		g := []int{}
		for _, v := range gs {
			gi, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			g = append(g, gi)
		}
		ufs := ""
		ufg := []int{}
		for i := 0; i < folds; i++ {
			ufs += ss[0]
			ufg = append(ufg, g...)
		}
		r = append(r, records{ufs, ufg})
	}
	return r
}

func partOne(input string) int {
	records := parseInput(input, 1)
	results := tryCombinatons(records)
	sum := 0
	for _, v := range results {
		sum += v
	}
	return sum
}

func partTwo(input string) int {
	records := parseInput(input, 5)
	results := tryCombinatons(records)
	sum := 0
	for _, v := range results {
		sum += v
	}
	return sum
}

func tryCombinatons(r []records) []int {
	m := map[rune]string{'0': ".", '1': "#"}
	counts := []int{}
	for _, v := range r {
		count := 0
		ss := strings.Split(v.conditions, "?")
		for i := 0; i < int(math.Pow(2, float64(len(ss)-1))); i++ {
			cond := v.conditions
			bin := reverse(fmt.Sprintf("%b", i))
			for _, bit := range bin {
				cond = strings.Replace(cond, "?", m[bit], 1)
			}
			cond = strings.ReplaceAll(cond, "?", ".")
			test := []int{}
			for _, t := range strings.Split(cond, ".") {
				if len(t) == 0 {
					continue
				}
				test = append(test, len(t))
			}
			if len(test) == len(v.groups) {
				match := true
				for k, t := range test {
					if t != v.groups[k] {
						match = false
					}
				}
				if match {
					count++
				}
			}
		}
		counts = append(counts, count)
	}
	return counts
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
