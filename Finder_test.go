package PrimeGo_test

import (
	. "side-projects/PrimeGo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Finder", func() {
	It("Should be able to check if prime", func() {
		Expect(IsPrime(1)).To(Equal(false))
		Expect(IsPrime(2)).To(Equal(true))
		Expect(IsPrime(4)).To(Equal(false))
		Expect(IsPrime(29)).To(Equal(true))
		Expect(IsPrime(30)).To(Equal(false))
	})
	It("Should find the nth prime", func() {
		Expect(FindPrime(1)).To(Equal(2))
		Expect(FindPrime(10)).To(Equal(29))
	})
})
