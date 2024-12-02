package day2

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/egbulmer/advent24/util"
)

func Puzzle1(r io.Reader) (int, error) {
	reports := parseInput(r)
	answer := 0

	for _, report := range reports {
		if isSafe(report) {
			answer += 1
		}
	}

	return answer, nil
}

func Puzzle2(r io.Reader) (int, error) {
	return 0, errors.New("unsolved")
}

func isSafe(report []int) bool {
	dirn := util.SignInt(report[1] - report[0])

	if dirn == 0 {
		// Must be increasing or decreasing.
		return false
	}

	for i := 0; i < len(report)-1; i++ {
		delta := report[i+1] - report[i]
		if util.SignInt(delta) != dirn {
			// Changed direction.
			return false
		}
		dist := util.AbsInt(delta)
		if dist < 0 || dist > 3 {
			// Too far away from previous value.
			return false
		}
	}

	return true
}

func parseInput(r io.Reader) [][]int {
	result := make([][]int, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		report := make([]int, len(parts))

		for i, part := range parts {
			n, _ := strconv.Atoi(part)
			report[i] = n
		}

		result = append(result, report)
	}

	return result
}
