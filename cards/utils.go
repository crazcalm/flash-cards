package flashcards

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

//PrintToScreen prints templates to standard out
func PrintToScreen(templ *template.Template, data interface{}, w io.Writer) {
	err := templ.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

//InSlice used to check of string is in slice
func InSlice(slice []string, s string) bool {
	result := false
	for _, item := range slice {
		if strings.Compare(item, s) == 0 {
			result = true
		}
	}
	return result
}

// Clear clears the screen.
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
