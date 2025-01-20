package clean_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestClean(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Clean Suite")
}
