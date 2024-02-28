package main

import "log"

/*
*
https://leetcode.cn/problems/search-in-rotated-sorted-array/
https://leetcode.com/problems/search-in-rotated-sorted-array/
*/
func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	log.Println("搜索旋转排序数组-两次二分:", search(nums, target))

	log.Println("搜索旋转排序数组-一次二分:", search2(nums, target))
}

// 时间复杂度：O(log n)
// 空间复杂度：O(1)
// 两次二分，分别对左部分区间和右部分区间进行二分查找
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	l, r := 0, length-1
	mid := l + (r-l)>>1
	// 如果正好找到了，提前返回
	if nums[mid] == target {
		return mid
	}

	// 寻找旋转之后数组的峰值，然后获取到2个升序数组的范围
	peakIndex := mid
	if nums[0] <= nums[mid] { // 左边部分有序
		for peakIndex+1 <= r && nums[peakIndex] < nums[peakIndex+1] {
			peakIndex += 1
		}
	} else {
		for peakIndex-1 >= l && nums[peakIndex] > nums[peakIndex-1] {
			peakIndex -= 1
		}
	}

	// 二分查找模板
	binarySearch := func(elements []int, target int) int {
		// 数组长度为空，返回-1
		if len(elements) == 0 {
			return -1
		}

		// 初始化判断区间的2个端点位置，mid通过这2个端点可以计算得出
		l, r := 0, len(elements)-1

		// l 与 r 是可能重合的
		for l <= r {
			// 计算mid
			mid := l + (r-l)>>1

			// 找到了值，直接返回
			if elements[mid] == target {
				return mid
			}

			// nums[mid] > target, 因为整体有序，说明target在 l~mid区间范围，需要移动r到mid-1位置，缩小一半范围
			if elements[mid] > target {
				// 因为当前判断的就是mid,所以需要移动到mid-1
				r = mid - 1
			} else if elements[mid] < target { // 相反
				l = mid + 1
			}
		}

		// 说明没找到结果，根据题意，返回-1
		return -1
	}

	// 先查找0~peakIndex
	res := binarySearch(nums[:peakIndex+1], target)
	if res != -1 {
		return res
	}

	// 再查找peakIndex+1~r
	res = binarySearch(nums[peakIndex+1:], target)
	if res == -1 {
		return -1
	}

	// 如果target是在右边的升序数组中，返回原来数组中的具体位置
	return len(nums[:peakIndex+1]) + res
}

// 时间复杂度：O(log n)
// 空间复杂度：O(1)
// 一次二分
func search2(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	l, r := 0, length-1
	if nums[l] == target {
		return l
	}

	if nums[r] == target {
		return r
	}

	// 把l,r理解成始终代表有序的区间部分就行
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}

		// 0-mid升序
		if nums[0] < nums[mid] {
			if nums[0] < target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // mid-r 升序
			if target < nums[r] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
