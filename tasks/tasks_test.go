package tasks_test

import (
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
		mockShortcutClient := support.MockShortcutClient{
			Stories: map[int]sc.Story{
				123: {
					Tasks: expectedTasks,
				},
			},
		}
		returnedTasks, err := tasks.ListTasks(mockGitRepo, mockShortcutClient)

		Expect(err).To(BeNil())
		Expect(returnedTasks).To(Equal(expectedTasks))
	})
})
