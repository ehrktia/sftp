package helpcontent

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

func TestHelpcontent(t *testing.T) {
	filname := filepath.Join("testdata", "helpcontent")
	inp, err := ioutil.ReadFile(filname)
	if err != nil {
		log.Fatalf("error occured %s", err)
	}
	in := string(inp)
	actual := Helpcontent()
	if in != actual {
		t.Errorf("exp out :%s actual out:%s", inp, actual)
	}
}
