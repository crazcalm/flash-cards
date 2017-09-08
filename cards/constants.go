package flashcards

const (
	//CARDFRONT template for the front of the card
	CARDFRONT = "Card front:\n{{.Front}}\n\n\n"
	//CARDBACK template for the back of the card
	CARDBACK = "Card back:\n{{.Back}}\n\n\n"
	//CARDHINT template for the card hint
	CARDHINT = "Card hint:\n{{.Hint}}\n\n\n"
	//HELPTEXT shows the commands
	HELPTEXT = "\n\n(n)ext (p)revious (f)lip (h)int (q)uit\n"
	//COUNTERTEXT template for card counter text
	COUNTERTEXT = "Count: %d/%d\n"
	//USERINPUTTEXT the template for taking in the user input
	USERINPUTTEXT = "\nUser input: "
	//RANDOMCARD used for the random card app
	RANDOMCARD = "Name: {{.Front}} ({{.Back}}) -- {{.Hint}}\n"
	//GROUPCARD used for the group card app
	GROUPCARD = "{{.Front}} ({{.Back}}) -- {{.Hint}}\n"
)
