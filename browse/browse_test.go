package browse_test

import (
	"os/exec"

	"github.com/carpeliam/gitshorty/support"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/carpeliam/gitshorty/browse"
	sc "github.com/carpeliam/gitshorty/generated"
)

type BrowserSpy struct {
	openedURL string
}

func (browserSpy *BrowserSpy) OpenURL(URL string) (*exec.Cmd, error) {
	browserSpy.openedURL = URL
	return nil, nil
}

var _ = Describe("BrowseStory", func() {
	It("should open the browser to the URL of the story", func() {
		mockGitRepo := &support.MockGitRepository{CurrentBranchName: "gitshorty-sc-123"}
		mockShortcutClient := support.MockShortcutClient{
			Stories: map[int]sc.Story{
				123: {AppUrl: "https://app.shortcut.com/gitshorty/story/123"},
			},
		}
		browserSpy := &BrowserSpy{}

		error := browse.BrowseStory(mockGitRepo, mockShortcutClient, browserSpy)

		Expect(error).To(BeNil())
		Expect(browserSpy.openedURL).To(Equal("https://app.shortcut.com/gitshorty/story/123"))
	})

	It("should return an error if the story is not found", func() {
		mockGitRepo := &support.MockGitRepository{CurrentBranchName: "main"}
		mockShortcutClient := support.MockShortcutClient{
			Stories: map[int]sc.Story{},
		}
		browserSpy := &BrowserSpy{}

		err := browse.BrowseStory(mockGitRepo, mockShortcutClient, browserSpy)

		Expect(err).NotTo(BeNil())
	})
})
