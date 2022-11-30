package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/vt100"
)

func main() {
	var (
		cli    = cmdline.NewBasicParser()
		cols   = cli.Option("-c, --columns").Int(72)
		suffix = cli.Option("-s, --suffix").String("...")

		tabWidth = cli.Option("-t, --tab-width", "number of spaces").Int(4)
		tab      = strings.Repeat(" ", tabWidth)
	)
	cli.Parse()
	s := bufio.NewScanner(os.Stdin)

	home := os.Getenv("HOME")
	at := vt100.Attributes()
	for s.Scan() {
		line := s.Text()
		line = strings.ReplaceAll(line, home, "~")
		// replace tabs
		if tabWidth > 0 {
			line = strings.ReplaceAll(line, "\t", tab)
		}

		if len(line) > cols {
			line = fmt.Sprintf("%s%s%v", line[:cols], at.Reset, suffix)
		}

		fmt.Println(line)
	}
}
