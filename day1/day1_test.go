package day1

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var sampleInput = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestPuzzle1Sample(t *testing.T) {
	r := strings.NewReader(sampleInput)
	n, err := Puzzle1(r)

	require.NoError(t, err)
	require.Equal(t, 11, n)
}

func TestPuzzle1(t *testing.T) {
	f, err := os.Open("input.txt")
	defer f.Close()

	require.NoError(t, err)

	n, err := Puzzle1(f)

	require.NoError(t, err)
	require.Equal(t, 1660292, n)
}
