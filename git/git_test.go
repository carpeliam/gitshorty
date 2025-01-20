package git_test

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/carpeliam/gitshorty/git"
)

var _ = Describe("Git", func() {
	path := os.Getenv("PATH")

	BeforeEach(func() {
		fixtures_path, _ := filepath.Abs("../support/fixtures")
		os.Setenv("PATH", fmt.Sprintf("%s:%s", fixtures_path, path))
	})
	AfterEach(func() {
		os.Setenv("PATH", path)
	})

	It("should know the current branch", func() {
		os.Setenv("GITSHORTY_TEST_CURRENT_BRANCH", "current-branch-sc-123")
		branchName := git.NewRepository().GetCurrentBranchName()
		Expect(branchName).To(Equal("current-branch-sc-123"))
	})

	It("can list local branches", func() {
		os.Setenv("GITSHORTY_TEST_CURRENT_BRANCH", "main")
		os.Setenv("GITSHORTY_TEST_LOCAL_BRANCHES", "main:current-branch-sc-123")
		branchNames := git.NewRepository().GetLocalBranchNames()
		Expect(branchNames).To(Equal([]string{"main", "current-branch-sc-123"}))
	})

	It("can list remote branches", func() {
		os.Setenv("GITSHORTY_TEST_CURRENT_BRANCH", "main")
		os.Setenv("GITSHORTY_TEST_REMOTE_BRANCHES", "origin/main:upstream/current-branch-sc-123")
		branchNames := git.NewRepository().GetRemoteBranchNames()
		Expect(branchNames).To(Equal([]string{"origin/main", "upstream/current-branch-sc-123"}))
	})

	It("can delete a local branch", func() {
		os.Setenv("GITSHORTY_TEST_CURRENT_BRANCH", "main")
		os.Setenv("GITSHORTY_TEST_LOCAL_BRANCHES", "main:current-branch-sc-123")
		err := git.NewRepository().DeleteLocalBranch("current-branch-sc-123")
		Expect(err).To(BeNil())
	})

	It("fails to delete a non-existent local branch", func() {
		os.Setenv("GITSHORTY_TEST_CURRENT_BRANCH", "main")
		os.Setenv("GITSHORTY_TEST_LOCAL_BRANCHES", "main:current-branch-sc-123")
		err := git.NewRepository().DeleteLocalBranch("current-branch-sc-456")
		Expect(err).NotTo(BeNil())
	})

	It("can delete a remote branch", func() {
		os.Setenv("GITSHORTY_TEST_CURRENT_BRANCH", "main")
		os.Setenv("GITSHORTY_TEST_REMOTE_BRANCHES", "origin/main:upstream/current-branch-sc-123")
		err := git.NewRepository().DeleteRemoteBranch("upstream/current-branch-sc-123")
		Expect(err).To(BeNil())
	})

	It("can delete a remote branch with a slash in its name", func() {
		os.Setenv("GITSHORTY_TEST_CURRENT_BRANCH", "main")
		os.Setenv("GITSHORTY_TEST_REMOTE_BRANCHES", "origin/feature/sc-123")
		err := git.NewRepository().DeleteRemoteBranch("origin/feature/sc-123")
		Expect(err).To(BeNil())
	})

	It("fails to delete a non-existent remote branch", func() {
		os.Setenv("GITSHORTY_TEST_CURRENT_BRANCH", "main")
		os.Setenv("GITSHORTY_TEST_REMOTE_BRANCHES", "origin/main:upstream/current-branch-sc-123")
		err := git.NewRepository().DeleteRemoteBranch("upstream/current-branch-sc-456")
		Expect(err).NotTo(BeNil())
	})
})
