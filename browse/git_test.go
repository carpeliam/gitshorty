package browse_test

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/carpeliam/gitshorty/browse"
)

var _ = Describe("Git", func() {
	path := os.Getenv("PATH")

	BeforeEach(func() {
		fixtures_path, _ := filepath.Abs("./fixtures")
		os.Setenv("PATH", fmt.Sprintf("%s:%s", fixtures_path, path))
	})
	AfterEach(func() {
		os.Setenv("PATH", path)
		os.Unsetenv("GITSHORTY_TEST_CURRENT_BRANCH")
	})

	It("should know the current branch", func() {
		os.Setenv("GITSHORTY_TEST_CURRENT_BRANCH", "current-branch-sc-123")
		branchName := browse.NewRepository().GetCurrentBranchName()
		Expect(branchName).To(Equal("current-branch-sc-123"))
	})

})
