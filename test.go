package main

import (
	"flag"
	"github.com/crazcalm/flash-cards/cards"
)

var csvFile = flag.String("f", "", "file: path to csv file")

func main() {
	//CSV file that was passed in
	flag.Parse()

	//Run the app
	flashcards.RandomCardApp(*csvFile)
}
