package main

import (
	"github.com/alecthomas/kong"

	"github.com/aina-saa/maws2json/process"
)

var CLI struct {
	File      string `help:"Input file or '-' for stdin" name:"file" type:"existingfile" short:"f" default:"-"`
	Timestamp bool   `help:"Add timestamp to output." short:"t" negatable:"true" default:"true"`
	Wind      bool   `help:"Process WIND messages." negatable:"true" default:"true"`
	Log       bool   `help:"Process LOG messages." negatable:"true" default:"true"`
	Ptu       bool   `help:"Process PTU messages." negatable:"true" default:"true"`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("maws2json"),
		kong.Description("Reads MAWS formatted data and converts it into JSON output stream."),
	)
	switch ctx.Command() {
	default:
		// we dont have a subcommand so we always drop into default
		process.Process(CLI.File, CLI.Timestamp, CLI.Wind, CLI.Log, CLI.Ptu)
	}
}

// eof
