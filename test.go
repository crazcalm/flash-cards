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


func main(){
	// holds all the cards
	var cards Cards

	// tempt variable to hold a single card
	var tempt Card

	// opens the test csv file
	file, err := os.Open("test_data.csv")
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

		if err != nil{
			log.Fatal(err)
		}
		
		// Tempt variable
		tempt = Card{record[0], record[1], record[2]}

		// Fill up the Cards
		cards.Cards = append(cards.Cards, tempt)

		// Testing out the template stuff
		templ, err := template.New("test").Parse("\n--------------\nFront: {{.Front}}\nBack: {{.Back}}\nHint: {{.Hint}}\n")
		if err != nil {
			log.Fatal(err)
		}
		err = templ.Execute(os.Stdout, tempt)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Testing out using range in templates
	templ2, err := template.New("test2").Parse("{{range .Cards}}------------\nFront: {{.Front}}\n{{end}}")
	if err != nil {
		log.Fatal(err)
	}
	err = templ2.Execute(os.Stdout, cards)
	if err != nil {
		log.Fatal(err)
	}


	input := bufio.NewScanner(os.Stdin)
	var userInput string
	var count int

	// Testing out the user interface loop
	for input.Scan() {
		userInput = input.Text()
		fmt.Println(userInput)
		fmt.Println(cards.Cards[count])

		if strings.Compare(userInput, "q") == 0 {
			break
		}
		count++
	}

	fmt.Println("The program completed running")

}