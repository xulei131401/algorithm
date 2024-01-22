package main

import "log"

/*
*
https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
*/
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	log.Println("在排序数组中查找元素的第一个和最后一个位置:", searchRange(nums, target))
}

// 时间复杂度： O(log n)
// 空间复杂度：O(1)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 先找最左边的target
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if nums[low] != target {
		return []int{-1, -1}
	}

	// 找右边比target大的第一个数
	left := low
	high = len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	right := low - 1
	return []int{left, right}
}
