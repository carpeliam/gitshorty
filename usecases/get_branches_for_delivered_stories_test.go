package usecases_test

import (
	sc "github.com/carpeliam/gitshorty/generated"
	"github.com/carpeliam/gitshorty/git"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/carpeliam/gitshorty/support"
	"github.com/carpeliam/gitshorty/usecases"
)

var _ = Describe("GetBranchesForDeliveredStories", func() {
	localBranchesOnly := usecases.GetBranchesForDeliveredStoriesOptions{IncludeLocal: true}
	remoteBranchesOnly := usecases.GetBranchesForDeliveredStoriesOptions{IncludeRemote: true}
	localAndRemoteBranches := usecases.GetBranchesForDeliveredStoriesOptions{IncludeLocal: true, IncludeRemote: true}
	It("returns local branch names for all delivered stories", func() {
		mockShortcutClient := &support.MockShortcutClient{
			Stories: map[int]sc.Story{
				123: {Id: 123, Completed: true},
				456: {Id: 456, Completed: false},
			},
		}
		mockGitRepo := &support.MockGitRepository{
			LocalBranchNames: []string{
				"gitshorty-sc-123",
				"gitshorty-sc-456",
				"main",
			},
			RemoteBranchNames: []string{"origin/gitshorty-sc-123"},
		}

		branches, err := usecases.GetBranchesForDeliveredStories(mockGitRepo, mockShortcutClient, localBranchesOnly)

		Expect(err).To(BeNil())
		Expect(branches).To(Equal([]git.Branch{{Name: "gitshorty-sc-123", Type: git.Local}}))
	})
	It("returns remote branch names for all delivered stories", func() {
		mockShortcutClient := &support.MockShortcutClient{
			Stories: map[int]sc.Story{
				123: {Id: 123, Completed: true},
				456: {Id: 456, Completed: false},
			},
		}
		mockGitRepo := &support.MockGitRepository{
			LocalBranchNames:  []string{"gitshorty-sc-123", "main"},
			RemoteBranchNames: []string{"origin/gitshorty-sc-123", "origin/gitshorty-sc-456"},
		}

		branches, err := usecases.GetBranchesForDeliveredStories(mockGitRepo, mockShortcutClient, remoteBranchesOnly)

		Expect(err).To(BeNil())
		Expect(branches).To(Equal([]git.Branch{{Name: "origin/gitshorty-sc-123", Type: git.Remote}}))

	})
	It("returns mixed branch names for all delivered stories", func() {
		mockShortcutClient := &support.MockShortcutClient{
			Stories: map[int]sc.Story{
				123: {Id: 123, Completed: true},
				456: {Id: 456, Completed: false},
			},
		}
		mockGitRepo := &support.MockGitRepository{
			LocalBranchNames:  []string{"gitshorty-sc-123", "main"},
			RemoteBranchNames: []string{"origin/gitshorty-sc-123", "origin/gitshorty-sc-456"},
		}

		branches, err := usecases.GetBranchesForDeliveredStories(mockGitRepo, mockShortcutClient, localAndRemoteBranches)

		Expect(err).To(BeNil())
		Expect(branches).To(Equal([]git.Branch{
			{Name: "gitshorty-sc-123", Type: git.Local},
			{Name: "origin/gitshorty-sc-123", Type: git.Remote},
		}))
	})
})
