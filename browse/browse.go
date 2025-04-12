package browse

import (
	"fmt"

	"github.com/carpeliam/gitshorty/gitshorty"
)

func BrowseStory(gs gitshorty.GitShorty, browser Browser) error {
	story, err := gs.GetStoryForCurrentBranch()
	if err != nil {
		return err
	}
	if story == nil {
		return fmt.Errorf("story not found for current branch")
	}
	_, err = browser.OpenURL(story.AppUrl)
	return err
}
