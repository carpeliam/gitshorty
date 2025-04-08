package support

import sc "github.com/carpeliam/gitshorty/generated"

type MockGitRepository struct {
	CurrentBranchName     string
	LocalBranchNames      []string
	RemoteBranchNames     []string
	DeletedLocalBranches  []string
	DeletedRemoteBranches []string
}

func (repository MockGitRepository) GetCurrentBranchName() string {
	return repository.CurrentBranchName
}

func (repository MockGitRepository) GetLocalBranchNames() []string {
	return repository.LocalBranchNames
}

func (repository MockGitRepository) GetRemoteBranchNames() []string {
	return repository.RemoteBranchNames
}

func (repository *MockGitRepository) DeleteLocalBranch(branchName string) error {
	repository.DeletedLocalBranches = append(repository.DeletedLocalBranches, branchName)
	return nil
}

func (repository *MockGitRepository) DeleteRemoteBranch(branchName string) error {
	repository.DeletedRemoteBranches = append(repository.DeletedRemoteBranches, branchName)
	return nil
}

type MockShortcutClient struct {
	Member               sc.MemberInfo
	SearchStoriesByOwner map[string]sc.StorySearchResults
	Stories              map[int]sc.Story
	TaskUpdates          map[int]map[int64]sc.UpdateTask
}

func (mockShortcutClient MockShortcutClient) GetMemberInfo() (sc.MemberInfo, error) {
	return mockShortcutClient.Member, nil
}

func (mockShortcutClient MockShortcutClient) GetStory(publicID int) (sc.Story, error) {
	return mockShortcutClient.Stories[publicID], nil
}

func (mockShortcutClient MockShortcutClient) SearchStories(owner string) (sc.StorySearchResults, error) {
	return mockShortcutClient.SearchStoriesByOwner[owner], nil
}

func (mockShortcutClient *MockShortcutClient) UpdateTask(storyID int, tasks map[int64]sc.UpdateTask) error {
	if mockShortcutClient.TaskUpdates == nil {
		mockShortcutClient.TaskUpdates = make(map[int]map[int64]sc.UpdateTask)
	}
	mockShortcutClient.TaskUpdates[storyID] = tasks
	return nil
}
