package helpcontent

import (
	"io/ioutil"
	"testing"
)

func TestHelpcontent(t *testing.T) {
	t.Run("read from helpcntnt file",
		func(t *testing.T) {
			filname := "helpcontent"
			expected, _ := ioutil.ReadFile(filname)
			actual := Helpcontent()
			if string(expected) != actual {
				t.Errorf("expected:-%q actual:-%q", expected, actual)
			}

		})
}
