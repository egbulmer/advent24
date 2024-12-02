package util

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type PuzzleFunc func(io.Reader) (int, error)

func TestPuzzle(t *testing.T, puzzle PuzzleFunc, filename string, want int) {
	f, err := os.Open(filename)
	defer f.Close()
	require.NoError(t, err)

	answer, err := puzzle(f)

	require.NoError(t, err)
	require.Equal(t, want, answer)

	fmt.Printf("Answer = %v\n", answer)
}
