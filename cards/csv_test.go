package flashcards

import (
	"path/filepath"
	"testing"
)

func TestCreateCards(t *testing.T) {
	var cards Cards

	err := CreateCards(filepath.Join("test_files", "test_data.csv"), &cards.Cards)
	if err != nil {
		t.Fail()
	}

	if len(cards.Cards) == 0 {
		t.Fail()
	}
}
