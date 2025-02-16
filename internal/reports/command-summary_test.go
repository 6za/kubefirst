package reports

import (
	"strings"
	"testing"
)

// TestStyleMessage test if a regular string get transformed into a new styled string and keep its original content
// not changed
func TestStyleMessage(t *testing.T) {

	message := "kubefirst rock!"
	got := StyleMessage(message)
	wanted := strings.Contains(got, message)

	if wanted != true {
		t.Error("returned string doesnt contain the content of the initial message")
	}

}
