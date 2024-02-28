package main

import (
	_ "embed"

	"github.com/alecthomas/kong"
	"github.com/hoomy-official/exporter-unifi-protect/cmd/exporter-unifi-protect/commads"
	c "github.com/hoomy-official/go-shared/pkg/cmd"
)

const (
	name        = "unifi-protect"
	description = "Exporter for Unifi protect"
)

//nolint:gochecknoglobals // these global variables exist to be overridden during build
var (
	license string

	version     = "dev"
	commit      = "dirty"
	date        = "latest"
	buildSource = "source"
)

func main() {
	cli := CMD{
		Commons: &c.Commons{
			Version: c.NewVersion(name, version, commit, buildSource, date),
			Licence: c.NewLicence(license),
		},
		Serve: &commads.Serve{},
	}

	ctx := kong.Parse(
		&cli,
		kong.Name(name),
		kong.Description(description),
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run(cli.Commons))
}

type CMD struct {
	*c.Commons
	Serve *commads.Serve `cmd:"serve"`
}
