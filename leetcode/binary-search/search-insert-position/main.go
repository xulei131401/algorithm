package main

import (
	"log"
	"sort"
)

/*
*
https://leetcode.cn/problems/search-insert-position/description/
https://leetcode.com/problems/search-insert-position/description/
*/
func main() {
	nums := []int{-1, 0, 3, 5, 8, 9, 12}
	target := 13

	log.Println("[手写二分]值:", searchInsert(nums, target))
	log.Println("[内置二分]值:", searchInsert1(nums, target))
}

// 时间复杂度：O(logN)，其中 n 是数组的长度。
// 空间复杂度：O(1)。
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)>>1 + low

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		}
	}

	// 如果target不存在，low肯定超出了数组maxIndex，返回这个位置正好符合题意
	return low
}

// 时间复杂度：O(logN)，其中 n 是数组的长度。
// 空间复杂度：O(1)。
func searchInsert1(nums []int, target int) int {
	index := sort.SearchInts(nums, target)
	if len(nums) == index {
		// 说明没有target
		return len(nums)
	}

	return index
}
