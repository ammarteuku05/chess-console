package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive", 5, 5},
		{"negative", -5, 5},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Abs(tt.input))
		})
	}
}

func TestSign(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive", 10, 1},
		{"negative", -10, -1},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Sign(tt.input))
		})
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  []int
		expectErr bool
	}{
		{"valid move a2 a3", "a2 a3", []int{6, 0, 5, 0}, false},
		{"valid move b8 c6", "b8 c6", []int{0, 1, 2, 2}, false},
		{"invalid format - too many parts", "a2 a3 a4", nil, true},
		{"invalid format - one part", "a2", nil, true},
		{"invalid coordinate - out of bounds", "z9 a2", nil, true},
		{"invalid coordinate - length", "a 2", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr, sc, er, ec, err := ParseInput(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected[0], sr)
				assert.Equal(t, tt.expected[1], sc)
				assert.Equal(t, tt.expected[2], er)
				assert.Equal(t, tt.expected[3], ec)
			}
		})
	}
}

func TestParseCoord(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectedR int
		expectedC int
		expectErr bool
	}{
		{"valid a1", "a1", 7, 0, false},
		{"valid h8", "h8", 0, 7, false},
		{"invalid length", "a11", 0, 0, true},
		{"out of bounds row high", "a9", 0, 0, true},
		{"out of bounds row low", "a0", 0, 0, true},
		{"out of bounds col", "z1", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, c, err := parseCoord(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedR, r)
				assert.Equal(t, tt.expectedC, c)
			}
		})
	}
}
