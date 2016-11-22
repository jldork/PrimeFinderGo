package PrimeGo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPrimeGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrimeGo Suite")
}
