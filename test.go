package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

const (
	//CARDFRONT template for the front of the card
	CARDFRONT = "Card front:\n{{.Front}}\n\n\n"
	//CARDBACK template for the back of the card
	CARDBACK = "Card back:\n{{.Back}}\n\n\n"
	//CARDHINT template for the card hint
	CARDHINT = "Card Hint:\n{{.Hint}}\n\n\n"
	//HELPTEXT shows the commands
	HELPTEXT = "\n\n(n)ext (p)revious (f)lip (h)int (q)uit\n"
	//COUNTERTEXT template for card counter text
	COUNTERTEXT = "Count: %d/%d\n"
	//USERINPUTTEXT the template for taking in the user input
	USERINPUTTEXT = "\nUser input: "
)

var csvFile = flag.String("f", "", "file: path to csv file")

func init() {
	//I need to double check that is the proper way to seed the rand generator.
	rand.Seed(time.Now().UnixNano())
}

//Card is template for a flash card
type Card struct {
	Front string
	Back  string
	Hint  string
}

//Cards struct for all cards
type Cards struct {
	Cards []Card
}

//Shuffle shuffles the cards
func (c Cards) Shuffle() {
	numOfCards := len(c.Cards)
	var tempt Card
	var swapIndex int

	// I need to figure out the cases that I will ignore...
	// I also need to figure out the each cases...
	if numOfCards > 1 {
		numOfCards--
	}

	for index := range c.Cards {
		swapIndex = rand.Intn(numOfCards)
		tempt = c.Cards[index]
		c.Cards[index] = c.Cards[swapIndex]
		c.Cards[swapIndex] = tempt
	}
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
func PrintToScreen(templ *template.Template, data interface{}) {
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

//InSlice used to check of string is in slice
func InSlice(slice []string, s string) bool {
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

	if InSlice(flips, input) {
		result = "flip"
	} else if InSlice(hints, input) {
		result = "hint"
	}
	return result
}

//TemplateString used to create the correct card template output
func TemplateString(c Card, face string) *template.Template {
	var result *template.Template
	if strings.Compare(face, "hint") == 0 {
		result = CreateTemplate("test3", CARDHINT)
	} else if strings.Compare(face, "flip") == 0 {
		result = CreateTemplate("test4", CARDBACK)
	} else {
		result = CreateTemplate("test5", CARDFRONT)
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
			count--
		}
	}
	return count
}

// Clear clears the screen.
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	//Clear the screen
	Clear()

	//CSV file that was passed in
	flag.Parse()

	// holds all the cards
	cards := CreateCards(*csvFile)
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
	fmt.Printf(COUNTERTEXT, count + 1, len(cards.Cards))
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
		fmt.Printf(COUNTERTEXT, count + 1, len(cards.Cards))
		fmt.Printf(HELPTEXT)
		fmt.Printf(USERINPUTTEXT)
	}

	fmt.Println("The program completed running")

}
