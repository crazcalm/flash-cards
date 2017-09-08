package flashcards

import (
	"bytes"
	"testing"
	"text/template"
	"strings"
)

func TestPrintToScreen(t *testing.T){
	// testing template
	templ, err := template.New("testing1").Parse("{{.Front}} {{.Back}} {{.Hint}}")
	if err != nil {
		t.Errorf("PrintToScreen failed when creating the template: %s\n", err)
	}

	// Test cases
	var tests = []struct{
		card		Card
		expected	string
	}{
		{
			Card{"Front", "Back", "Hint"},
			"Front Back Hint",
		},
		{
			Card{"NANA", "NA", "Batman"},
			"NANA NA Batman",
		},
	}

	b := new(bytes.Buffer)

	for _, test := range tests{
		PrintToScreen(templ, test.card, b)

		got := b.String()

		if strings.Compare(got, test.expected) != 0 {
			t.Errorf("PrintToScreen: %s != %s", test.expected, got)
		}

		b.Reset()
	}
}