package tasks

import (
	sc "github.com/carpeliam/gitshorty/generated"
	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/shortcut"
	"github.com/carpeliam/gitshorty/usecases"
)

func ListTasks(repo git.Repository, shortcutClient shortcut.Client) ([]sc.Task, error) {
	story, err := usecases.GetStoryByBranchName(repo.GetCurrentBranchName(), shortcutClient)
	return story.Tasks, err
}
