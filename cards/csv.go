package flashcards

import (
	"github.com/artonge/go-csv-tag"
)

//CreateCards created the flashcards
func CreateCards(fileName string, dest *[]Card){
	csvtag.Load(csvtag.Config{
		Path: fileName,
		Dest: dest,
	})
}
