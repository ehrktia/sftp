package helpcontent

import (
	"log"
	"os"
	"testing"
)

func TestHelpcontent(t *testing.T) {
	log := log.New(os.Stderr, "Msg:", log.Ldate|log.Ltime|log.Lshortfile)
	t.Run("open file", func(t *testing.T) {
		got, err := getFile(log)
		if err != nil {
			t.Fatal(err)
		}
		if got == nil {
			t.Fatal("error reading file")
		}
	})
	t.Run("read from helpcntnt file",
		func(t *testing.T) {
			got := Helpcontent(log)
			if got == "" {
				t.Fatal("error reading file")
			}
		})
}
