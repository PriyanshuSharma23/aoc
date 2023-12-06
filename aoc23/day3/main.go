package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Println(part2(string(data)))
}

func solve(data string) int {
	lines := strings.Split(data, "\n")
	lines = lines[0 : len(lines)-1]
	// fmt.Println(lines, len(lines))
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	charPresent := false
	builder, total := 0, 0

	h, w := len(grid), len(grid[0])

	for i := 0; i < h; i++ {
		// fmt.Printf("Parsing %s\n", string(grid[i]))
		for j := 0; j < w; j++ {
			if unicode.IsDigit(rune(grid[i][j])) {
				builder = builder*10 + int(grid[i][j]-'0')
				if !charPresent {
					charPresent = isSymbolInSorroundings(grid, i, j, h, w)
				}
			} else {
				if builder != 0 && charPresent {
					total += builder
					// println("adding", builder)
				}

				builder = 0
				charPresent = false
			}
			// println(builder, charPresent)
		}

		if builder != 0 && charPresent {
			// println("adding", builder)
			total += builder
		}
		builder = 0
		charPresent = false
	}

	return total
}

var directionsX []int = []int{-1, 0, 1}
var directionsY []int = []int{-1, 0, 1}

func isSymbol(s byte) bool {
	return !unicode.IsDigit(rune(s)) && s != '.'
}

func isSymbolInSorroundings(grid [][]byte, i int, j int, h int, w int) bool {
	for _, dx := range directionsX {
		for _, dy := range directionsY {
			if dx == 0 && dy == 0 {
				continue
			}

			i_ := i + dx
			j_ := j + dy

			if i_ < h && j_ < w && i_ >= 0 && j_ >= 0 {
				if isSymbol(grid[i_][j_]) {
					return true
				}
			}
		}
	}

	return false
}

func part2(data string) int {
	lines := strings.Split(data, "\n")
	lines = lines[0 : len(lines)-1]
	// fmt.Println(lines, len(lines))
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	var total int = 0
	h, w := len(grid), len(grid[0])

	for i := 0; i < h; i++ {
		// fmt.Printf("Parsing %s\n", string(grid[i]))
		for j := 0; j < w; j++ {
		if grid[i][j] == '*' {
				neighbors := getNeighboringNums(grid, i, j, h, w)
				fmt.Println(neighbors)
				if len(neighbors) == 2 {
					total += neighbors[0] * neighbors[1]
				}
			}
		}
	}

	return total
}

func getNeighboringNums(grid [][]byte, i int, j int, h int, w int) []int {
	nums := []int{}

	sides := [][][]int{
		{
			{-1, -1}, {-1, 0}, {-1, 1},
		},
		{
			{0, -1},
		},
		{
			{0, 1},
		},
		{
			{1, -1}, {1, 0}, {1, 1},
		},
	}

	for _, side := range sides {
		continous := false
		for _, points := range side {
			di, dj := points[0], points[1]
			i_ := i + di
			j_ := j + dj

			if inBounds(i_, j_, h, w) && unicode.IsDigit(rune(grid[i_][j_])) {
				if !continous {
					nums = append(nums, propNum(grid, i_, j_, h, w))
				}
				continous = true
			} else {
				continous = false
			}
		}
	}

	return nums
}

func inBounds(i, j, h, w int) bool {
	return i < h && j < w && i >= 0 && j >= 0
}

func propNum(grid [][]byte, i, j, h, w int) int {
	number := 0
	iter := j
	for ; inBounds(i, iter, h, w) && unicode.IsDigit(rune(grid[i][iter])); iter-- {
	}

	iter++ // back to start of number
	for ; inBounds(i, iter, h, w) && unicode.IsDigit(rune(grid[i][iter])); iter++ {
		number = number*10 + int(grid[i][iter]-'0')
	}

	return number
}
