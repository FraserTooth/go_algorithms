package remove_dupes_from_sorted

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Remove Duplicates in Place Suite")
}

var _ = Describe("Tests", func() {
	It("Basic tests", func() {
		Describe("Should Return Length After Elements Removed - Case 1", func() {
			var input = []int{1, 1, 2}
			var output = []int{1, 2}
			var newLen = RemoveDuplicates(input)
			Expect(newLen).To(Equal(2))
			Expect(input[:newLen]).To(Equal(output))
		})
		Describe("Should Return Length After Elements Removed - Case 2", func() {
			var input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
			var output = []int{0, 1, 2, 3, 4}
			var newLen = RemoveDuplicates(input)
			Expect(newLen).To(Equal(5))
			Expect(input[:newLen]).To(Equal(output))
		})
	})
	It("Edge Cases", func() {
		Describe("Empty Array", func() {
			var input = []int{}
			Expect(RemoveDuplicates(input)).To(Equal(0))
		})
		Describe("All the Same", func() {
			var input = []int{0, 0, 0, 0, 0, 0, 0}
			Expect(RemoveDuplicates(input)).To(Equal(1))
		})
	})
})
