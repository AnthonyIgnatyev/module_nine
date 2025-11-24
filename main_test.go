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
	testCases := [][]int{
		{},
		{42},
		{5},
		{1, 2, 3},
		{10, 5, 8, 3, 9, 1},
		{10, 5, 1, 20},
		{0, 0, 0},
		{100, 200, 300, 250, 150},
	}

	for _, data := range testCases {
		expected := 0
		if len(data) == 0 {
			expected = 0
		} else if len(data) == 1 {
			expected = data[0]
		} else {
			expected = data[0]
			for _, v := range data {
				if v > expected {
					expected = v
				}
			}
		}

		actual := maximum(data)
		assert.Equal(t, expected, actual, "Для данных %v максимум должен быть %d", data, expected)
	}
}
