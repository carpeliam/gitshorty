package tasks

import (
	"slices"

	sc "github.com/carpeliam/gitshorty/generated"
	"github.com/carpeliam/gitshorty/gitshorty"
	"github.com/carpeliam/gitshorty/shortcut"
)

func GetTaskChanges(tasks []gitshorty.Task, ids []int64) map[int64]sc.UpdateTask {
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
