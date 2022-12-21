package trim

import (
	"bufio"
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
	}
}

type Trimmer struct {
	Columns  int    // Max line length
	Suffix   string // Added if trimmed
	TabWidth int    // Number of spaces tabs are replaced with

	ReplaceHome bool // Replace occurences of Trimmer.Home with ~
	Home        string
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

		if len(line) > t.Columns {
			line = fmt.Sprintf("%s%s%v", line[:t.Columns], at.Reset, t.Suffix)
		}

		fmt.Fprintln(w, line)
	}
}
