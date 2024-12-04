package day4

import (
	"errors"
	"io"
	"strings"
)

func Puzzle1(r io.Reader) (int, error) {
	dat, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	puzzle := newPuzzle(string(dat))
	return puzzle.solve(), nil
}

func Puzzle2(r io.Reader) (int, error) {
	return 0, errors.New("unsolved")
}

type puzzle struct {
	data []byte
	w, h int
}

type vec2 struct {
	x, y int
}

const (
	N int = iota
	S
	W
	E
	NE
	NW
	SE
	SW
)

var ErrDNE = errors.New("does not exist")

var xmas = []byte("XMAS")

var directions = []vec2{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
	{1, -1},
	{-1, -1},
	{1, 1},
	{-1, 1},
}

func newPuzzle(input string) puzzle {
	lines := strings.Split(input, "\n")
	w, h := 0, 0
	data := make([]byte, 0, 140*140)

	for _, line := range lines {
		chars := []byte(strings.TrimSpace(line))
		if len(chars) == 0 {
			continue
		}
		w = len(chars)
		for j := 0; j < w; j++ {
			data = append(data, chars[j])
		}
		h += 1
	}

	return puzzle{data, w, h}
}

func (p puzzle) charAt(pos vec2) (byte, error) {
	if pos.x < 0 || pos.x >= p.w || pos.y < 0 || pos.y >= p.h {
		return 0, ErrDNE
	}
	return p.data[pos.x*p.w+pos.y], nil
}

func (p puzzle) scan(pos vec2) int {
	count := 0
	for _, dirn := range directions {
		if p.findXmas(pos, dirn) {
			count += 1
		}
	}
	return count
}

func (p puzzle) findXmas(pos vec2, dirn vec2) bool {
	for i := 0; i < 4; i++ {
		ch, err := p.charAt(pos)
		if err == ErrDNE {
			return false
		}

		if ch != xmas[i] {
			return false
		}

		pos.x = pos.x + dirn.x
		pos.y = pos.y + dirn.y
	}
	return true
}

func (p puzzle) solve() int {
	count := 0
	for y := 0; y < p.h; y++ {
		for x := 0; x < p.w; x++ {
			pos := vec2{x, y}
			count += p.scan(pos)
		}
	}
	return count
}
