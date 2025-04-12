package gitshorty

import (
	"log/slog"
	"regexp"
	"strconv"

	sc "github.com/carpeliam/gitshorty/generated"
	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/shortcut"
)

type Branch struct {
	Name string
}

type Story struct {
	Id       int64
	Name     string
	Branches []Branch
}

type GitShorty struct {
	repository     git.Repository
	shortcutClient shortcut.Client
}

func NewGitShorty(repository git.Repository,
	shortcutClient shortcut.Client) GitShorty {
	return GitShorty{repository: repository, shortcutClient: shortcutClient}
}

func (gs GitShorty) GetStoryForCurrentBranch() (*sc.Story, error) {
	branchName := gs.repository.GetCurrentBranchName()
	return gs.GetStoryByBranchName(branchName)
}

func (gs GitShorty) GetStoryByBranchName(branchName string) (*sc.Story, error) {
	storyId, error := getStoryId(branchName)
	if error != nil {
		return nil, error
	}
	if storyId == nil {
		return nil, nil
	}
	story, err := gs.shortcutClient.GetStory(*storyId)
	return &story, err
}

func (gs GitShorty) GetMyStories() ([]Story, error) {
	localBranches := gs.repository.GetLocalBranchNames()
	remoteBranches := gs.repository.GetRemoteBranchNames()
	memberInfo, err := gs.shortcutClient.GetMemberInfo()
	if err != nil {
		return nil, err
	}
	searchResults, err := gs.shortcutClient.SearchStories(memberInfo.MentionName)
	if err != nil {
		return nil, err
	}
	stories := make([]Story, len(searchResults.Data))
	for i, searchResult := range searchResults.Data {
		branches := append(
			getBranchesForStoryId(localBranches, searchResult.Id),
			getBranchesForStoryId(remoteBranches, searchResult.Id)...,
		)
		stories[i] = Story{Id: searchResult.Id, Name: searchResult.Name, Branches: branches}
	}
	return stories, nil
}

func getBranchesForStoryId(branchNames []string, storyId int64) []Branch {
	branches := []Branch{}
	for _, branchName := range branchNames {
		id, err := getStoryId(branchName)
		if err != nil {
			slog.Error("Error getting ID", "error", err)
		}
		if id != nil && *id == int(storyId) {
			branches = append(branches, Branch{Name: branchName})
		}
	}
	return branches
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
