package trim

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gregoryv/vt100"
)

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
	Columns     int
	Suffix      string
	Home        string
	TabWidth    int
	ReplaceHome bool
}

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
