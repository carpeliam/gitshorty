package tasks

import (
	"fmt"
	"slices"

	sc "github.com/carpeliam/gitshorty/generated"
	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/shortcut"
	"github.com/carpeliam/gitshorty/usecases"
)

func ListTasks(repo git.Repository, shortcutClient shortcut.Client) ([]sc.Task, error) {
	currentBranch := repo.GetCurrentBranchName()
	story, err := usecases.GetStoryByBranchName(currentBranch, shortcutClient)
	if story == nil {
		return nil, fmt.Errorf("tasks not found for branch '%s'", currentBranch)
	}
	return story.Tasks, err
}

func GetTaskChanges(tasks []sc.Task, ids []int64) map[int64]sc.UpdateTask {
	updates := make(map[int64]sc.UpdateTask)
	for _, task := range tasks {
		shouldBeCompleted := slices.ContainsFunc(ids, func(id int64) bool {
			return id == task.Id
		})
		if task.Complete != shouldBeCompleted {
			updates[task.Id] = sc.UpdateTask{
				Complete: shouldBeCompleted,
			}
		}

	}
	return updates
}

func UpdateTasks(shortcutClient shortcut.Client, storyId int, updates map[int64]sc.UpdateTask) error {
	return shortcutClient.UpdateTask(storyId, updates)
}
