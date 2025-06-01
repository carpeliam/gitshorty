package main

import (
	// "errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	// "github.com/Songmu/prompter"
	"github.com/carpeliam/gitshorty/browse"
	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/gitshorty"
	"github.com/carpeliam/gitshorty/shortcut"
	"github.com/carpeliam/gitshorty/tasks"
	"github.com/carpeliam/gitshorty/ui"
	"github.com/carpeliam/gitshorty/version"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/urfave/cli/v2"
)

// type model struct {
//     // choices  []string           // items on the to-do list
// 	stories []gitshorty.Story
//     cursor   int                // which to-do list item our cursor is pointing at
//     selected map[int]struct{}   // which to-do items are selected
// 	input *huh.MultiSelect[int64]
// 	form   *huh.Form
// }

// func storiesToOptions(_stories []gitshorty.Story) *huh.MultiSelect[int64] {
// 	// options := make([]huh.Option[int64], len(stories))
// 	// for i, story := range stories {
// 	// 	options[i] = huh.NewOption(story.Name, story.Id).Selected(false)
// 	// }
// 	options := huh.NewOptions[int64](1, 2, 3)
// 	return huh.NewMultiSelect[int64]().Key("opts").Options(options...).Title("bam")
// }

// func initialModel(stories ...gitshorty.Story) model {
// 	return model{
// 		// Our to-do list is a grocery list
// 		// choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
// 		stories: stories,

// 		// A map which indicates which choices are selected. We're using
// 		// the  map like a mathematical set. The keys refer to the indexes
// 		// of the `choices` slice, above.

// 		selected: make(map[int]struct{}),

// 		input: storiesToOptions(stories),
// 		form: huh.NewForm(huh.NewGroup(storiesToOptions(stories))),
// 	}
// }
// func (m model) Init() tea.Cmd {
//     // Just return `nil`, which means "no I/O right now, please."
//     return m.form.Init()
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	slog.Info("cmd rcvd", "msg", msg)
// 	switch msg := msg.(type) {
// 	// case tea.WindowSizeMsg:
// 	// 	m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()

// 	case tea.KeyMsg:
// 		slog.Info("msg.String", "s", msg.String())
// 		switch msg.String() {
// 		// case "ctrl+c":
// 		// 	return m,
// 		case "esc", "q":
// 			return m, tea.Quit
// 		}
// 	}

// 	var cmds []tea.Cmd

// 	// Process the form
// 	form, cmd := m.form.Update(msg)
// 	f, ok := form.(*huh.Form)
// 	if ok {
// 		m.form = f
// 		cmds = append(cmds, cmd)
// 	}
// 	slog.Info("success", "ok", ok, "f", f, "cmd", cmd, "opts", f.Get("opts"))
// 	// form, cmd := m.form.Update(msg)
// 	// if f, ok := form.(*huh.Form); ok {
// 	// 	m.form = f
// 	// 	cmds = append(cmds, cmd)
// 	// }

// 	// m.input.
// 	if m.form.State == huh.StateCompleted {
// 		// Quit when the form is done.
// 		cmds = append(cmds, tea.Quit)
// 	}

// 	return m, tea.Batch(cmds...)
// }

// func (m model) View() string {
// 	return m.form.View()
// }

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

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
					gs := newGitShorty(ctx.String("api-token"))
					browser := browse.NewBrowser()
					return browse.BrowseStory(gs, browser)
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
					// if !ctx.Bool("local") && !ctx.Bool("remote") {
					// 	return errors.New("must specify at least one of --local or --remote")
					// }
					// if !ctx.Bool("dry-run") && !ctx.Bool("confirm") {
					// 	shouldContinue := prompter.YesNo("This is a destructive operation. Are you sure you want to continue? You can supply --dry-run to see what would be deleted without actually deleting anything, or --confirm to skip this prompt in the future.", false)
					// 	if !shouldContinue {
					// 		return nil
					// 	}
					// }
					// git := git.NewRepository()
					// shortcutClient := shortcut.NewShortcutClient(ctx.String("api-token"))
					f, err := tea.LogToFile("debug.log", "debug")
					if err != nil {
						fmt.Println("fatal:", err)
						os.Exit(1)
					}

					// sc := newGitShorty(ctx.String("api-token"))
					// stories, _ := sc.GetAcceptedStories()
					stories := []gitshorty.Story{
						{Id: 1, Name: "A story", Branches: []gitshorty.Branch{
							{Name: "test1-sc-1"},
							{Name: "test2-sc-1"},
						}},
						{Id: 2, Name: "Another story", Branches: []gitshorty.Branch{}},
					}

					m, err := tea.NewProgram(ui.NewModel(stories)).Run()
					branches := m.(ui.Model).Branches
					slog.Info(fmt.Sprint(branches))

					defer f.Close()

					return err

					// deletedBranches, err := clean.CleanBranches(git, shortcutClient, clean.CleanOpts{
					// 	DryRun:                ctx.Bool("dry-run"),
					// 	IncludeLocalBranches:  ctx.Bool("local"),
					// 	IncludeRemoteBranches: ctx.Bool("remote"),
					// })
					// if deletedBranches != nil {
					// 	fmt.Printf("Deleted branches: %s\n", strings.Join(deletedBranches, ", "))
					// 	return nil
					// }
					// return err
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
					shortcutClient := shortcut.NewShortcutClient(ctx.String("api-token"))
					gs := gitshorty.NewGitShorty(git.NewRepository(), shortcutClient)

					story, err := gs.GetStoryForCurrentBranch()
					if err == nil {
						if story == nil {
							return fmt.Errorf("no story found for current branch")
						}
						shortcutTasks := story.Tasks
						if len(shortcutTasks) == 0 {
							fmt.Println("No tasks found for the associated story")
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
					gs := newGitShorty(ctx.String("api-token"))
					stories, err := gs.GetMyStories()
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

func newGitShorty(apiToken string) gitshorty.GitShorty {
	return gitshorty.NewGitShorty(git.NewRepository(), shortcut.NewShortcutClient(apiToken))
}

// func nestedMultiSelect(stories ...gitshorty.Story) model {
// 	return initialModel(stories...)
// }

func taskListStatic(tasks []gitshorty.Task) []string {
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

func taskListDynamic(tasks []gitshorty.Task) *huh.MultiSelect[int64] {
	options := make([]huh.Option[int64], len(tasks))
	for i, task := range tasks {
		options[i] = huh.NewOption(task.Description, task.Id).Selected(task.Complete)
	}
	return huh.NewMultiSelect[int64]().Options(options...).Title("Press SPACE to toggle, ENTER to submit, '/' to filter")
}
