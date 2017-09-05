package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"text/template"
	"github.com/crazcalm/flash-cards/cards"
)

var csvFile = flag.String("f", "", "file: path to csv file")

func main() {
	//Clear the screen
	flashcards.Clear()

	//CSV file that was passed in
	flag.Parse()

	// holds all the cards
	cards := flashcards.CreateCards(*csvFile)
	cards.Shuffle()

	input := bufio.NewScanner(os.Stdin)
	var userInput string
	var count int
	var cardFace string
	var output *template.Template

	//Need to print the first card...
	cardFace = flashcards.InputCardFace(userInput)
	output = flashcards.TemplateString(cards.Cards[count], cardFace)
	flashcards.PrintToScreen(output, cards.Cards[count])
	fmt.Printf(flashcards.COUNTERTEXT, count + 1, len(cards.Cards))
	fmt.Printf(flashcards.HELPTEXT)
	fmt.Printf(flashcards.USERINPUTTEXT)

	// Testing out the user interface loop
	for input.Scan() {
		userInput = input.Text()

		// clears the screen
		flashcards.Clear()

		if flashcards.BreakLoop(userInput) {
			break
		}
		count = flashcards.CardSelectCounter(userInput, count)
		//Break if out of range
		if len(cards.Cards) <= count {
			break
		}
		cardFace = flashcards.InputCardFace(userInput)
		output = flashcards.TemplateString(cards.Cards[count], cardFace)
		flashcards.PrintToScreen(output, cards.Cards[count])
		fmt.Printf(flashcards.COUNTERTEXT, count + 1, len(cards.Cards))
		fmt.Printf(flashcards.HELPTEXT)
		fmt.Printf(flashcards.USERINPUTTEXT)
	}

	fmt.Println("The program completed running")

}
