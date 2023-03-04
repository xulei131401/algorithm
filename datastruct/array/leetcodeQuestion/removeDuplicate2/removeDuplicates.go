package main

import (
	"log"
)

/*
题目：https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/

给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。


说明：
为什么返回数值是整数，但输出的答案是数组呢？
请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
你可以想象内部操作如下:
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}


示例 1：
输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 不需要考虑数组中超出新长度后面的元素。

示例 2：
输入：nums = [0,0,1,1,1,1,2,3,3]
输出：7, nums = [0,0,1,1,2,3,3]
解释：函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。 不需要考虑数组中超出新长度后面的元素。


提示：
1.1 <= nums.length <= 3 * 104
2.-104 <= nums[i] <= 104
3.nums 已按升序排列


总结：
1.数组升序
2.只需要返回长度即可
3.原地修改数组的意思是，保证长度以内的元素正确即可，长度之外的元素忽略，也就是说删除的时候直接覆盖元素值即可。
不能交换，只能覆盖

*/

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 1, 2, 2}
	//l := removeDuplicates(nums)
	//l := removeDuplicates2(nums)
	l := removeDuplicates3(nums)
	log.Println("l,nums:", l, nums[:l])
}

// removeDuplicates 双指针法，O(n) O(1)
func removeDuplicates(nums []int) int {
	n := len(nums)
	// 长度小于等于2没有比较的意义了，直接返回
	if n <= 2 {
		return n
	}

	// 如何理解left-2,newTailIndex,right这三个索引的含义?
	// newTailIndex代表符合题意的数组长度
	// newTailIndex-2代表相同元素区间的左边界
	// right代表实际遍历过程中相同元素区间的右边界
	// 当左右边界的值不同的时候，意味着要保留当前遍历的值，把这个值覆盖到要保留的数组长度位置即可
	newTailIndex, right := 2, 2
	for right < n {
		if nums[newTailIndex-2] != nums[right] {
			nums[newTailIndex] = nums[right]
			newTailIndex++
		}
		right++
	}
	return newTailIndex
}

// removeDuplicates2 双指针法通用写法，O(n) O(1)
func removeDuplicates2(nums []int) int {
	sameLengthLimit := 2 // 最大相同元素个数
	n := len(nums)
	if n <= sameLengthLimit {
		return n
	}

	f := func(limit int) int {
		newTailIndex, right := limit, limit
		for right < n {
			// 当前值要保留
			if nums[newTailIndex-limit] != nums[right] {
				nums[newTailIndex] = nums[right]
				newTailIndex++
			}
			// 不确定当前值是否要保留的话就一直向后遍历
			right++
		}

		return newTailIndex
	}

	return f(sameLengthLimit)
}

// removeDuplicates3 直接遍历通用写法，O(n) O(1)
func removeDuplicates3(nums []int) int {
	f := func(k int) int {
		newLength := 0
		for _, v := range nums {
			// 保留当前值的逻辑
			// 1.根据题意，前k个数可以直接保留
			// 2.当前值与前k个值不同的时候保留
			if newLength < k || nums[newLength-k] != v {
				nums[newLength] = v
				newLength++
			}
		}
		return newLength
	}
	return f(2)
}
