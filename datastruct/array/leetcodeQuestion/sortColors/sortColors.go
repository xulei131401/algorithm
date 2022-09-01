package main

import "log"

/**
题目：https://leetcode.cn/problems/sort-colors/
颜色分类

给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

必须在不使用库的sort函数的情况下解决这个问题。

示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]

提示：

n == nums.length
1 <= n <= 300
nums[i] 为 0、1 或 2

进阶：

你可以不使用代码库中的排序函数来解决这道题吗？
你能想出一个仅使用常数空间的一趟扫描算法吗？


*/
func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	log.Println("颜色分类-双指针:", nums)

	nums1 := []int{2, 0, 2, 1, 1, 0}
	sortColors1(nums1)
	log.Println("颜色分类-单指针:", nums1)
}

// sortColors 双指针 O(n)O(1) 一趟排序
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		// 不断的把2交换到数组尾部，并且移动p2位置
		for ; i <= p2 && nums[i] == 2; p2-- {
			if nums[p2] != 2 {
				nums[i], nums[p2] = nums[p2], nums[i]
			}
		}

		// 0 的话置换到数组首部，那么中间就剩下1了
		if nums[i] == 0 {
			if nums[p0] != 0 {
				nums[i], nums[p0] = nums[p0], nums[i]
			}

			p0++
		}
	}
}

// sortColors1 单指针 O(n) O(1) 两趟排序
func sortColors1(nums []int) {
	var swapColors func(colors []int, target int) int
	swapColors = func(colors []int, target int) int {
		var start int
		for i, c := range colors {
			if c == target {
				colors[i], colors[start] = colors[start], colors[i]
				start++
			}
		}
		return start
	}

	// 先遍历一次，把0放到最左边
	lastZeroIndex := swapColors(nums, 0)

	// 除去0的部分，再遍历一次，把1放到最左边
	swapColors(nums[lastZeroIndex:], 1)
}
