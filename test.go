package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
	"github.com/crazcalm/flash-cards/cards"
)

var csvFile = flag.String("f", "", "file: path to csv file")

// BreakLoop handles the cases that break the loop
func BreakLoop(input string) bool {
	quits := []string{"q", "quit"}
	var answer bool
	for _, quit := range quits {
		if strings.Compare(input, quit) == 0 {
			answer = true
		}
	}
	return answer
}

//InputCardFace used to figure out what part of the card the user wants to see
func InputCardFace(input string) string {
	flips := []string{"f", "flip"}
	hints := []string{"h", "hint"}
	result := "front"

	if flashcards.InSlice(flips, input) {
		result = "flip"
	} else if flashcards.InSlice(hints, input) {
		result = "hint"
	}
	return result
}

//CardSelectCounter use to increment and decrement the card counter
func CardSelectCounter(input string, count int) int {
	forward := []string{"n", "next"}
	previous := []string{"p", "previous"}

	for _, f := range forward {
		if strings.Compare(f, input) == 0 {
			count++
		}
	}

	for _, p := range previous {
		if strings.Compare(p, input) == 0 {
			if count > 0 {
				count--
			}
		}
	}
	return count
}

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
	cardFace = InputCardFace(userInput)
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

		if BreakLoop(userInput) {
			break
		}
		count = CardSelectCounter(userInput, count)
		//Break if out of range
		if len(cards.Cards) <= count {
			break
		}
		cardFace = InputCardFace(userInput)
		output = flashcards.TemplateString(cards.Cards[count], cardFace)
		flashcards.PrintToScreen(output, cards.Cards[count])
		fmt.Printf(flashcards.COUNTERTEXT, count + 1, len(cards.Cards))
		fmt.Printf(flashcards.HELPTEXT)
		fmt.Printf(flashcards.USERINPUTTEXT)
	}

	fmt.Println("The program completed running")

}
