package main

import (
	"log"
	"os"

	"github.com/carpeliam/gitshorty/browse"
	"github.com/carpeliam/gitshorty/version"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "sc",
		Usage:                "command-line client for Shortcut",
		Version:              version.Version,
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "api-token",
				Usage:    "API Token to use with Shortcut API",
				EnvVars:  []string{"SHORTCUT_API_TOKEN"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "browse",
				Usage: "open story associated with current branch in web browser",
				Action: func(ctx *cli.Context) error {
					apiToken := ctx.String("api-token")
					git := browse.NewRepository()
					shortcut := browse.NewShortcutReader(apiToken)
					browser := browse.NewBrowser()
					return browse.BrowseStory(git, shortcut, browser)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
