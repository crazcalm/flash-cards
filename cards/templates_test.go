package flashcards

import (
	"bytes"
	"strings"
	"testing"
)

func TestTemplateString(t *testing.T) {
	// Test card
	card := Card{"FRONT!!!", "BACK!!!", "HINT!!!"}

	// test cases
	var tests = []struct {
		input    string
		expected []string
	}{
		{
			"hint",
			[]string{
				"Card hint:",
				"HINT!!!",
			},
		},
		{
			"flip",
			[]string{
				"Card back:",
				"BACK!!!",
			},
		},
		{
			"default",
			[]string{
				"Card front:",
				"FRONT!!!",
			},
		},
		{
			"",
			[]string{
				"Card front:",
				"FRONT!!!",
			},
		},
	}

	// buffer used to check the result
	b := new(bytes.Buffer)

	for _, test := range tests {
		templ := TemplateString(card, test.input)

		err := templ.Execute(b, card)
		if err != nil {
			t.Errorf("TemplateString failed when executing the template: %s", err)
		}
		got := b.String()

		// Checks that the output has all the required parts
		for _, item := range test.expected {
			if strings.Contains(got, item) != true {
				t.Errorf("TemplateString: (%s) does not contain (%s)", got, item)
			}
		}

		// Clears the buffer
		b.Reset()
	}
}
