package main

import (
	"os"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/trim"
)

func main() {
	var (
		cli = cmdline.NewBasicParser()
		t   = trim.Trimmer{
			Columns: cli.Option("-c, --columns").Int(72),
			Suffix:  cli.Option("-s, --suffix").String("..."),
			TabWidth: cli.Option("-t, --tab-width",
				"number of spaces",
			).Int(4),

			Home:        os.Getenv("HOME"),
			ReplaceHome: true,
			PathLen: cli.Option("-l, --path-length",
				"trim paths to length",
				"0 means no trimming and min length is 6",
			).Int(20),

			ConsecutiveSpace: cli.Option(
				"-S, --consecutive-space",
				"number of consecutive spaces to replace with one",
				"0 means no replacement",
			).Uint8(4),
		}
	)
	cli.Parse()

	t.Trim(os.Stdout, os.Stdin)
}
