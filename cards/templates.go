package flashcards

import (
	"log"
	"strings"
	"text/template"
)

//CreateTemplate is responsible for creating the templates
func CreateTemplate(name, words string) *template.Template {
	templ, err := template.New(name).Parse(words)
	if err != nil {
		log.Fatal(err)
	}
	return templ
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
