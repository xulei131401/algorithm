package main

import "fmt"

/*
题目：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

删除有序数组中的重复项

给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
大前提是已经排序过了：

给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么nums的前 k 个元素应该保存最终结果。

将最终结果插入nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

*/

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//l := removeDuplicates(nums)
	l := removeDuplicates1(nums)
	fmt.Println("l,nums:", l, nums)
}

func removeDuplicates(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1
	preNum := nums[0]
	for _, v := range nums {
		if v != preNum {
			preNum = v
			nums[count] = v
			count++
		}
	}
	return count
}

func removeDuplicates2(nums []int) int {
	nui := 0
	for i := 1; i < len(nums); i++ {
		if nums[nui] != nums[i] {
			nui++
			nums[nui] = nums[i]
		}
	}
	return nui + 1
}
