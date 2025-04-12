package tasks_test

import (
	"github.com/carpeliam/gitshorty/gitshorty"
	"github.com/carpeliam/gitshorty/support"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sc "github.com/carpeliam/gitshorty/generated"
	"github.com/carpeliam/gitshorty/tasks"
)

var _ = Describe("Tasks", func() {
	It("returns a list of tasks", func() {
		expectedTasks := []sc.Task{
			{Id: 1, Description: "Task 1"},
			{Id: 2, Description: "Task 2"},
			{Id: 3, Description: "Task 3"},
		}
		mockGitRepo := &support.MockGitRepository{CurrentBranchName: "magic-sc-123"}
		mockShortcutClient := &support.MockShortcutClient{
			Stories: map[int]sc.Story{
				123: {
					Tasks: expectedTasks,
				},
			},
		}
		returnedTasks, err := tasks.ListTasks(gitshorty.NewGitShorty(mockGitRepo, mockShortcutClient))

		Expect(err).To(BeNil())
		Expect(returnedTasks).To(Equal(expectedTasks))
	})

	It("should return an error if the story is not found", func() {
		mockGitRepo := &support.MockGitRepository{CurrentBranchName: "main"}
		mockShortcutClient := &support.MockShortcutClient{}

		returnedTasks, err := tasks.ListTasks(gitshorty.NewGitShorty(mockGitRepo, mockShortcutClient))

		Expect(err).NotTo(BeNil())
		Expect(returnedTasks).To(BeNil())
	})

	It("detects changes in task completion", func() {
		storyTasks := []sc.Task{
			{Id: 1, StoryId: 123, Description: "Task 1", Complete: false},
			{Id: 2, StoryId: 123, Description: "Task 2", Complete: false},
			{Id: 3, StoryId: 123, Description: "Task 3", Complete: true},
			{Id: 4, StoryId: 123, Description: "Task 4", Complete: true},
		}
		ids := []int64{1, 3}

		changes := tasks.GetTaskChanges(storyTasks, ids)
		Expect(changes).To(HaveKeyWithValue(int64(1), sc.UpdateTask{Complete: true}))
		Expect(changes).To(HaveKeyWithValue(int64(4), sc.UpdateTask{Complete: false}))
	})

	It("submits tasks updates to Shortcut", func() {
		mockShortcutClient := &support.MockShortcutClient{
			Stories: map[int]sc.Story{
				123: {
					Tasks: []sc.Task{
						{Id: 1, StoryId: 123, Description: "Task 1", Complete: false},
					},
				},
			},
		}
		err := tasks.UpdateTasks(mockShortcutClient, 123, map[int64]sc.UpdateTask{
			1: {Complete: true},
		})

		Expect(err).To(BeNil())
		Expect(mockShortcutClient.TaskUpdates[123]).To(HaveKeyWithValue(int64(1), sc.UpdateTask{Complete: true}))
	})
})
