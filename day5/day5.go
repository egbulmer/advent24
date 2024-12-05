package day5

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Puzzle1(r io.Reader) (int, error) {
	rules, updates := parseInput(r)

	// For every number, find all the numbers that it must proceed. For example, in the sample
	// input 61 must proceed 75 and 97, 13 must proceed 29, 47, 53, 61 and 75. 97 does not proceed
	// any other number.
	mustProceed := make(map[int][]int)
	for _, rule := range rules {
		numbers, ok := mustProceed[rule.after]
		if !ok {
			numbers = make([]int, 0)
		}
		numbers = append(numbers, rule.before)
		mustProceed[rule.after] = numbers
	}

	answer := 0

	for _, update := range updates {
		if update.isOrdered(mustProceed) {
			fmt.Println("ordered\t\t", update)
			answer += update.middleNumber()
		} else {
			fmt.Println("unordered\t", update)
		}
	}

	return answer, nil
}

func Puzzle2(r io.Reader) (int, error) {
	return 0, errors.New("unsolved")
}

type rule struct {
	before, after int
}

type update []int

func (u update) middleNumber() int {
	return u[len(u)/2]
}

func (u update) isOrdered(mustProceed map[int][]int) bool {

	for i := len(u) - 1; i > 0; i -= 1 {
		curr, prev := u[i], u[i-1]

		// Lookup all the numbers that this number must proceed.
		numbers, ok := mustProceed[curr]
		if !ok {
			return i != 1
		}

		// Check that the number before it is in this list. If it's not then it is
		// out-of-order.
		index := slices.Index(numbers, prev)
		if index == -1 {
			return false
		}
	}

	return true
}

func parseInput(r io.Reader) ([]rule, []update) {
	rules := make([]rule, 0)
	updates := make([]update, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			before, _ := strconv.Atoi(parts[0])
			after, _ := strconv.Atoi(parts[1])
			rules = append(rules, rule{before, after})
		} else if strings.Contains(line, ",") {
			update := make([]int, 0)
			for _, part := range strings.Split(line, ",") {
				page, _ := strconv.Atoi(part)
				update = append(update, page)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates
}
