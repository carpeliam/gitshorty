package git

import (
	"os/exec"
	"strings"
)

type Repository interface {
	GetCurrentBranchName() string
	GetLocalBranchNames() []string
	GetRemoteBranchNames() []string
	DeleteLocalBranch(branchName string) error
	DeleteRemoteBranch(branchName string) error
	DeleteBranch(branch Branch) error
}

type BranchType int

const (
	Local BranchType = iota
	Remote
)

type Branch struct {
	Name string
	Type BranchType
}

type BranchIterator interface {
	GetBranches() []Branch
}
type LocalBranchIterator struct {
	repo Repository
}

func (iterator LocalBranchIterator) GetBranches() []Branch {
	branchNames := iterator.repo.GetLocalBranchNames()
	branches := make([]Branch, len(branchNames))
	for i, branchName := range branchNames {
		branches[i] = Branch{Name: branchName, Type: Local}
	}
	return branches
}
func NewLocalBranchIterator(repo Repository) BranchIterator {
	return LocalBranchIterator{repo: repo}
}

type RemoteBranchIterator struct {
	repo Repository
}

func (iterator RemoteBranchIterator) GetBranches() []Branch {
	branchNames := iterator.repo.GetRemoteBranchNames()
	branches := make([]Branch, len(branchNames))
	for i, branchName := range branchNames {
		branches[i] = Branch{Name: branchName, Type: Remote}
	}
	return branches
}
func NewRemoteBranchIterator(repo Repository) BranchIterator {
	return RemoteBranchIterator{repo: repo}
}

type GitRepository struct{}

func (repository GitRepository) GetCurrentBranchName() string {
	output, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	return strings.TrimSpace(string(output))
}

func (repository GitRepository) GetLocalBranchNames() []string {
	output, _ := exec.Command("git", "--no-pager", "branch", "--format=%(refname:short)").Output()
	return strings.Split(strings.TrimSpace(string(output)), "\n")
}

func (repository GitRepository) GetRemoteBranchNames() []string {
	output, _ := exec.Command("git", "--no-pager", "branch", "--remote", "--format=%(refname:short)").Output()
	return strings.Split(strings.TrimSpace(string(output)), "\n")
}

func (repository GitRepository) DeleteLocalBranch(branchName string) error {
	_, err := exec.Command("git", "branch", "-D", branchName).Output()
	return err
}

func (repository GitRepository) DeleteRemoteBranch(remoteBranchName string) error {
	remoteAndBranch := strings.SplitN(remoteBranchName, "/", 2)
	remote := remoteAndBranch[0]
	branch := remoteAndBranch[1]
	_, err := exec.Command("git", "push", "--delete", remote, branch).Output()
	return err
}

func (repository GitRepository) DeleteBranch(branch Branch) error {
	switch branch.Type {
	case Local:
		return repository.DeleteLocalBranch(branch.Name)
	case Remote:
		return repository.DeleteRemoteBranch(branch.Name)
	}
	return nil
}

func NewRepository() Repository {
	return GitRepository{}
}
