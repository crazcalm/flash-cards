package main

import (
	"flag"
	"github.com/crazcalm/flash-cards/cards"
)

var csvFile = flag.String("f", "", "file: path to csv file")
var shuffle = flag.Bool("s", true, "Shuffle the cards")

func main() {
	//CSV file that was passed in
	flag.Parse()

	//cards will hold the cards
	var cards flashcards.Cards

	flashcards.CreateCards(*csvFile, &cards.Cards)
	cards.Shuffle()

	//Run the app
	flashcards.FlashcardApp(cards, *shuffle)
}
