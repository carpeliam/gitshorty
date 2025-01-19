package browse_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBrowse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Browse Suite")
}
