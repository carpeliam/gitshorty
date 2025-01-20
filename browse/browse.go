package browse

import (
	"fmt"
	"regexp"
	"strconv"
)

func getStoryId(branchName string) (int, error) {
	re := regexp.MustCompile(`sc-(\d+)$`)
	match := re.FindStringSubmatch(branchName)
	if len(match) != 2 {
		return 0, fmt.Errorf("story ID not found in branch name '%s'", branchName)
	}
	return strconv.Atoi(match[1])
}

func BrowseStory(repository Repository, backlog Backlog, browser Browser) error {
	branch := repository.GetCurrentBranchName()
	storyId, error := getStoryId(branch)
	if error != nil {
		return error
	}
	story, error := backlog.GetStory(storyId)
	if error != nil {
		return error
	}
	_, error = browser.OpenURL(story.AppUrl)
	return error
}
