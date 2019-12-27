package helpcontent

import (
	"io/ioutil"
	"log"
)

func Helpcontent() string {
	var filname = "helpcontent"
	content, err := ioutil.ReadFile(filname)
	if err != nil {
		log.Fatal("ERROR: Reading helpcontent file")
	}
	return string(content)
    
}
