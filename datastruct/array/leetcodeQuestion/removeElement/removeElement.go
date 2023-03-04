package main

import "fmt"

/*
*
https://leetcode-cn.com/problems/remove-element/
移除元素
给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

注意：原地删除的意思就会把要删除的值移动到数组的末尾，然后返回新的长度，超出长度的部分当做是不存在
*/
func main() {
	arr := []int{3, 2, 2, 1}
	arr1 := []int{3, 2, 2, 1}
	val := 2
	fmt.Println("原地删除后长度-双指针1：", removeElement(arr, val), arr)
	fmt.Println("原地删除后长度-双指针2：", removeElement1(arr1, val), arr1)
}

// removeElement1 双指针 O(n) O(1)
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	for left <= right {
		// 左边界始终是要删除的元素
		for nums[left] != val && left < right {
			left++
		}

		// 右边界始终是不删除的元素
		for nums[right] == val && right > left {
			right--
		}

		// 当两端都移动到同一个元素位置时，判断是否等于要删除的元素，不是则停止运行
		if left == right && nums[left] != val {
			break
		}

		// 交换当前元素，并各自向前移动，此时right位置的元素应该保留，但是因为right-1了，所以最终长度是right+1
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}

	// 之所以不用left，是因为最终left可能等于right，也可能大于right，等于的情况下用left-1就不对了，所以用right+1更直接
	return right + 1
}

// removeElement1 双指针 O(n) O(1)
func removeElement1(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	// left代表第一个要被删除的元素位置
	// right代表第一个要保留元素的位置
	// right的值覆盖left的值，因为不需要老驴元素顺序
	left, right, length := 0, 0, len(nums)
	for right < length {
		if nums[right] != val {
			// 此时要保留当前位置的值
			nums[left] = nums[right]
			left++
		}

		// 一直向前移动
		right++
	}

	return left
}
