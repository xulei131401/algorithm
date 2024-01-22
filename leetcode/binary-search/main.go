package main

import "log"

/*
*
https://leetcode.cn/problems/binary-search/description/
https://leetcode.com/problems/binary-search/description/
*/
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	log.Println("res Index:", search(nums, target))
}

// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func search(nums []int, target int) int {
	// 数组长度为空，返回-1
	if len(nums) == 0 {
		return -1
	}

	// 初始化判断区间的2个端点位置，mid通过这2个端点可以计算得出
	l, r := 0, len(nums)-1

	// l 与 r 是可能重合的
	for l <= r {
		// 计算mid
		mid := l + (r-l)>>1

		// 找到了值，直接返回
		if nums[mid] == target {
			return mid
		}

		// nums[mid] > target, 因为整体有序，说明target在 l~mid区间范围，需要移动r到mid-1位置，缩小一半范围
		if nums[mid] > target {
			// 因为当前判断的就是mid,所以需要移动到mid-1
			r = mid - 1
		} else if nums[mid] < target { // 相反
			l = mid + 1
		}
	}

	// 说明没找到结果，根据题意，返回-1
	return -1
}
