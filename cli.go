package main

import (
	"github.com/urfave/cli"
)

func NewApp(name, usage, version string) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Action = MatchByToml
	app.Version = version
	app.Commands = getCommands()
	app.Flags = getFlags()
	return app
}

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "show mentor list",
			Action: List,
		},
		{
			Name:   "match",
			Usage:  "match enter and mentor",
			Action: Match,
		},
	}
}

func getFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "d directory",
			Usage: "directory of mentor.toml and ratio.toml",
		},
	}
}
