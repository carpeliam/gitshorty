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
}

type BranchIterator interface {
	GetBranchNames() []string
}
type LocalBranchIterator struct {
	repo Repository
}

func (iterator LocalBranchIterator) GetBranchNames() []string {
	return iterator.repo.GetLocalBranchNames()
}
func NewLocalBranchIterator(repo Repository) BranchIterator {
	return LocalBranchIterator{repo: repo}
}

type RemoteBranchIterator struct {
	repo Repository
}

func (iterator RemoteBranchIterator) GetBranchNames() []string {
	return iterator.repo.GetRemoteBranchNames()
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

func NewRepository() Repository {
	return GitRepository{}
}
