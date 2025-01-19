package browse

import (
	"os/exec"
	"strings"
)

type Repository interface {
	GetCurrentBranchName() string
}

type GitRepository struct{}

func (repository GitRepository) GetCurrentBranchName() string {
	output, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	return strings.TrimSpace(string(output))
}

func NewRepository() GitRepository {
	return GitRepository{}
}
