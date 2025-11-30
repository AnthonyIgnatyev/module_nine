package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	testSizes := []int{
		0, -1, -100,
		1, 5, 10, 42, 100, 1000,
		100_000, 10_000_000,
	}

	for _, size := range testSizes {
		result := generateRandomElements(size)
		if size <= 0 {
			assert.Equal(t, []int{}, result, "Для size <= 0 должен возвращаться пустой срез")
		} else {
			assert.Len(t, result, size, "Длина среза должна быть равна size")
		}
	}
}

func TestMaximum(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{42}, 42},
		{[]int{10, 5, 20}, 20},
		{[]int{3, 1, 4, 1, 5, 9, 2}, 9},
		{[]int{0, 0, 0}, 0},
		{[]int{7, 3, 7, 2}, 7},
	}

	for _, tc := range testCases {
		result := maximum(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}
