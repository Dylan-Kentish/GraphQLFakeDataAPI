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