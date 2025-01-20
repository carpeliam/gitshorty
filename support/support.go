package support

import sc "github.com/carpeliam/gitshorty/generated"

type MockGitRepository struct {
	CurrentBranchName string
	LocalBranchNames  []string
	RemoteBranchNames []string
	DeletedBranches   []string
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

func (repository *MockGitRepository) DeleteBranch(branchName string) error {
	repository.DeletedBranches = append(repository.DeletedBranches, branchName)
	return nil
}

type MockShortcutClient struct {
	Stories map[int]sc.Story
}

func (mockShortcutClient MockShortcutClient) GetStory(publicID int) (sc.Story, error) {
	return mockShortcutClient.Stories[publicID], nil
}
