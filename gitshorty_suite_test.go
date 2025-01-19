package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGitshorty(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gitshorty Suite")
}
