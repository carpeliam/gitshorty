package clean_test

import (
	"io"
	"log"

	"github.com/carpeliam/gitshorty/support"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/carpeliam/gitshorty/clean"
	sc "github.com/carpeliam/gitshorty/generated"
)

var _ = Describe("Clean", func() {
	BeforeEach(func() {
		log.SetOutput(io.Discard)
	})
	Context("IncludeLocalBranches is true", func() {
		It("should delete local branches associated with delivered stories", func() {
			mockGitRepo := &support.MockGitRepository{
				CurrentBranchName: "dev",
				LocalBranchNames: []string{
					"gitshorty-sc-123",
					"gitshorty-sc-456",
					"main",
					"dev",
				},
				DeletedLocalBranches: []string{},
			}
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {
						Id:        123,
						Completed: true,
					},
					456: {
						Id:        456,
						Completed: false,
					},
				},
			}

			deletedBranches, err := clean.CleanBranches(mockGitRepo, mockShortcutClient, clean.CleanOpts{IncludeLocalBranches: true, DryRun: false})

			Expect(err).To(BeNil())
			Expect(deletedBranches).To(Equal([]string{"gitshorty-sc-123"}))
			Expect(mockGitRepo.DeletedLocalBranches).To(Equal([]string{"gitshorty-sc-123"}))
			Expect(mockGitRepo.DeletedRemoteBranches).To(BeNil())
		})
		It("should avoid deleting the current branch", func() {
			mockGitRepo := &support.MockGitRepository{
				CurrentBranchName: "gitshorty-sc-123",
				LocalBranchNames: []string{
					"gitshorty-sc-123",
					"main",
				},
				DeletedLocalBranches: []string{},
			}
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {
						Id:        123,
						Completed: true,
					},
				},
			}

			deletedBranches, err := clean.CleanBranches(mockGitRepo, mockShortcutClient, clean.CleanOpts{IncludeLocalBranches: true, DryRun: false})

			Expect(err).To(BeNil())
			Expect(deletedBranches).To(Equal([]string{}))
			Expect(mockGitRepo.DeletedLocalBranches).To(Equal([]string{}))
			Expect(mockGitRepo.DeletedRemoteBranches).To(BeNil())
		})
		It("returns branches without deleting them if dry run", func() {
			mockGitRepo := &support.MockGitRepository{
				CurrentBranchName: "main",
				LocalBranchNames: []string{
					"gitshorty-sc-123",
					"main",
				},
				DeletedLocalBranches: []string{},
			}
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {
						Id:        123,
						Completed: true,
					},
				},
			}

			deletedBranches, err := clean.CleanBranches(mockGitRepo, mockShortcutClient, clean.CleanOpts{IncludeLocalBranches: true, DryRun: true})

			Expect(err).To(BeNil())
			Expect(deletedBranches).To(Equal([]string{"gitshorty-sc-123"}))
			Expect(mockGitRepo.DeletedLocalBranches).To(Equal([]string{}))
			Expect(mockGitRepo.DeletedRemoteBranches).To(BeNil())
		})
	})
	Context("IncludeRemoteBranches is true", func() {
		It("should delete remote branches associated with delivered stories", func() {
			mockGitRepo := &support.MockGitRepository{
				CurrentBranchName: "dev",
				LocalBranchNames:  []string{"dev"},
				RemoteBranchNames: []string{
					"origin/gitshorty-sc-123",
					"upstream/gitshorty-sc-123",
					"origin/gitshorty-sc-456",
					"origin/main",
				},
			}
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {
						Id:        123,
						Completed: true,
					},
					456: {
						Id:        456,
						Completed: false,
					},
				},
			}

			deletedBranches, err := clean.CleanBranches(mockGitRepo, mockShortcutClient, clean.CleanOpts{IncludeRemoteBranches: true, DryRun: false})

			Expect(err).To(BeNil())
			Expect(deletedBranches).To(Equal([]string{"origin/gitshorty-sc-123", "upstream/gitshorty-sc-123"}))
			Expect(mockGitRepo.DeletedRemoteBranches).To(Equal([]string{"origin/gitshorty-sc-123", "upstream/gitshorty-sc-123"}))
			Expect(mockGitRepo.DeletedLocalBranches).To(BeNil())
		})
		It("returns branches without deleting them if dry run", func() {
			mockGitRepo := &support.MockGitRepository{
				CurrentBranchName: "main",
				LocalBranchNames:  []string{"dev"},
				RemoteBranchNames: []string{
					"origin/gitshorty-sc-123",
					"upstream/gitshorty-sc-123",
					"origin/gitshorty-sc-456",
					"origin/main",
				},
				DeletedRemoteBranches: []string{},
			}
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {
						Id:        123,
						Completed: true,
					},
				},
			}

			deletedBranches, err := clean.CleanBranches(mockGitRepo, mockShortcutClient, clean.CleanOpts{IncludeRemoteBranches: true, DryRun: true})

			Expect(err).To(BeNil())
			Expect(deletedBranches).To(Equal([]string{"origin/gitshorty-sc-123", "upstream/gitshorty-sc-123"}))
			Expect(mockGitRepo.DeletedRemoteBranches).To(Equal([]string{}))
			Expect(mockGitRepo.DeletedLocalBranches).To(BeNil())
		})
	})
	Context("IncludeLocalBranches and IncludeRemoteBranches are true", func() {
		It("should delete local and remote branches associated with delivered stories", func() {
			mockGitRepo := &support.MockGitRepository{
				CurrentBranchName: "dev",
				LocalBranchNames:  []string{"main", "gitshorty-sc-123"},
				RemoteBranchNames: []string{
					"origin/main",
					"origin/gitshorty-sc-123",
				},
			}
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {
						Id:        123,
						Completed: true,
					},
				},
			}
			deletedBranches, err := clean.CleanBranches(mockGitRepo, mockShortcutClient, clean.CleanOpts{IncludeLocalBranches: true, IncludeRemoteBranches: true})

			Expect(err).To(BeNil())
			Expect(deletedBranches).To(Equal([]string{"gitshorty-sc-123", "origin/gitshorty-sc-123"}))
		})
	})
})
