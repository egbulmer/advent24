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
	puzzle := newPuzzle(string(dat), modeLetters)
	return puzzle.solve(), nil
}

func Puzzle2(r io.Reader) (int, error) {
	dat, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	puzzle := newPuzzle(string(dat), modeCross)
	return puzzle.solve(), nil
}

var ErrDNE = errors.New("does not exist")
var xmas = []byte("XMAS")

type mode int

const (
	modeLetters mode = iota
	modeCross
)

type puzzle struct {
	data []byte
	w, h int
	mode mode
}

type vec2 struct {
	x, y int
}

func (v1 vec2) plus(v2 vec2) vec2 {
	return vec2{v1.x + v2.x, v1.y + v2.y}
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

var corners = []vec2{
	directions[NW], directions[NE], directions[SW], directions[SE],
}

func newPuzzle(input string, mode mode) puzzle {
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

	return puzzle{data, w, h, mode}
}

func (p puzzle) charAt(pos vec2) (byte, error) {
	if pos.x < 0 || pos.x >= p.w || pos.y < 0 || pos.y >= p.h {
		return 0, ErrDNE
	}
	return p.data[pos.y*p.w+pos.x], nil
}

func (p puzzle) scanForLetters(pos vec2) int {
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

		pos = pos.plus(dirn)
	}
	return true
}

func (p puzzle) scanForCrosses(pos vec2) int {
	ch, _ := p.charAt(pos)
	if ch != 'A' {
		return 0
	}

	corners, err := p.cornersAt(pos)
	if err == ErrDNE {
		return 0
	}

	tl, tr, bl, br := corners[0], corners[1], corners[2], corners[3]
	if isMAS(tl, tr, bl, br) {
		return 1
	}

	return 0
}

func (p puzzle) cornersAt(center vec2) ([4]byte, error) {
	var result [4]byte

	for i, dirn := range corners {
		pos := center.plus(dirn)
		ch, err := p.charAt(pos)
		if err != nil {
			return result, err
		}
		result[i] = ch
	}

	return result, nil
}

func (p puzzle) solve() int {
	count := 0
	for y := 0; y < p.h; y++ {
		for x := 0; x < p.w; x++ {
			pos := vec2{x, y}
			switch p.mode {
			case modeLetters:
				count += p.scanForLetters(pos)
			case modeCross:
				count += p.scanForCrosses(pos)
			}
		}
	}
	return count
}

func isMS(a, b byte) bool {
	return (a == 'M' && b == 'S') || (a == 'S' && b == 'M')
}

func isMAS(tl, tr, bl, br byte) bool {
	return isMS(tl, br) && isMS(tr, bl)
}
