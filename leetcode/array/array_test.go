// @author hongjun500
// @date 2023/8/29 13:48
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestBinarySearch 测试二分查找
func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	target := 3
	index := binarySearch(nums, target)
	assert.Equal(t, 2, index)
	t.Log(index)

	index = binarySearch2(nums, target)
	assert.Equal(t, 2, index)
	t.Log(index)
}

// TestRemoveElement 测试移除元素 快慢指针
func TestRemoveElement(t *testing.T) {
	nums := []int{1, 2, 3, 2, 5}
	val := 2
	length := removeElement(nums, val)
	assert.Equal(t, 3, length)
	t.Log(length)

	nums = []int{1, 2, 3, 2, 5}
	val = 3
	length = removeElement2(nums, val)
	assert.Equal(t, 4, length)
	t.Log(length)
}

// TestSortedSquares 测试有序数组的平方 双指针
func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	squares := sortedSquares(nums)
	assert.Equal(t, []int{0, 1, 9, 16, 100}, squares)
	t.Log(squares)
}

// TestMinSubArrayLen 测试长度最小的子数组 滑动窗口
func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	length := minSubArrayLen(target, nums)
	assert.Equal(t, 2, length)
	t.Log(length)

	nums = []int{1, 4, 4}
	target = 4
	length = minSubArrayLen(target, nums)
	assert.Equal(t, 1, length)
	t.Log(length)

	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	target = 11
	length = minSubArrayLen(target, nums)
	assert.Equal(t, 0, length)
	t.Log(length)
}

func TestGenerateMatrix(t *testing.T) {
	n := 3
	matrix := generateMatrix(n)
	assert.Equal(t, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, matrix)
	t.Log(matrix)

	n = 1
	matrix = generateMatrix(n)
	assert.Equal(t, [][]int{{1}}, matrix)
	t.Log(matrix)
}
