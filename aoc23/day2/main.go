package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("failed to read file")
	}

	println(solve(string(file)))
}

func solve(data string) int {

	lines := strings.Split(strings.TrimSpace(data), "\n")
	total := 0

	// colorMap := map[string]int{
	// 	"red":   12,
	// 	"green": 13,
	// 	"blue":  14,
	// }

	for _, line := range lines {
		game := strings.Split(line, ":")[1]
		rolls := strings.Split(game, ";")

		// lineValid := true
		red, blue, green := 0, 0, 0

		for _, roll := range rolls {
			colors := strings.Split(roll, ",")

			for _, c := range colors {
				temp := strings.Split(strings.TrimSpace(c), " ")
				numStr, col := temp[0], temp[1]
				num, _ := strconv.Atoi(numStr)

				switch col {
				case "red":
					red = max(red, num)
				case "green":
					green = max(green, num)
				case "blue":
					blue = max(blue, num)
				}
			}
		}

		total += red * green * blue
	}

	return total
}
