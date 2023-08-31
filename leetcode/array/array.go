// @author hongjun500
// @date 2023/8/28 14:44
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 数组

package array

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// binarySearch 二分查找
func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right { // 注意这里是小于等于
		mid := left + (right-left)/2 // 防止溢出
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target { // 注意这里是大于
			right = mid - 1 // 注意这里是减一
		} else if nums[mid] < target { // 注意这里是小于
			left = mid + 1 // 注意这里是加一
		}
	}
	return -1
}

// binarySearch2 二分查找第二种写法
func binarySearch2(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right { // 注意这里是小于
		mid := left + (right-left)/2 // 防止溢出
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target { // 注意这里是大于
			right = mid // 注意这里不减一
		} else if nums[mid] < target { // 注意这里是小于
			left = mid + 1 // 注意这里是加一
		}
	}
	return -1
}

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
// removeElement 移除元素 快慢指针
func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}
	return left
}

// removeElement2 移除元素第二种写法 相向双指针
func removeElement2(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] == val {
			for left <= right && nums[right] == val {
				right--
			}
			if left <= right {
				nums[left], nums[right] = nums[right], nums[left]
				right--
			}
		} else {
			left++
		}
	}
	return left
}

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
// sortedSquares 有序数组的平方 双指针
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	// 左右指针
	l, r := 0, n-1
	// 从后往前填充
	for pos := n - 1; pos >= 0; pos-- {
		if lv, rv := nums[l]*nums[l], nums[r]*nums[r]; lv > rv {
			// 左指针的平方大于右指针的平方 则填充左指针的平方到最后
			ans[pos] = lv
			l++
		} else {
			ans[pos] = rv
			r--
		}
	}
	return ans
}

// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
// minSubArrayLen 长度最小的子数组 滑动窗口
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	ans := n + 1 // 最小长度
	l, r := 0, 0 // 左右指针
	sum := 0
	for r < n {
		sum += nums[r] // 右指针向右移动
		for sum >= s {
			ans = r - l + 1 // 更新最小长度
			sum -= nums[l]  // 减去最左边的值从而缩小子数组的长度
			l++             // 左指针向右移动
		}
		r++
	}
	if ans == n+1 { // 如果没有找到符合条件的子数组 则返回0
		return 0
	}
	return ans
}

// 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
// generateMatrix 螺旋矩阵
func generateMatrix(n int) [][]int {
	// 初始化矩阵
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n) // 初始化每一行
	}
	// 定义上下左右四个边界
	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1 // 从 1 开始填充
	for top <= bottom && left <= right {
		// 从左到右
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		// 从上到下
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		// 从右到左
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		// 从下到上
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
