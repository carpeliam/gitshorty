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
	Id        int64
	Completed bool
	AppUrl    string
	Name      string
	Tasks     []Task
	Branches  []Branch
}

type Task struct {
	Id          int64
	Complete    bool
	Description string
	StoryId     int64
}

type GitShorty struct {
	repository     git.Repository
	shortcutClient shortcut.Client
}

func NewGitShorty(repository git.Repository, shortcutClient shortcut.Client) GitShorty {
	return GitShorty{repository: repository, shortcutClient: shortcutClient}
}

func (gs GitShorty) GetStoryForCurrentBranch() (*Story, error) {
	branchName := gs.repository.GetCurrentBranchName()
	return gs.GetStoryByBranchName(branchName)
}

func (gs GitShorty) GetStoryByBranchName(branchName string) (*Story, error) {
	storyId, error := getStoryId(branchName)
	if error != nil {
		return nil, error
	}
	if storyId == nil {
		return nil, nil
	}
	story, err := gs.shortcutClient.GetStory(*storyId)
	wrappedStory := wrapStory(story)
	return &wrappedStory, err
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

func wrapStory(story sc.Story) Story {
	return Story{
		Id:        story.Id,
		AppUrl:    story.AppUrl,
		Completed: story.Completed,
		Tasks:     wrapTasks(story.Tasks),
	}
}

func wrapTasks(scTasks []sc.Task) []Task {
	tasks := make([]Task, len(scTasks))
	for i, scTask := range scTasks {
		tasks[i] = Task{
			Id:          scTask.Id,
			Complete:    scTask.Complete,
			Description: scTask.Description,
			StoryId:     scTask.StoryId,
		}
	}
	return tasks
}
