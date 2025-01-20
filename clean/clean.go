package clean

import (
	"log/slog"

	"github.com/carpeliam/gitshorty/git"
	"github.com/carpeliam/gitshorty/shortcut"
	"github.com/carpeliam/gitshorty/usecases"
)

type CleanOpts struct {
	DryRun                bool
	IncludeLocalBranches  bool
	IncludeRemoteBranches bool
}

func CleanBranches(repo git.Repository, shortcutClient shortcut.Client, opts CleanOpts) ([]string, error) {
	isDryRun := opts.DryRun
	deletedBranches := []string{}
	if opts.IncludeLocalBranches {
		localDeletedBranches, err := cleanLocalBranches(repo, shortcutClient, isDryRun)
		if err != nil {
			return localDeletedBranches, err
		}
		deletedBranches = localDeletedBranches
	}
	if opts.IncludeRemoteBranches {
		remoteDeletedBranches, err := cleanRemoteBranches(repo, shortcutClient, isDryRun)
		if err != nil {
			return remoteDeletedBranches, err
		}
		deletedBranches = append(deletedBranches, remoteDeletedBranches...)
	}
	return deletedBranches, nil
}

func cleanLocalBranches(repo git.Repository, shortcutClient shortcut.Client, isDryRun bool) ([]string, error) {
	currentBranch := repo.GetCurrentBranchName()
	localBranchNames := repo.GetLocalBranchNames()
	deletedBranches := []string{}

	for _, branchName := range localBranchNames {
		if branchName == currentBranch {
			slog.Info("Skipping branch, as it's the current branch", "branch", branchName)
			continue
		}
		story, err := usecases.GetStoryByBranchName(branchName, shortcutClient)
		if err != nil {
			return deletedBranches, err
		}
		if story == nil {
			continue
		}

		if story.Completed {
			if isDryRun {
				deletedBranches = append(deletedBranches, branchName)
			} else {
				err = repo.DeleteLocalBranch(branchName)
				if err != nil {
					return deletedBranches, err
				} else {
					deletedBranches = append(deletedBranches, branchName)
				}
			}
		}
	}
	return deletedBranches, nil
}

func cleanRemoteBranches(repo git.Repository, shortcutClient shortcut.Client, isDryRun bool) ([]string, error) {
	remoteBranchNames := repo.GetRemoteBranchNames()
	deletedBranches := []string{}

	for _, branchName := range remoteBranchNames {
		story, err := usecases.GetStoryByBranchName(branchName, shortcutClient)
		if err != nil {
			return deletedBranches, err
		}
		if story == nil {
			continue
		}

		if story.Completed {
			if isDryRun {
				deletedBranches = append(deletedBranches, branchName)
			} else {
				err = repo.DeleteRemoteBranch(branchName)
				if err != nil {
					return deletedBranches, err
				} else {
					deletedBranches = append(deletedBranches, branchName)
				}
			}
		}
	}

	return deletedBranches, nil
}
