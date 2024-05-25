package lc

import (
	"math"
)

// 返回最长的递增子序列长度
func LengthOfLIS(nums []int) int {
	m := make(map[int]int)
	maxLen := 1
	for i := range nums {
		len := l(nums, i, m)
		if len > maxLen {
			maxLen = len
		}
	}
	return maxLen
}

// 返回从下标i开始，最长的递增子序列长度
func l(nums []int, i int, m map[int]int) int {
	if rs, ok := m[i]; ok {
		// fmt.Println("避免重复遍历：", i)
		return rs
	}
	if i == len(nums)-1 { // 递归终止条件
		return 1
	}
	maxLen := 1
	for j := i + 1; j < len(nums); j++ {
		if nums[j] > nums[i] {
			maxLen = int(math.Max(float64(l(nums, j, m)+1), float64(maxLen)))
		}
	}
	m[i] = maxLen
	return maxLen
}

// 动态规划
func LengthOfLISByDp(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	rs := 1
	for i := len(nums) - 2; i >= 0; i-- { // 求从下标i开始，最长的递增子序列长度
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				max := dp[j] + 1
				if max > dp[i] {
					dp[i] = max
					if max > rs {
						rs = max
					}
				}
			}
		}
	}
	return rs
}
