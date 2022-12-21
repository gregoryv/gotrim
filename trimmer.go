package trim

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gregoryv/vt100"
)

// NewTrimmer returns a trimmer with a column width of 72, replacing
// tabs with 4 spaces and $HOME with ~. Trimmed lines are indicated
// with suffix '...'.
func NewTrimmer() *Trimmer {
	return &Trimmer{
		Columns:     72,
		Suffix:      "...",
		Home:        os.Getenv("HOME"),
		TabWidth:    4,
		ReplaceHome: true,
		PathLen:     20,
	}
}

type Trimmer struct {
	Columns  int    // Max line length
	Suffix   string // Added if trimmed
	TabWidth int    // Number of spaces tabs are replaced with

	ReplaceHome bool // Replace occurences of Trimmer.Home with ~
	Home        string

	PathLen int
}

// Trim trims lines from the reader and writes them to the writer.
// Stops when reader reaches io.EOF.
func (t *Trimmer) Trim(w io.Writer, r io.Reader) {
	var (
		s   = bufio.NewScanner(r)
		tab = strings.Repeat(" ", t.TabWidth)
		at  = vt100.Attributes()
	)

	for s.Scan() {
		line := s.Text()
		if t.ReplaceHome {
			line = strings.ReplaceAll(line, t.Home, "~")
		}
		// replace tabs
		if t.TabWidth > 0 {
			line = strings.ReplaceAll(line, "\t", tab)
		}
		if t.PathLen > 0 {
			line = TrimPaths(t.PathLen, line)
		}
		if len(line) > t.Columns {
			line = fmt.Sprintf("%s%s%v", line[:t.Columns], at.Reset, t.Suffix)
		}
		fmt.Fprintln(w, line)
	}
}

// TrimPaths shortens first path in the given string to max cols. Min
// cols is 6.
func TrimPaths(cols int, in string) string {
	if len(in) < cols {
		return in
	}
	if cols < 6 {
		cols = 6
	}
	r := bufio.NewReader(strings.NewReader(in))

	// before any path
	before, err := r.ReadString('/')
	if err != nil {
		// no path found
		return before
	}
	var buf bytes.Buffer
	buf.WriteString(before)

	// shorten path
	prefix := "/"
	if strings.HasSuffix(prefix, "~/") {
		prefix = "~/"
	}
	sep := "..."
	max := cols - len(prefix) - len(sep)

	dir, _ := r.ReadString(' ')
	if len(dir) > max {
		buf.WriteString(sep)

		short := dir[len(dir)-max:]
		i := strings.Index(short, "/")
		if i > -1 {
			short = short[i:]
		}
		buf.WriteString(short)
	}
	// the rest
	io.Copy(&buf, r)
	return buf.String()
}
