package flashcards

//Card is a struct to represent a single flash card
type Card struct {
	Front 	string `csv:"front"`
	Back  	string `csv:"back"`
	Hint  	string `csv:"hint"`
	Flipped	bool
}

//Flip tracks the flipped state of the card
func (c *Card) Flip(){
	if c.Flipped {
		c.Flipped = false
	}else{
		c.Flipped = true
	}
}
