package helpcontent

/* declare inps and out vars */
var helpsl = []string{"this is a command line operated program",
	"To get list of available option  provide the command option -help=help",
	"ex: uname=TE123455 pwd=password",
	"To generate key use key mode"}

/* all operations performed here 1:1 */
func Helpcontent() []string {
	return helpsl
}
