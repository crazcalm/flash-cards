package flashcards

import (
	"bufio"
	"fmt"
	"os"
	"text/template"
)

//FlashcardApp is used to run the terminal flashcard app
func FlashcardApp(cards FlashCards, shuffle bool) {
	//Clear the screen
	Clear()

	//Shuffles the cards
	if shuffle {
		cards.Shuffle()
	}

	input := bufio.NewScanner(os.Stdin)
	var userInput string
	var count int
	var cardFace string
	var output *template.Template

	//Need to print the first card...
	cardFace = InputCardFace(userInput)
	output = TemplateString(&cards.GetCards()[count], cardFace)
	PrintToScreen(output, cards.GetCards()[count], os.Stdout)
	fmt.Printf(COUNTERTEXT, count+1, len(cards.GetCards()))
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
		if len(cards.GetCards()) <= count {
			break
		}
		cardFace = InputCardFace(userInput)
		output = TemplateString(&cards.GetCards()[count], cardFace)
		PrintToScreen(output, cards.GetCards()[count], os.Stdout)
		fmt.Printf(COUNTERTEXT, count+1, len(cards.GetCards()))
		fmt.Printf(HELPTEXT)
		fmt.Printf(USERINPUTTEXT)
	}

	fmt.Println("The program completed running")

}
