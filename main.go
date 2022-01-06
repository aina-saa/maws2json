package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/aina-saa/maws2json/process"
	"github.com/aina-saa/maws2json/version"
)

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	info := fmt.Sprintf("%s ver. %s\n %s built at %s on %s\n Author: %s", app.Model.Name, version.BuildVersion, version.BuildSha, version.BuildTime, version.BuildHost, version.Author)
	fmt.Println(info)
	app.Exit(0)
	return nil
}

var CLI struct {
	File      string      `help:"Input file or '-' for stdin" name:"file" type:"existingfile" short:"f" default:"-"`
	Timestamp bool        `help:"Add timestamp to output." short:"t" negatable:"true" default:"true"`
	Wind      bool        `help:"Process WIND messages." negatable:"true" default:"true"`
	Log       bool        `help:"Process LOG messages." negatable:"true" default:"true"`
	Ptu       bool        `help:"Process PTU messages." negatable:"true" default:"true"`
	Version   VersionFlag `name:"version" help:"Print version information and quit"`
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
