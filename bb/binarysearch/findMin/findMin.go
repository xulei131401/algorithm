package main

import "log"

/*
*
题目：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/

寻找旋转排序数组中的最小值

已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

示例 1：

输入：nums = [3,4,5,1,2]
输出：1
解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
示例 2：

输入：nums = [4,5,6,7,0,1,2]
输出：0
解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。
示例 3：

输入：nums = [11,13,15,17]
输出：11
解释：原数组为 [11,13,15,17] ，旋转 4 次得到输入数组。

提示：

n == nums.length
1 <= n <= 5000
-5000 <= nums[i] <= 5000
nums 中的所有整数 互不相同
nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转

注意：
1.旋转数组的意思是原来数组是有序的，经过旋转变成现在的样子
2.旋转的含义，数组末尾的元素移动到数组头部，依次向后移动
3.数组元素没有相同的
*/
func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	log.Println("寻找旋转排序数组中的最小值-二分查找:", findMin(nums))
	log.Println("寻找旋转排序数组中的最大值-二分查找:", findMax(nums))
}

// findMin O(logn) O(1)
func findMin(nums []int) int {
	// 数组有序或者部分有序，都可以使用二分法进行遍历
	low, high := 0, len(nums)-1
	for low < high {
		pivotIndex := low + (high-low)>>1
		// 最小值在左侧，收缩high
		if nums[pivotIndex] < nums[high] {
			high = pivotIndex
		} else {
			// 最小值在右侧，收缩low
			low = pivotIndex + 1
		}
	}

	return nums[low]
}

// findMax O(logn) O(1)
func findMax(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivotIndex := low + (high-low)>>1
		if nums[pivotIndex] > nums[high] {
			low = pivotIndex
		} else {
			high = pivotIndex - 1
		}
	}

	return nums[low]
}
