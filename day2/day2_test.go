package day2

import (
	"testing"

	"github.com/egbulmer/advent24/util"
)

func TestPuzzle1Sample(t *testing.T) {
	util.TestPuzzle(t, Puzzle1, "sample.txt", 2)
}

func TestPuzzle1(t *testing.T) {
	util.TestPuzzle(t, Puzzle1, "input.txt", 526)
}

func TestPuzzle2Sample(t *testing.T) {
	util.TestPuzzle(t, Puzzle2, "sample.txt", 0)
}

func TestPuzzle2(t *testing.T) {
	util.TestPuzzle(t, Puzzle2, "input.txt", 0)
}
