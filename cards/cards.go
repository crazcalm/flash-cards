package flashcards

import (
	"math/rand"
	"time"
)

// Cards is a data structure to hold all of the flash cards
type Cards struct {
	Cards [] Card
}


func init(){
	// Seedning the random generator
	rand.Seed(time.Now().UnixNano())
}

// Shuffle shuffles the cards in place
func (c Cards) Shuffle(){
	numOfCards := len(c.Cards)
	var tempt Card
	var swapIndex int

	if numOfCards > 0 {
		numOfCards--
	}

	for index := range c.Cards {
		swapIndex = rand.Intn(numOfCards)
		tempt = c.Cards[index]
		c.Cards[index] = c.Cards[swapIndex]
		c.Cards[swapIndex] = tempt 
	}
}