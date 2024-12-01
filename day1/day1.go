package day1

import (
	"bufio"
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
	left, right, err := parseInput(r)
	if err != nil {
		return 0, err
	}

	incidence := make(map[int]int)
	for _, num := range right {
		count, _ := incidence[num]
		incidence[num] = count + 1
	}

	similarity := 0
	for _, num := range left {
		score, _ := incidence[num]
		similarity += num * score
	}

	return similarity, nil
}

// parseInput consumes the reader line-by-line, parsing a left and right number
// from each list.
func parseInput(r io.Reader) ([]int, []int, error) {
	left := make([]int, 0)
	right := make([]int, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 0 {
			// Line is empty.
			continue
		}

		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])

		left = append(left, int(l))
		right = append(right, int(r))
	}

	return left, right, nil
}
