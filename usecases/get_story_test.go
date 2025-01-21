package usecases_test

import (
	sc "github.com/carpeliam/gitshorty/generated"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/carpeliam/gitshorty/support"
	. "github.com/carpeliam/gitshorty/usecases"
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
})
