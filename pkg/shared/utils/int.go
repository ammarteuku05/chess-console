package utils

import (
	"errors"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func ParseInput(input string) (sr, sc, er, ec int, err error) {
	parts := strings.Fields(strings.ToLower(input))
	if len(parts) != 2 {
		return 0, 0, 0, 0, errors.New("input must be two coordinates (e.g., a2 b3)")
	}

	sr, sc, err = parseCoord(parts[0])
	if err != nil {
		return
	}

	er, ec, err = parseCoord(parts[1])
	return
}

func parseCoord(coord string) (int, int, error) {
	if len(coord) != 2 {
		return 0, 0, errors.New("invalid coordinate format")
	}

	col := int(coord[0] - 'a')
	row := 8 - int(coord[1]-'0')

	if row < 0 || row > 7 || col < 0 || col > 7 {
		return 0, 0, errors.New("coordinate out of bounds")
	}

	return row, col, nil
}
