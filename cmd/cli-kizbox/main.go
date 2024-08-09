package main

import (
	_ "embed"

	"github.com/alecthomas/kong"
	"github.com/hoomy-official/cli-kizbox/cmd/cli-kizbox/commands"
	"github.com/hoomy-official/cli-kizbox/cmd/cli-kizbox/commands/devices"
	"github.com/hoomy-official/cli-kizbox/cmd/cli-kizbox/globals"
	"github.com/vanyda-official/go-shared/pkg/cmd"
)

//nolint:gochecknoglobals // these global variables exist to be overridden during build
var (
	name    = "kizbox"
	group   = "hoomy"
	license string

	version     = "dev"
	commit      = "dirty"
	date        = "latest"
	buildSource = "source"
)

type CLI struct {
	*cmd.Commons
	*cmd.Config
	*globals.Globals

	Venitian commands.VenitianCmd `cmd:"venitians"`
	Devices  devices.Cmd          `cmd:"devices" help:"list devices available in the current system"`
	Listen   commands.ListenCmd   `cmd:"listen" help:"listen events in the current system"`
	Discover commands.DiscoverCmd `cmd:"discover" help:"list for systems available"`
}

func main() {
	cli := CLI{
		Commons: &cmd.Commons{
			Version: cmd.NewVersion(name, version, commit, buildSource, date),
			Licence: cmd.NewLicence(license),
		},
		Config:  cmd.NewConfig(name, cmd.WithGroup(group)),
		Globals: &globals.Globals{},
	}

	ctx := kong.Parse(
		&cli,
		kong.Name(name),
		kong.Description("Simple cli for managing my home automation"),
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run(cli.Globals, cli.Commons))
}
