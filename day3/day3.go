package day3

import (
	"errors"
	"io"
	"regexp"
	"strconv"
)

type mul struct{ x, y int }

func Puzzle1(r io.Reader) (int, error) {
	muls := parseInput(r)
	answer := 0

	for _, mul := range muls {
		answer += mul.x * mul.y
	}

	return answer, nil
}

func Puzzle2(r io.Reader) (int, error) {
	return 0, errors.New("unsolved")
}

func parseInput(r io.Reader) []mul {
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
