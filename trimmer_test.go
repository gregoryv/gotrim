package trim

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func ExampleTrimmer_Trim() {
	r := strings.NewReader(`A short line

Next line has $HOME replaced with ~
/home/johndoe/src/github.com/gregoryv/trim
two tabs:		a
three...:			b
`)

	t := NewTrimmer()
	t.Home = "/home/johndoe"
	t.PathLen = 20
	t.ConsecutiveSpace = 4 // replace 4 concecutive spaces with one
	t.Trim(os.Stdout, r)

	// output:
	// A short line
	//
	// Next line has $HOME replaced with ~
	// ~/.../gregoryv/trim
	// two tabs:  a
	// three...:   b
}

func ExampleTrimPaths() {
	dir := `/home/johndoe/src/github.com/gregoryv/trim`
	v := TrimPaths(20, dir)
	fmt.Println(v)
	// output:
	// /.../gregoryv/trim
}

func ExampleTrimPaths_inline() {
	dir := `Code found in ~/src/github.com/gregoryv/trim directory`
	v := TrimPaths(20, dir)
	fmt.Println(v)
	// output:
	// Code found in ~/.../gregoryv/trim directory
}

func TestTrimPaths(t *testing.T) {
	cases := []struct {
		len int
		in  string
		out string
	}{
		{len: 10, in: "/a/b", out: "/a/b"},
		{len: 6, in: "/a/b/c/d", out: "/.../d"},
		{len: -1, in: "~/a/b/c/d", out: "~/.../d"},
		{len: 8, in: "/aabbccddee", out: "/...ddee"},
	}
	for _, c := range cases {
		if v := TrimPaths(c.len, c.in); v != c.out {
			t.Errorf("TrimPaths(%v, %q)\ngot: %q\nexp: %q", c.len, c.in, v, c.out)
		}
	}
}
