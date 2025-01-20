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
	Describe("CleanLocalBranches", func() {
		It("should delete branches associated with delivered stories", func() {
			mockGitRepo := &support.MockGitRepository{
				CurrentBranchName: "dev",
				LocalBranchNames: []string{
					"gitshorty-sc-123",
					"gitshorty-sc-456",
					"main",
					"dev",
				},
				DeletedBranches: []string{},
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

			deletedBranches, err := clean.CleanLocalBranches(mockGitRepo, mockShortcutClient)

			Expect(err).To(BeNil())
			Expect(deletedBranches).To(Equal([]string{"gitshorty-sc-123"}))
			Expect(mockGitRepo.DeletedBranches).To(Equal([]string{"gitshorty-sc-123"}))
		})
		It("should avoid deleting the current branch", func() {
			mockGitRepo := &support.MockGitRepository{
				CurrentBranchName: "gitshorty-sc-123",
				LocalBranchNames: []string{
					"gitshorty-sc-123",
					"main",
				},
				DeletedBranches: []string{},
			}
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {
						Id:        123,
						Completed: true,
					},
				},
			}

			deletedBranches, err := clean.CleanLocalBranches(mockGitRepo, mockShortcutClient)

			Expect(err).To(BeNil())
			Expect(deletedBranches).To(Equal([]string{}))
			Expect(mockGitRepo.DeletedBranches).To(Equal([]string{}))
		})
	})
})
