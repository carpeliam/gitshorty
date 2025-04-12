package browse

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/gitshorty"
	"github.com/carpeliam/gitshorty/shortcut"
)

func getStoryId(branchName string) (int, error) {
	re := regexp.MustCompile(`sc-(\d+)$`)
	match := re.FindStringSubmatch(branchName)
	if len(match) != 2 {
		return 0, fmt.Errorf("story ID not found in branch name '%s'", branchName)
	}
	return strconv.Atoi(match[1])
}

func BrowseStory(repository git.Repository, shortcutClient shortcut.Client, browser Browser) error {
	currentBranch := repository.GetCurrentBranchName()
	story, err := gitshorty.GetStoryByBranchName(currentBranch, shortcutClient)
	if err != nil {
		return err
	}
	if story == nil {
		return fmt.Errorf("story not found for branch '%s'", currentBranch)
	}
	_, err = browser.OpenURL(story.AppUrl)
	return err
}
