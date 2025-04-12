package gitshorty_test

import (
	sc "github.com/carpeliam/gitshorty/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/carpeliam/gitshorty/gitshorty"
	"github.com/carpeliam/gitshorty/support"
)

var _ = Describe("Getting stories", func() {
	Describe("GetStoryByBranchName", func() {
		It("gets a Shortcut Story based on the given git branch name", func() {
			mockShortcutClient := &support.MockShortcutClient{
				Stories: map[int]sc.Story{
					123: {Id: 123},
				},
			}

			story, error := GetStoryByBranchName("coolwhip-sc-123", mockShortcutClient)

			Expect(error).To(BeNil())
			Expect(story.Id).To(BeEquivalentTo(123))
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

			stories, error := GetMyStories(mockGitRepo, mockShortcutClient)
			Expect(error).To(BeNil())
			Expect(stories).To(HaveLen(1))
			Expect(stories[0]).To(Equal(Story{Id: 123, Name: "Deliverance", Branches: []Branch{{Name: "deliver-sc-123"}, {Name: "origin/deliver-sc-123"}}}))
		})
	})
})
