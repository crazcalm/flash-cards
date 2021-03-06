package flashcards

import (
	"bytes"
	"strings"
	"testing"
	"text/template"
)

func TestInSlice(t *testing.T) {
	//Test cases
	var tests = []struct {
		List   []string
		Item   string
		Answer bool
	}{
		{[]string{"yes", "no", "maybe"}, "yes", true},
		{[]string{"yes", "no", "maybe"}, "idk", false},
		{[]string{}, "", false},
	}

	for _, test := range tests {
		result := InSlice(test.List, test.Item)
		if result != test.Answer {
			t.Fail()
		}
	}
}

func TestPrintToScreen(t *testing.T) {
	// testing template
	templ, err := template.New("testing1").Parse("{{.Front}} {{.Back}} {{.Hint}}")
	if err != nil {
		t.Errorf("PrintToScreen failed when creating the template: %s\n", err)
	}

	// Test cases
	var tests = []struct {
		card     Card
		expected string
	}{
		{
			Card{"Front", "Back", "Hint", false},
			"Front Back Hint",
		},
		{
			Card{"NANA", "NA", "Batman", false},
			"NANA NA Batman",
		},
	}

	b := new(bytes.Buffer)

	for _, test := range tests {
		PrintToScreen(templ, test.card, b)

		got := b.String()

		if strings.Compare(got, test.expected) != 0 {
			t.Errorf("PrintToScreen: %s != %s", test.expected, got)
		}

		b.Reset()
	}
}
