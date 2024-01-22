package main

import "log"

/*
*
https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/submissions/
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/submissions/
*/
func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 3
	log.Println("搜索旋转排序数组 II:", search(nums, target))
}

// 时间复杂度：O(n),其中 nnn 是数组的长度, 最坏情况下所有元素均相等且不为target
// 空间复杂度：O(1)
func search(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}

	if length == 1 {
		return nums[0] == target
	}

	l, r := 0, length-1
	for l <= r {
		mid := l + (r-l)>>2
		if nums[mid] == target {
			return true
		}

		// 因为有重复元素，此时无法判断哪边有序，只能缩小一个范围继续查找
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
		} else if nums[l] <= nums[mid] {
			// 因为有重复元素，所以需要考虑相等的情况
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}
