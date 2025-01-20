package clean

import (
	"log/slog"

	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/shortcut"
	"github.com/carpeliam/gitshorty/usecases"
)

func CleanLocalBranches(repo git.Repository, shortcutClient shortcut.Client) ([]string, error) {
	currentBranch := repo.GetCurrentBranchName()
	localBranchNames := repo.GetLocalBranchNames()
	deletedBranches := []string{}

	for _, branchName := range localBranchNames {
		if branchName == currentBranch {
			slog.Info("Skipping branch, as it's the current branch", "branch", branchName)
			continue
		}
		story, error := usecases.GetStoryByBranchName(branchName, shortcutClient)
		if error != nil {
			return deletedBranches, error
		}
		if story == nil {
			continue
		}

		if story.Completed {
			error = repo.DeleteBranch(branchName)
			if error != nil {
				return deletedBranches, error
			} else {
				deletedBranches = append(deletedBranches, branchName)
			}
		}
	}
	return deletedBranches, nil
}
