package day3

import (
	"io"
	"regexp"
	"strconv"
)

type mul struct{ x, y int }

func Puzzle1(r io.Reader) (int, error) {
	muls := parseInput1(r)
	answer := process(muls)
	return answer, nil
}

func Puzzle2(r io.Reader) (int, error) {
	muls := parseInput2(r)
	answer := process(muls)
	return answer, nil
}

func process(muls []mul) int {
	answer := 0
	for _, mul := range muls {
		answer += mul.x * mul.y
	}
	return answer
}

func parseInput1(r io.Reader) []mul {
	result := make([]mul, 0)
	data, _ := io.ReadAll(r)

	re := regexp.MustCompile(`mul\((\d+?),(\d+?)\)`)
	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		result = append(result, mul{x, y})
	}

	return result
}

func parseInput2(r io.Reader) []mul {
	result := make([]mul, 0)
	data, _ := io.ReadAll(r)

	re := regexp.MustCompile(`(?:mul\((\d+?),(\d+?)\)|do(n't)?\(\))`)
	matches := re.FindAllStringSubmatch(string(data), -1)

	enabled := true

	for _, match := range matches {
		switch {
		case match[0] == "do()":
			enabled = true
		case match[0] == "don't()":
			enabled = false
		case enabled:
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			result = append(result, mul{x, y})
		}

	}

	return result
}
