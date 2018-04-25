package flashcards

import (
	"log"
	"strings"
	"text/template"
)

//createTemplate is responsible for creating the templates
func createTemplate(name, words string) *template.Template {
	templ, err := template.New(name).Parse(words)
	if err != nil {
		log.Fatal(err)
	}
	return templ
}

//TemplateString used to create the correct card template output
func TemplateString(c *Card, face string) *template.Template {
	var result *template.Template
	if strings.Compare(face, "hint") == 0 {
		result = createTemplate("test3", CARDHINT)
	} else if strings.Compare(face, "flip") == 0 {

		//Checks the flipped state and flips accordingly
		if c.Flipped {
			result = createTemplate("test5", CARDFRONT)
		} else {
			result = createTemplate("test4", CARDBACK)
		}
		//Change the flip state
		c.Flip()
	} else {
		result = createTemplate("test5", CARDFRONT)
	}
	return result
}
