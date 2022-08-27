package helpcontent

import (
	"io"
	"log"
)

func Helpcontent() string {
	var filname = "helpcontent"
	content, err := io.ReadAll(filname)
	if err != nil {
		log.Fatal("ERROR: Reading helpcontent file")
	}
	return string(content)

}
