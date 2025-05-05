package gitshorty_test

import (
	sc "github.com/carpeliam/gitshorty/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/carpeliam/gitshorty/gitshorty"
	"github.com/carpeliam/gitshorty/support"
)

var _ = Describe("Getting stories", func() {
	Describe("GetStoryForCurrentBranch", func() {
		It("gets a Shortcut Story based on the current git branch", func() {
			mockGitRepo := &support.MockGitRepository{
				CurrentBranchName: "coolwhip-sc-123",
			}
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {Id: 123},
				},
			}
			gitshorty := NewGitShorty(mockGitRepo, mockShortcutClient)

			story, err := gitshorty.GetStoryByBranchName("coolwhip-sc-123")

			Expect(err).To(BeNil())
			Expect(story.Id).To(BeEquivalentTo(123))
		})
	})

	Describe("GetStoryByBranchName", func() {
		It("gets a Shortcut Story based on the given git branch name", func() {
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {Id: 123},
				},
			}
			gitshorty := NewGitShorty(&support.MockGitRepository{}, mockShortcutClient)

			story, err := gitshorty.GetStoryByBranchName("coolwhip-sc-123")

			Expect(err).To(BeNil())
			Expect(story.Id).To(BeEquivalentTo(123))
		})
	})

	Describe("GetAcceptedStories", func() {
		It("gets stories and their branches that have been completed", func() {
			mockGitRepo := &support.MockGitRepository{
				LocalBranchNames: []string{
					"gitshorty-sc-123",
					"gitshorty-sc-456",
					"main",
					"dev",
				},
				RemoteBranchNames: []string{
					"origin/gitshorty-sc-123",
					"origin/gitshorty-sc-456",
					"origin/main",
				},
			}
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {Id: 123, Name: "All's well that ends well", Completed: true},
					456: {Id: 456, Completed: false},
				},
			}
			gitshorty := NewGitShorty(mockGitRepo, mockShortcutClient)

			stories, err := gitshorty.GetAcceptedStories()
			Expect(err).To(BeNil())
			Expect(stories).To(HaveLen(1))
			Expect(stories[0]).To(Equal(Story{
				Id:        123,
				Completed: true,
				Name:      "All's well that ends well",
				Tasks:     []Task{},
				Branches: []Branch{
					{Name: "gitshorty-sc-123"},
					{Name: "origin/gitshorty-sc-123"},
				}},
			))
		})
	})

	Describe("GetMyStories", func() {
		It("gets stories and their branches associated with token holder", func() {
			mockGitRepo := &support.MockGitRepository{
				LocalBranchNames: []string{"main", "deliver-sc-123", "blocked-sc-456"}, RemoteBranchNames: []string{"origin/deliver-sc-123"},
			}
			mockShortcutClient := &support.MockShortcutClient{
				Member: sc.MemberInfo{Name: "Susan Saranwrap", MentionName: "susansaranwrap"},
				SearchStoriesByOwner: map[string]sc.StorySearchResults{
					"susansaranwrap": {Data: []sc.StorySearchResult{{Id: 123, Name: "Deliverance"}}},
				},
			}
			gitshorty := NewGitShorty(mockGitRepo, mockShortcutClient)

			stories, err := gitshorty.GetMyStories()
			Expect(err).To(BeNil())
			Expect(stories).To(HaveLen(1))
			Expect(stories[0]).To(Equal(Story{Id: 123, Name: "Deliverance", Branches: []Branch{{Name: "deliver-sc-123"}, {Name: "origin/deliver-sc-123"}}}))
		})
	})
	It("also handles errors", Pending)
})
