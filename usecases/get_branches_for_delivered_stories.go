package usecases

import (
	"log/slog"

	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/shortcut"
)

type GetBranchesForDeliveredStoriesOptions struct {
	IncludeLocal  bool
	IncludeRemote bool
}

func GetBranchesForDeliveredStories(repo git.Repository, shortcutClient shortcut.Client, options GetBranchesForDeliveredStoriesOptions) ([]git.Branch, error) {
	branches := []git.Branch{}
	if options.IncludeLocal {
		slog.Debug("Getting local branches")
		iterator := git.NewLocalBranchIterator(repo)
		localBranches, err := getBranchesForDeliveredStories(iterator, shortcutClient)
		if err != nil {
			slog.Warn("Error while determining local branches for delivered stories; skipping remote branches", "error", err)
			return localBranches, err
		}
		branches = localBranches
	}
	if options.IncludeRemote {
		slog.Debug("Getting remote branches")
		iterator := git.NewRemoteBranchIterator(repo)
		remoteBranches, err := getBranchesForDeliveredStories(iterator, shortcutClient)
		if remoteBranches != nil {
			branches = append(branches, remoteBranches...)
		}
		if err != nil {
			return branches, err
		}
	}
	return branches, nil
}
func getBranchesForDeliveredStories(iterator git.BranchIterator, shortcutClient shortcut.Client) ([]git.Branch, error) {
	branches := iterator.GetBranches()

	deliveredBranches := []git.Branch{}
	for _, branch := range branches {
		slog.Debug("Considering branch", "branch", branch.Name)
		story, err := GetStoryByBranchName(branch.Name, shortcutClient)
		if err != nil {
			slog.Warn("Error while fetching story", "error", err, "branch", branch.Name)
			return nil, err
		}
		if story == nil {
			continue
		}
		if story.Completed {
			deliveredBranches = append(deliveredBranches, branch)
		}
	}
	return deliveredBranches, nil
}
