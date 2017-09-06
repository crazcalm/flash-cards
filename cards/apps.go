package flashcards

import (
	"bufio"
	"fmt"
	"os"
	"text/template"
	"math"
)

//GroupCardsApp Breaks the cards into groups
func GroupCardsApp(csvFile string, numOfGroups int) {
	//Clear the screen
	Clear()

	//Holds all the cards
	cards := CreateCards(csvFile, true)

	//Shuffles the cards
	cards.Shuffle()

	numOfPeopleInGroups := math.Ceil(float64(len(cards.Cards))/float64(numOfGroups))

	//count is used for the internal loop
	var count int

	for i:= 1; i<=numOfGroups; i++ {
		fmt.Println("Group %f:", i)
		for count<len(cards.Cards) {
			output := CreateTemplate("test7", GROUPCARD)
			fmt.Print(count, " ")
			PrintToScreen(output, cards.Cards[count])
			count++
			if count > 0 && math.Mod(float64(count), float64(numOfPeopleInGroups)) == 0{
				break
			} 
		}
	}
	
}


//RandomCardApp Prints out one random card
func RandomCardApp(csvFile string) {
	//Clear the screen
	Clear()

	//Holds all the cards
	cards := CreateCards(csvFile, false)

	//Shuffles the cards
	cards.Shuffle()

	template := CreateTemplate("test6", RANDOMCARD)
	PrintToScreen(template, cards.Cards[0])
}

//FlashcardApp is used to run the terminal flashcard app
func FlashcardApp(csvFile string) {
	//Clear the screen
	Clear()

	// holds all the cards
	cards := CreateCards(csvFile, false)
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
