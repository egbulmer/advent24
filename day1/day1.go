package day1

import (
	"bufio"
	"errors"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Puzzle1(r io.Reader) (int, error) {
	left, right, err := parseInput(r)
	if err != nil {
		return 0, err
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0

	for i := 0; i < len(left); i++ {
		n := left[i] - right[i]
		if n < 0 {
			n = -n
		}
		distance += n
	}

	return distance, nil
}

func Puzzle2(r io.Reader) (int, error) {
	return 0, errors.New("unsolved")
}

// parseInput consumes the reader line-by-line, parsing a left and right number
// from each list.
func parseInput(r io.Reader) ([]int, []int, error) {
	left := make([]int, 0)
	right := make([]int, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Split(scanner.Text(), " ")

		lpart := parts[0]
		rpart := parts[len(parts)-1]

		l, err := strconv.ParseInt(strings.TrimSpace(lpart), 10, 32)
		if err != nil {
			return nil, nil, err
		}
		left = append(left, int(l))

		r, err := strconv.ParseInt(strings.TrimSpace(rpart), 10, 32)
		if err != nil {
			return nil, nil, err
		}
		right = append(right, int(r))
	}

	return left, right, nil
}