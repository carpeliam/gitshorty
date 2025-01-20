package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/carpeliam/gitshorty/browse"
	"github.com/carpeliam/gitshorty/clean"
	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/shortcut"
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
					git := git.NewRepository()
					shortcutClient := shortcut.NewShortcutClient(ctx.String("api-token"))
					browser := browse.NewBrowser()
					return browse.BrowseStory(git, shortcutClient, browser)
				},
			},
			{
				Name:  "clean",
				Usage: "delete local branches associated with delivered stories",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:     "force",
						Aliases:  []string{"f"},
						Usage:    "currently required in order to delete local branches",
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					git := git.NewRepository()
					shortcutClient := shortcut.NewShortcutClient(ctx.String("api-token"))
					deletedBranches, err := clean.CleanLocalBranches(git, shortcutClient)
					if err == nil {
						slog.Info("Deleted branches", "branches", deletedBranches)
					}
					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
