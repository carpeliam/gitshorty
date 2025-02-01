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

func GetBranchesForDeliveredStories(repo git.Repository, shortcutClient shortcut.Client, options GetBranchesForDeliveredStoriesOptions) ([]string, error) {
	branches := []string{}
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
func getBranchesForDeliveredStories(iterator git.BranchIterator, shortcutClient shortcut.Client) ([]string, error) {
	branchNames := iterator.GetBranchNames()

	branches := []string{}
	for _, branchName := range branchNames {
		slog.Debug("Considering branch", "branch", branchName)
		story, err := GetStoryByBranchName(branchName, shortcutClient)
		if err != nil {
			slog.Warn("Error while fetching story", "error", err, "branch", branchName)
			return nil, err
		}
		if story == nil {
			continue
		}
		if story.Completed {
			branches = append(branches, branchName)
		}
	}
	return branches, nil
}
