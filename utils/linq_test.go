package utils

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transform", func() {
	It("Transforms each element", func() {
		count := 10
		slice := make([]int, count)
		expected := make([]string, count)

		for i := 0; i < count; i++ {
			slice[i] = i
			expected[i] = fmt.Sprint(i)
		}

		result := Transform(slice, func(i int) string { return fmt.Sprint(i) })

		Expect(result).To(ContainElements(expected))
	})
})

var _ = Describe("TransformValues", func() {
	It("Transforms each value", func() {
		count := 10
		m := make(map[int]int, count)
		expected := make([]string, count)

		for i := 0; i < count; i++ {
			m[i] = i
			expected[i] = fmt.Sprint(i)
		}

		result := TransformValues(m, func(i int) string { return fmt.Sprint(i) })

		Expect(result).To(ContainElements(expected))
	})
})

var _ = Describe("Where", func() {
	It("Filters slice", func() {
		count := 10
		slice := make([]int, count)

		for i := 0; i < count; i++ {
			slice[i] = i
		}

		result := Where(slice, func(i int) bool { return i < 5 })

		Expect(result).To(HaveLen(5))
		Expect(result).To(ContainElements(BeNumerically("<", 5)))
	})
})

var _ = Describe("ValuesWhere", func() {
	It("Filters values", func() {
		count := 10
		m := make(map[int]int, count)

		for i := 0; i < count; i++ {
			m[i] = i
		}

		result := ValuesWhere(m, func(i int) bool { return i < 5 })

		Expect(result).To(HaveLen(5))
		Expect(result).To(ContainElements(BeNumerically("<", 5)))
	})
})

var _ = Describe("OrderedValues", func() {
	It("returns the values in order", func() {
		count := 10
		m := make(map[int]int, count)
		expected := make([]int, count)

		for i := count - 1; i >= 0; i-- {
			m[i] = i
		}
		for i := 0; i < count; i++ {
			expected[i] = i
		}

		result := OrderedValues(m)

		Expect(result).To(Equal(expected))
	})
})

var _ = Describe("Single", func() {
	slice := make([]int, 5)

	BeforeEach(func() {
		for i := 0; i < len(slice); i++ {
			slice[i] = 0
		}
	})

	When("no matches", func() {
		It("returns an error", func() {
			match, err := Single(slice, func(item int) bool {
				return item == 1
			})

			Expect(match).To(BeNil())
			Expect(err).To(MatchError("no items match the condition"))
		})
	})

	When("one matches", func() {
		BeforeEach(func() {
			slice[2] = 1
		})

		It("returns the match", func() {
			match, err := Single(slice, func(item int) bool {
				return item == 1
			})

			Expect(*match).To(Equal(1))
			Expect(err).To(BeNil())
		})
	})

	When("multiple matches", func() {
		It("returns an error", func() {
			match, err := Single(slice, func(item int) bool {
				return item == 0
			})

			Expect(match).To(BeNil())
			Expect(err).To(MatchError("multiple items match the condition"))
		})
	})
})
