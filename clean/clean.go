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
		localDeletedBranches, err := cleanBranches(newLocalGitRepo(repo), shortcutClient, isDryRun)
		if err != nil {
			slog.Warn("Error while cleaning local branches, not attempting to clean remote branches", "error", err)
			return localDeletedBranches, err
		}
		deletedBranches = localDeletedBranches
	}
	if opts.IncludeRemoteBranches {
		remoteDeletedBranches, err := cleanBranches(newRemoteGitRepo(repo), shortcutClient, isDryRun)
		if err != nil {
			return remoteDeletedBranches, err
		}
		deletedBranches = append(deletedBranches, remoteDeletedBranches...)
	}
	return deletedBranches, nil
}

func cleanBranches(g gitRepo, shortcutClient shortcut.Client, isDryRun bool) ([]string, error) {
	branchNames := g.getBranchNames()
	deletedBranches := []string{}
	var lastError error

	for _, branchName := range branchNames {
		slog.Debug("Considering branch", "branch", branchName)

		if g.shouldSkipBranch(branchName) {
			slog.Info("Skipping branch, as it's the current branch", "branch", branchName)
			continue
		}
		story, err := usecases.GetStoryByBranchName(branchName, shortcutClient)
		if err != nil {
			slog.Warn("Error while fetching story", "error", err, "branch", branchName)
			lastError = err
			continue
		}
		if story == nil {
			continue
		}

		if story.Completed {
			if isDryRun {
				deletedBranches = append(deletedBranches, branchName)
			} else {
				err = g.deleteBranch(branchName)
				if err != nil {
					slog.Warn("Error while deleting branch", "error", err, "branch", branchName)
					lastError = err
					continue
				} else {
					deletedBranches = append(deletedBranches, branchName)
				}
			}
		}
	}
	return deletedBranches, lastError
}

type gitRepo interface {
	shouldSkipBranch(branchName string) bool
	getBranchNames() []string
	deleteBranch(branchName string) error
}
type localGitRepo struct{ repo git.Repository }

func newLocalGitRepo(repo git.Repository) localGitRepo { return localGitRepo{repo: repo} }
func (l localGitRepo) shouldSkipBranch(branchName string) bool {
	return branchName == l.repo.GetCurrentBranchName()
}
func (l localGitRepo) getBranchNames() []string { return l.repo.GetLocalBranchNames() }
func (l localGitRepo) deleteBranch(branchName string) error {
	return l.repo.DeleteLocalBranch(branchName)
}

type remoteGitRepo struct{ repo git.Repository }

func newRemoteGitRepo(repo git.Repository) remoteGitRepo        { return remoteGitRepo{repo: repo} }
func (r remoteGitRepo) shouldSkipBranch(branchName string) bool { return false }
func (r remoteGitRepo) getBranchNames() []string                { return r.repo.GetRemoteBranchNames() }
func (r remoteGitRepo) deleteBranch(branchName string) error {
	return r.repo.DeleteRemoteBranch(branchName)
}
