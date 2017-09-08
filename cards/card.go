package flashcards

//Card is a struct to represent a single flash card
type Card struct {
	Front string `csv:"front"`
	Back  string `csv:"back"`
	Hint  string `csv:"hint"`
}
