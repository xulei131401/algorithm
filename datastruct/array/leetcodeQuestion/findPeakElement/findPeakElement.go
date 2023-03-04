package main

import (
	"log"
)

/**
寻找峰值
题目：https://leetcode.cn/problems/find-peak-element/

峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞ 。

你必须实现时间复杂度为 O(log n) 的算法来解决此问题。



示例 1：

输入：nums = [1,2,3,1]
输出：2
解释：3 是峰值元素，你的函数应该返回其索引 2。
示例 2：

输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5
解释：你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。


提示：

1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
对于所有有效的 i 都有 nums[i] != nums[i + 1]

注意：
1.测试用例数组中nums[i] != nums[i+1]
*/

func main() {
	nums := []int{2, 1}
	log.Println("寻找峰值(直接遍历1)索引:", findPeakElement(nums))
	log.Println("寻找峰值(直接遍历2)索引:", findPeakElement2(nums))
	log.Println("寻找峰值(二分查找)索引:", findPeakElement3(nums))
}

// findPeakElement O(n) O(1)
func findPeakElement(nums []int) int {
	length := len(nums)
	// 根据题意，一个元素的时候，峰值就是这个元素
	if length == 1 {
		return 0
	}

	right := 0
	for right < length {
		// 找到第一个峰值，则直接退出
		if right > 0 && nums[right] <= nums[right-1] {
			break
		}

		right++
	}

	return right - 1
}

// findPeakElement2 O(n) O(1)
func findPeakElement2(nums []int) int {
	idx := 0
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}

	return idx
}

// findPeakElement3 O(logN) O(1)
func findPeakElement3(nums []int) int {
	length := len(nums)
	left, right := 0, length-1 // 初始位置
	// 循环条件，left必须小于等于right
	for left <= right {
		mid := (right + left) >> 1

		// 峰值条件，mid >mid-1,mid>mid+1,注意边界值
		if mid > 0 && mid < length-1 && nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			return mid
		}

		if mid < length-1 && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
