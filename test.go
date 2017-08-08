package main

import (
	"os"
	"log"
	"encoding/csv"
	"io"
	"text/template"
	"fmt"
	"bufio"
	"strings"
)

const (
	//TestCSV test csv
	TestCSV = "test_data.csv"
	//CARDFRONT template for the front of the card
	CARDFRONT = "Card front:\n{{.Front}}\n"
	//CARDBACK template for the back of the card
	CARDBACK = "Card back:\n{{.Back}}\n"
	//CARDHINT template for the card hint
	CARDHINT = "Card Hint:\n{{.Hint}}\n"
)

//Card is template for a flash card
type Card struct {
	Front		string
	Back		string
	Hint		string
}

//Cards struct for all cards
type Cards struct {
	Cards		[]Card
}


//CreateTemplate is responsible for creating the templates
func CreateTemplate(name, words string) *template.Template {
	templ, err := template.New(name).Parse(words)
	if err != nil {
		log.Fatal(err)
	}
	return templ
}

//PrintToScreen prints templates to standard out
func PrintToScreen(templ *template.Template, data interface{}){
	err := templ.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}


// CreateCards creates the flash cards
func CreateCards(fileName string) Cards {
	// holds all the cards
	var cards Cards

	// tempt variable to hold a single card
	var tempt Card

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// reading the CSV file
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Tempt variable for card
		tempt = Card{record[0], record[1], record[2]}

		// Append to list
		cards.Cards = append(cards.Cards, tempt)
	}
	return cards
}

// BreakLoop handles the cases that break the loop
func BreakLoop(input string) bool{
	quits := []string{"q", "quit" }
	var answer bool
	for _, quit := range quits {
		if strings.Compare(input, quit) == 0 {
			answer = true
		}
	}
	return answer
}

//InSlice used to check of string is in slice
func InSlice(slice []string, s string) bool{
	result := false
	for _, item := range slice {
		if strings.Compare(item, s) == 0 {
			result = true
		}
	}
	return result
}

//InputCardFace used to figure out what part of the card the user wants to see
func InputCardFace(input string) string {
	flips := []string{"f", "flip"}
	hints := []string{"h", "hint"}
	result := "front"

	if InSlice(flips, input){
		result = "flip"
	} else if InSlice(hints, input){
		result = "hint"
	}
	return result
}

//TemplateString used to create the correct card template output
func TemplateString(c Card, face string) *template.Template {
	var result *template.Template
	if strings.Compare(face, "hint") == 0 {
		result = CreateTemplate("test3", CARDHINT)
	}else if strings.Compare(face, "flip") == 0 {
		result = CreateTemplate("test4", CARDBACK)
	}else {
		result = CreateTemplate("test5", CARDFRONT)
	}
	return result
}

func main(){
	// holds all the cards
	cards := CreateCards(TestCSV)

	// Testing out using range in templates
	templ2 := CreateTemplate("test2", "{{range .Cards}}------------\nFront: {{.Front}}\n{{end}}")
	PrintToScreen(templ2, cards)

	input := bufio.NewScanner(os.Stdin)
	var userInput string
	var count int
	var cardFace string
	var output *template.Template

	// Testing out the user interface loop
	for input.Scan() {
		userInput = input.Text()
		fmt.Println(userInput)
		fmt.Printf("Count: %d\n", count)
		fmt.Println(cards.Cards[count])

		cardFace = InputCardFace(userInput)
		output = TemplateString(cards.Cards[count], cardFace)
		PrintToScreen(output, cards.Cards[count])

		if BreakLoop(userInput) {
			break
		}else {
			count++
		}
		//Break if out of range
		if len(cards.Cards) <= count {
			break
		}
	}

	fmt.Println("The program completed running")

}