package flashcards

import (
	"testing"
)

func TestBreakLoop(t *testing.T) {
	var tests = []struct {
		Input  string
		Answer bool
	}{
		{"q", true},
		{"quit", true},
		{"", false},
		{"no", false},
	}

	for _, test := range tests {
		result := BreakLoop(test.Input)
		if result != test.Answer {
			t.Fail()
		}
	}
}

func TestInputCardFace(t *testing.T) {
	var tests = []struct {
		Input  string
		Answer string
	}{
		{"f", "flip"},
		{"flip", "flip"},
		{"h", "hint"},
		{"hint", "hint"},
		{"", "front"},
		{"hello", "front"},
	}

	for _, test := range tests {
		result := InputCardFace(test.Input)
		if result != test.Answer {
			t.Fail()
		}
	}
}

func TestCardSelectCounter(t *testing.T) {
	var tests = []struct {
		Input   string
		Counter int
		Answer  int
	}{
		{"n", 0, 1},
		{"next", 0, 1},
		{"p", 1, 0},
		{"previous", 1, 0},
		{"p", 0, 0},
		{"", 1, 1},
		{"hi", 1, 1},
	}

	for _, test := range tests {
		result := CardSelectCounter(test.Input, test.Counter)

		if result != test.Answer {
			t.Fail()
		}
	}
}
