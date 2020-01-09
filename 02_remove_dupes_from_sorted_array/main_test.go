package remove_dupes_from_sorted

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Tests", func() {
	It("Basic tests", func() {
		Describe("Should Return Length After Elements Removed - Case 1", func() {
			Expect(RemoveDuplicates([]int{1, 1, 2})).To(Equal(2))
		})
		Describe("Should Return Length After Elements Removed - Case 2", func() {
			Expect(RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})).To(Equal(5))
		})
	})
	It("Edge Cases", func() {
		Describe("Empty Array", func() {
			Expect(RemoveDuplicates([]int{})).To(Equal(0))
		})
		Describe("All the Same", func() {
			Expect(RemoveDuplicates([]int{0, 0, 0, 0, 0, 0, 0})).To(Equal(1))
		})
	})
})
