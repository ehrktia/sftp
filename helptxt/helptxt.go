package helptxt

import (
	"strings"
)

/* set the vals for flags using func one to one */
func SetHelpTxt(helpsl []string) string {
	outtxt := strings.Join(helpsl, "\n")
	return outtxt
}
