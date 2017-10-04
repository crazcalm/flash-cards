package flashcards

import (
	"testing"
)

func TestFlip(t *testing.T) {
	var tests = []struct {
		Card   Card
		Answer bool
	}{
		{
			Card{"Front", "Back", "Hint", true},
			false,
		},
		{
			Card{"Front", "Back", "Hint", false},
			true,
		},
	}

	for _, test := range tests {
		test.Card.Flip()

		if test.Card.Flipped != test.Answer {
			t.Errorf("TestFlip: %v != %v", test.Card.Flipped, test.Answer)
		}
	}
}
