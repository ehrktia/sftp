package helpcontent

import (
	"testing"
)

func TestHelpcontent(t *testing.T) {
	t.Run("read from helpcntnt file",
		func(t *testing.T) {
			got := Helpcontent()
			t.Logf("%v\n", got)
		})
}
