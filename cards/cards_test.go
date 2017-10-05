package flashcards

import (
	"path/filepath"
	"testing"
)

func TestGetCards(t *testing.T) {
	var cards Cards

	err := CreateCards(filepath.Join("test_files", "test_data.csv"), &cards.Cards)
	if err != nil {
		t.Fail()
	}

	if len(cards.GetCards()) == 0 {
		t.Fail()
	}
}

func TestShuffle(t *testing.T) {
	var cards Cards

	err := CreateCards(filepath.Join("test_files", "test_data.csv"), &cards.Cards)
	if err != nil {
		t.Fail()
	}

	//Checking two cards to lower the chance of false positives
	c1 := cards.GetCards()[0]
	c2 := cards.GetCards()[1]

	cards.Shuffle()

	if cards.GetCards()[0] == c1 && cards.GetCards()[1] == c2 {
		t.Fail()
	}
}
