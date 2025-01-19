package browse_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/carpeliam/gitshorty/browse"
	"github.com/carpeliam/gitshorty/sc"
)

type MockGitRepository struct {
	currentBranchName string
}

func (repository MockGitRepository) GetCurrentBranchName() string {
	return repository.currentBranchName
}

type BrowserSpy struct {
	openedURL string
}

func (browserSpy *BrowserSpy) OpenURL(URL string) (*exec.Cmd, error) {
	browserSpy.openedURL = URL
	return nil, nil
}

type MockShortcutReader struct {
	story sc.Story
}

func (mockShortcutReader MockShortcutReader) GetStory(publicID int) (sc.Story, error) {
	return mockShortcutReader.story, nil
}

var _ = Describe("BrowseStory", func() {
	It("should open the browser to the URL of the story", func() {
		mockGitRepo := MockGitRepository{currentBranchName: "gitshorty-sc-123"}
		browserSpy := &BrowserSpy{}
		mockShortcutReader := MockShortcutReader{
			story: sc.Story{AppUrl: "https://app.shortcut.com/gitshorty/story/123"},
		}

		error := browse.BrowseStory(mockGitRepo, mockShortcutReader, browserSpy)

		Expect(error).To(BeNil())
		Expect(browserSpy.openedURL).To(Equal("https://app.shortcut.com/gitshorty/story/123"))
	})
})
