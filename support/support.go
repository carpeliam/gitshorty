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
	Stories map[int]sc.Story
}

func (mockShortcutClient MockShortcutClient) GetStory(publicID int) (sc.Story, error) {
	return mockShortcutClient.Stories[publicID], nil
}
