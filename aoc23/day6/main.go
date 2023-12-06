package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	ans := solve_part_2(string(data))
	fmt.Println("Answer: ", ans)
}

func solve_part_1(data string) int {
	lines := strings.Split(data, "\n")

	times, _ := convertStringToFloatSlice(strings.Split(lines[0], ":")[1])
	distances, _ := convertStringToFloatSlice(strings.Split(lines[1], ":")[1])

	total := 1
	fmt.Println(times, distances)

	for i := range times {
		time := times[i]
		distance := distances[i]

		D := math.Sqrt(float64(time*time - 4*distance))

		r0 := (time - D) / 2
		r1 := (time + D) / 2

		a := math.Ceil(r0)
		b := math.Floor(r1)

		// m for exluding edges (strictly greater)
		m := 0
		if r0 == a {
			m++
		}
		if r1 == b {
			m++
		}

		total *= int(b-a+1) - m
	}

	return total
}

func solve_part_2(data string) int {
	lines := strings.Split(data, "\n")

	time := convertStringToFloat(strings.Split(lines[0], ":")[1])
	distance := convertStringToFloat(strings.Split(lines[1], ":")[1])

	fmt.Println(time, distance)

	D := math.Sqrt(float64(time*time - 4*distance))

	r0 := (time - D) / 2
	r1 := (time + D) / 2

	a := math.Ceil(r0)
	b := math.Floor(r1)

	fmt.Println(r0, r1, a, b)

	m := 0
	if r0 == a {
		m++
	}
	if r1 == b {
		m++
	}

	return int(b-a+1) - m
}

func convertStringToFloat(str string) float64 {

	number := 0
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			number = number*10 + int(ch-'0')
		}
	}

	return float64(number)
}

func convertStringToFloatSlice(input string) ([]float64, error) {
	// Split the input string by white spaces
	strSlice := strings.Fields(input)

	// Convert each substring to an integer
	var floatSlice []float64
	for _, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		floatSlice = append(floatSlice, float64(num))
	}

	return floatSlice, nil
}
