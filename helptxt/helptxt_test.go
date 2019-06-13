package helptxt

import (
	"strings"
	"testing"
)

func TestSetHelpTxt(t *testing.T) {
	inp := []string{"this", "is", "test"}
	out := strings.Join(inp, "\n")
	if out != SetHelpTxt(inp) {
		t.Errorf("expected values is : %s\n received value : %s ",
			out, SetHelpTxt(inp))
	}

}
