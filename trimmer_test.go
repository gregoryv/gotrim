package trim

import (
	"os"
	"strings"
)

func ExampleTrimmer_Trim() {
	r := strings.NewReader(`A short line

Next line has $HOME replaced with ~
/home/johndoe/src/github.com/gregoryv/trim`)

	t := NewTrimmer()
	t.Home = "/home/johndoe"
	t.Trim(os.Stdout, r)

	// output:
	// A short line
	//
	// Next line has $HOME replaced with ~
	// ~/src/github.com/gregoryv/trim
}
