package main

import (
	"flag"
	"github.com/crazcalm/flash-cards/cards"
)

var csvFile = flag.String("f", "", "file: path to csv file")
var numOfGroups = flag.Int("g", 0, "number of groups")

func main() {
	//CSV file that was passed in
	flag.Parse()

	//Run the app
	flashcards.GroupCardsApp(*csvFile, *numOfGroups)
}
