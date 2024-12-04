package day4

import (
	"testing"

	"github.com/egbulmer/advent24/util"
)

func TestPuzzle1Sample(t *testing.T) {
	util.TestPuzzle(t, Puzzle1, "sample.txt", 18)
}

func TestPuzzle1(t *testing.T) {
	util.TestPuzzle(t, Puzzle1, "input.txt", 2562)
}

func TestPuzzle2Sample(t *testing.T) {
	util.TestPuzzle(t, Puzzle2, "sample.txt", 4)
}

func TestPuzzle2(t *testing.T) {
	util.TestPuzzle(t, Puzzle2, "input.txt", 566)
}
