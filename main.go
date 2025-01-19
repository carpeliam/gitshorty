package main

import (
	"log"
	"os"

	"github.com/carpeliam/gitshorty/browse"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "sc",
		Usage:                "command line client for Shortcut",
		Version:              "0.0.1",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "browse",
				Usage: "open story associated with current branch in web browser",
				Action: func(*cli.Context) error {
					apiToken := os.Getenv("SHORTCUT_API_TOKEN")
					shortcut := browse.NewShortcutReader(apiToken)
					browser := browse.NewBrowser()
					git := browse.NewRepository()
					return browse.BrowseStory(git, shortcut, browser)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
