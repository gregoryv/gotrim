package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gregoryv/cmdline"
)

func main() {
	var (
		cli    = cmdline.NewBasicParser()
		cols   = cli.Option("-c, --columns").Int(72)
		suffix = cli.Option("-s, --suffix").String("...")
	)
	cli.Parse()
	s := bufio.NewScanner(os.Stdin)

	home := os.Getenv("HOME")

	for s.Scan() {
		line := s.Text()
		line = strings.ReplaceAll(line, home, "~")
		if len(line) > cols {
			line = fmt.Sprintf("%s%s", line[:cols], suffix)
		}

		fmt.Println(line)
	}
}
