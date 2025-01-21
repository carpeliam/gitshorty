package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/carpeliam/gitshorty/browse"
	"github.com/carpeliam/gitshorty/clean"
	sc "github.com/carpeliam/gitshorty/generated"
	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/shortcut"
	"github.com/carpeliam/gitshorty/tasks"
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
						Name:    "confirm",
						Aliases: []string{"c", "y"},
						Usage:   "auto-confirm deletion, without prompting",
					},
					&cli.BoolFlag{
						Name:  "dry-run",
						Usage: "print out branches that would be deleted",
					},
					&cli.BoolFlag{
						Name:  "local",
						Usage: "delete local branches",
						Value: true,
					},
					&cli.BoolFlag{
						Name:  "remote",
						Usage: "delete remote branches",
						Value: false,
					},
				},
				Action: func(ctx *cli.Context) error {
					if !ctx.Bool("local") && !ctx.Bool("remote") {
						return errors.New("must specify at least one of --local or --remote")
					}
					if !ctx.Bool("dry-run") && !ctx.Bool("confirm") {
						shouldContinue := prompter.YesNo("This is a destructive operation. Are you sure you want to continue? You can supply --dry-run to see what would be deleted without actually deleting anything, or --confirm to skip this prompt in the future.", false)
						if !shouldContinue {
							return nil
						}
					}
					git := git.NewRepository()
					shortcutClient := shortcut.NewShortcutClient(ctx.String("api-token"))
					deletedBranches, err := clean.CleanBranches(git, shortcutClient, clean.CleanOpts{
						DryRun:                ctx.Bool("dry-run"),
						IncludeLocalBranches:  ctx.Bool("local"),
						IncludeRemoteBranches: ctx.Bool("remote"),
					})
					if err == nil {
						slog.Info("Deleted branches", "branches", deletedBranches)
					}
					return err
				},
			},
			{
				Name:  "tasks",
				Usage: "display tasks associated with the current branch's story",
				Action: func(ctx *cli.Context) error {
					git := git.NewRepository()
					shortcutClient := shortcut.NewShortcutClient(ctx.String("api-token"))
					tasks, err := tasks.ListTasks(git, shortcutClient)
					if err == nil {
						if len(tasks) == 0 {
							fmt.Println("No tasks found")
						} else {
							fmt.Println(strings.Join(taskList(tasks), "\n"))
						}
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

func taskList(tasks []sc.Task) []string {
	taskStrings := make([]string, len(tasks))
	for i, task := range tasks {
		var taskString strings.Builder
		taskString.WriteString("[")
		if task.Complete {
			taskString.WriteString("x")
		} else {
			taskString.WriteString(" ")
		}
		taskString.WriteString("] ")
		taskString.WriteString(task.Description)
		taskStrings[i] = taskString.String()
	}
	return taskStrings
}
