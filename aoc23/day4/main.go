package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	fmt.Println(solve(string(data)))
}

func solve(data string) int {
	lines := strings.Split(data, "\n")
	lines = lines[:len(lines)-1]

	counts := make([]int, len(lines))
	for i := range counts {
		counts[i]++
	}

	for i, line := range lines {
		game := strings.Split(line, ": ")
		cards := strings.Split(game[1], " | ")

		winningCards := make(map[string]bool)

		winCards := strings.Split(cards[0], " ")
		playerCards := strings.Split(cards[1], " ")

		subtotal := 0

		for _, card := range winCards {
			if card == "" {
				continue
			}
			winningCards[card] = true
		}

		for _, card := range playerCards {
			if winningCards[card] {
				subtotal++
			}
		}


		for j := i + 1; subtotal > 0; subtotal-- {
			counts[j] += counts[i]
			j++
		}
	}

	sum := 0
	for _, c := range counts {
		sum += c
	}
	return sum
}
