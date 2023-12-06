package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	ans := solve_part_2(string(data))
	fmt.Println("Answer: ", ans)
}

func stringsToNum(array []string) ([]int, error) {
	returnSlice := make([]int, len(array))
	for i, str := range array {
		num, error := strconv.Atoi(str)
		if error != nil {
			return []int{}, error
		}
		returnSlice[i] = num
	}

	return returnSlice, nil
}

func solve(data string) int {
	paragraphs := strings.Split(data, "\n\n")
	// fmt.Println(lines[2])
	seeds, _ := stringsToNum(strings.Split(strings.Split(paragraphs[0], ": ")[1], " "))

	paragraphs = paragraphs[1:]
	ranges := make([][][]int, len(paragraphs))

	for i, paragraph := range paragraphs {
		lines := strings.Split(strings.TrimSpace(paragraph), "\n")
		lines = lines[1:]

		range_ := make([][]int, len(lines))
		for i, line := range lines {
			r, _ := stringsToNum(strings.Split(line, " "))
			range_[i] = r
		}

		ranges[i] = range_
	}

	locations := append([]int{}, seeds...)
	for _, range_ := range ranges {
		for j := range locations {
			locations[j] = mapper(range_, locations[j])
		}
	}

	return slices.Min(locations)
}

func mapper(haystack [][]int, needle int) int {
	for _, hay := range haystack {
		d := needle - hay[1]
		if d >= 0 && d < hay[2] {
			return hay[0] + d
		}
	}

	return needle
}

func solve_part_2(data string) int {
	paragraphs := strings.Split(data, "\n\n")
	rawSeeds, _ := stringsToNum(strings.Split(strings.Split(paragraphs[0], ": ")[1], " "))

	paragraphs = paragraphs[1:]
	ranges := make([][][]int, len(paragraphs))

	for i, paragraph := range paragraphs {
		lines := strings.Split(strings.TrimSpace(paragraph), "\n")
		lines = lines[1:]

		range_ := make([][]int, len(lines))
		for i, line := range lines {
			r, _ := stringsToNum(strings.Split(line, " "))
			range_[i] = r
		}

		ranges[i] = range_
	}

	seeds := make([][]int, len(rawSeeds)/2)
	// fmt.Println(rawSeeds, len(rawSeeds), len(seeds))
	for i := 0; i < len(rawSeeds); i += 2 {
		seeds[i/2] = []int{rawSeeds[i], rawSeeds[i] + rawSeeds[i+1]}
	}

	// fmt.Println(seeds)

	for _, range_ := range ranges {
		newSeedRanges := make([][]int, 0)

		for len(seeds) > 0 {
			seed := seeds[len(seeds)-1]
			seeds = seeds[:len(seeds)-1]

			ss, se := seed[0], seed[1]
			// fmt.Println("Seed", ss, se)

			fnd := false
			for _, rule := range range_ {
				rs, re := rule[1], rule[1]+rule[2]
				off := rule[0] - rule[1]

				// fmt.Println(rs, re)

				is := max(rs, ss)
				ie := min(re, se)

				// fmt.Println("Intersection: ", is, ie)

				if is < ie {
					newSeedRanges = append(newSeedRanges, []int{is + off, ie + off}) // Map with offset

					if is > ss {
						seeds = append(seeds, []int{ss, is})
					}
					if ie < se {
						seeds = append(seeds, []int{ie, se})
					}

					fnd = true

					break
				}
				// fmt.Println("Seeds :", seeds)
				// fmt.Println("New Seeds :", newSeedRanges)

			}
			if !fnd {
				newSeedRanges = append(newSeedRanges, []int{ss, se}) // No mapping found
			}
			// fmt.Println("Seeds: ", seeds, newSeedRanges)
		}

		seeds = newSeedRanges
		// fmt.Println(seeds)
		// break
	}

	minimum := int(math.Inf(0))
	fmt.Println(seeds)
	for _, seed := range seeds {
		minimum = min(seed[0], minimum)
	}

	return minimum
}
