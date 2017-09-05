package flashcards

import (
	"strings"
)

// BreakLoop handles the cases that break the loop
func BreakLoop(input string) bool {
	quits := []string{"q", "quit"}
	var answer bool
	for _, quit := range quits {
		if strings.Compare(input, quit) == 0 {
			answer = true
		}
	}
	return answer
}

//InputCardFace used to figure out what part of the card the user wants to see
func InputCardFace(input string) string {
	flips := []string{"f", "flip"}
	hints := []string{"h", "hint"}
	result := "front"

	if InSlice(flips, input) {
		result = "flip"
	} else if InSlice(hints, input) {
		result = "hint"
	}
	return result
}

//CardSelectCounter use to increment and decrement the card counter
func CardSelectCounter(input string, count int) int {
	forward := []string{"n", "next"}
	previous := []string{"p", "previous"}

	for _, f := range forward {
		if strings.Compare(f, input) == 0 {
			count++
		}
	}

	for _, p := range previous {
		if strings.Compare(p, input) == 0 {
			if count > 0 {
				count--
			}
		}
	}
	return count
}
