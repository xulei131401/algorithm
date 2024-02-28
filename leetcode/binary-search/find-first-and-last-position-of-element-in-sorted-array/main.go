package main

import (
	"log"
	"sort"
)

/*
*
https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
*/
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	log.Println("[手写二分]在排序数组中查找元素的第一个和最后一个位置:", searchRange(nums, target))

	nums1 := []int{5, 7, 7, 8, 8, 10}
	target1 := 8
	log.Println("[内置二分]在排序数组中查找元素的第一个和最后一个位置:", searchRange1(nums1, target1))

	nums2 := []int{5, 7, 7, 8, 8, 10}
	target2 := 8
	log.Println("[固定模板]在排序数组中查找元素的第一个和最后一个位置:", searchRange2(nums2, target2))
}

// 时间复杂度： O(log n)
// 空间复杂度：O(1)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 找左边区域target的最小位置
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	// 说明target不存在
	if nums[low] != target {
		return []int{-1, -1}
	}

	// 找右边区域target最大位置
	left := low
	high = len(nums) - 1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	right := low - 1 // low-1正好是正确的位置
	return []int{left, right}
}

func searchRange1(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}

	// sort.SearchInts 固定返回target的正确位置，但这个位置不一定合法
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// 找左边界
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			high = mid - 1 // 右边往左边靠
		}
	}

	// 判断索引是不是越界了
	if low < 0 || low >= len(nums) {
		return []int{-1, -1}
	}

	// 说明target不存在
	if nums[low] != target {
		return []int{-1, -1}
	}

	// 找右边界
	left := low
	high = len(nums) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			low = mid + 1 // 左边往右边靠
		}
	}

	// 超出边界
	if high < 0 || high >= len(nums) {
		return []int{left, left}
	}

	if nums[high] == target {
		return []int{left, high}
	}

	return []int{left, left}
}
