package usecases

import (
	"regexp"
	"strconv"

	sc "github.com/carpeliam/gitshorty/generated"
	"github.com/carpeliam/gitshorty/shortcut"
)

func GetStoryByBranchName(branchName string, shortcutClient shortcut.Client) (*sc.Story, error) {
	storyId, error := getStoryId(branchName)
	if error != nil {
		return nil, error
	}
	if storyId == nil {
		return nil, nil
	}
	story, err := shortcutClient.GetStory(*storyId)
	return &story, err
}

func getStoryId(branchName string) (*int, error) {
	re := regexp.MustCompile(`sc-(\d+)$`)
	match := re.FindStringSubmatch(branchName)
	if len(match) != 2 {
		return nil, nil
	}
	id, err := strconv.Atoi(match[1])
	return &id, err
}
