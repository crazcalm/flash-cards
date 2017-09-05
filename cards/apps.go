package flashcards

import (
	"bufio"
	"fmt"
	"os"
	"text/template"
)

//FlashcardApp is used to run the terminal flashcard app
func FlashcardApp(csvFile string) {
	//Clear the screen
	Clear()

	//CSV file that was passed in
	//flag.Parse()

	// holds all the cards
	cards := CreateCards(csvFile)
	cards.Shuffle()

	input := bufio.NewScanner(os.Stdin)
	var userInput string
	var count int
	var cardFace string
	var output *template.Template

	//Need to print the first card...
	cardFace = InputCardFace(userInput)
	output = TemplateString(cards.Cards[count], cardFace)
	PrintToScreen(output, cards.Cards[count])
	fmt.Printf(COUNTERTEXT, count+1, len(cards.Cards))
	fmt.Printf(HELPTEXT)
	fmt.Printf(USERINPUTTEXT)

	// Testing out the user interface loop
	for input.Scan() {
		userInput = input.Text()

		// clears the screen
		Clear()

		if BreakLoop(userInput) {
			break
		}
		count = CardSelectCounter(userInput, count)
		//Break if out of range
		if len(cards.Cards) <= count {
			break
		}
		cardFace = InputCardFace(userInput)
		output = TemplateString(cards.Cards[count], cardFace)
		PrintToScreen(output, cards.Cards[count])
		fmt.Printf(COUNTERTEXT, count+1, len(cards.Cards))
		fmt.Printf(HELPTEXT)
		fmt.Printf(USERINPUTTEXT)
	}

	fmt.Println("The program completed running")

}
