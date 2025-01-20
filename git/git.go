package git

import (
	"os/exec"
	"strings"
)

type Repository interface {
	GetCurrentBranchName() string
	GetLocalBranchNames() []string
	DeleteBranch(branchName string) error
}

type GitRepository struct{}

func (repository GitRepository) GetCurrentBranchName() string {
	output, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	return strings.TrimSpace(string(output))
}

func (repository GitRepository) GetLocalBranchNames() []string {
	output, _ := exec.Command("git", "--no-pager", "branch", "--format='%(refname:short)'").Output()
	return strings.Split(strings.TrimSpace(string(output)), "\n")
}

func (repository GitRepository) DeleteBranch(branchName string) error {
	_, err := exec.Command("git", "branch", "-D", branchName).Output()
	return err
}

func NewRepository() Repository {
	return GitRepository{}
}
