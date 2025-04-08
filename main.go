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
	"github.com/carpeliam/gitshorty/usecases"
	"github.com/carpeliam/gitshorty/version"
	"github.com/charmbracelet/huh"
	"github.com/urfave/cli/v2"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelError)

	app := &cli.App{
		Name:                   "sc",
		Usage:                  "command-line client for Shortcut",
		Version:                version.Version,
		EnableBashCompletion:   true,
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "api-token",
				Usage:    "API Token to use with Shortcut API",
				EnvVars:  []string{"SHORTCUT_API_TOKEN"},
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Usage:   "enable verbose logging",
				Aliases: []string{"V"},
				Action: func(ctx *cli.Context, verbose bool) error {
					switch ctx.Count("verbose") {
					case 1:
						slog.SetLogLoggerLevel(slog.LevelInfo)
					default:
						slog.SetLogLoggerLevel(slog.LevelDebug)
					}
					return nil
				},
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
					if deletedBranches != nil {
						fmt.Printf("Deleted branches: %s\n", strings.Join(deletedBranches, ", "))
						return nil
					}
					return err
				},
			},
			{
				Name:  "tasks",
				Usage: "display tasks associated with the current branch's story",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "interactive",
						Usage: "check tasks off as you complete them",
						Value: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					git := git.NewRepository()
					shortcutClient := shortcut.NewShortcutClient(ctx.String("api-token"))
					shortcutTasks, err := tasks.ListTasks(git, shortcutClient)
					if err == nil {
						if len(shortcutTasks) == 0 {
							fmt.Println("No shortcutTasks found")
						} else {
							if ctx.Bool("interactive") {
								var values []int64
								if err = taskListDynamic(shortcutTasks).Value(&values).Run(); err != nil {
									return err
								} else {
									storyId := shortcutTasks[0].StoryId
									return tasks.UpdateTasks(shortcutClient, int(storyId), tasks.GetTaskChanges(shortcutTasks, values))
								}
							} else {
								fmt.Println(strings.Join(taskListStatic(shortcutTasks), "\n"))
							}
						}
					}
					return err
				},
			},
			{
				Name:  "mywork",
				Usage: "display branches associated with the current user's stories",
				Action: func(ctx *cli.Context) error {
					git := git.NewRepository()
					shortcutClient := shortcut.NewShortcutClient(ctx.String("api-token"))
					stories, err := usecases.GetMyStories(git, shortcutClient)
					if err == nil {
						for _, story := range stories {
							fmt.Printf("[sc-%d]\t%s\n", story.Id, story.Name)
							for _, branch := range story.Branches {
								fmt.Printf(" * %s\n", branch.Name)
							}
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

func taskListStatic(tasks []sc.Task) []string {
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

func taskListDynamic(tasks []sc.Task) *huh.MultiSelect[int64] {
	options := make([]huh.Option[int64], len(tasks))
	for i, task := range tasks {
		options[i] = huh.NewOption(task.Description, task.Id).Selected(task.Complete)
	}
	return huh.NewMultiSelect[int64]().Options(options...).Title("Press SPACE to toggle, ENTER to submit, '/' to filter")
}
