package flashcards

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// CreateCards creates the flash cards
func CreateCards(fileName string, skipHeader bool) Cards {
	// holds all the cards
	var cards Cards

	// tempt variable to hold a single card
	var tempt Card

	//count allows me to skip the first entry
	var count int

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
		
		if skipHeader && count == 0{
			
		}else {
			// Tempt variable for card
			tempt = Card{record[0], record[1], record[2]}

			// Append to list
			cards.Cards = append(cards.Cards, tempt)
		}
		count++
	}
	return cards
}
