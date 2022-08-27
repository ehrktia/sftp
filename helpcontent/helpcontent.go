package helpcontent

import (
	"io"
	"log"
	"os"
)

func Helpcontent() string {
	var filname = "./helpcontent/helpcontent"
	f, err := os.Open(filname)
	if err != nil {
		panic("missing help text file")
	}
	d, err := io.ReadAll(f)
	if err != nil && err != io.EOF {
		log.Printf("%s-%s","ERROR",err)
	}
	return string(d)
}
