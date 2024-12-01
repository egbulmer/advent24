package day1

import (
	"testing"

	"github.com/egbulmer/advent24/testutil"
)

func TestPuzzle1Sample(t *testing.T) {
	testutil.TestPuzzle(t, Puzzle1, "sample.txt", 11)
}

func TestPuzzle1(t *testing.T) {
	testutil.TestPuzzle(t, Puzzle1, "input.txt", 1660292)
}
