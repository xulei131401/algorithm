package main

import (
	"log"
)

/*
*
折半插入排序
*/
func main() {
	list := []int{4, 3, 5, 1, 2, 9}
	log.Println("折半插入排序:", InsertSort(list))
}

// InsertSort 折半插入排序 最好时间复杂度O(n),最坏时间复杂度O(n^2),空间复杂度O(1)
func InsertSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		// 0~i区间的元素有序，所以可以使用二分法减少操作次数
		// 利用二分查找的特性缩小区间，最终i只需要把low~i之前的数字移动一次即可
		low, high := 0, i
		for low <= high {
			middle := low + (high-low)>>1
			if nums[middle] > nums[i] {
				high = middle - 1
			} else {
				low = middle + 1
			}
		}

		// 想把i位置元素移动到low位置，倒着两两交换
		for j := i - 1; j >= low; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}

	return nums
}
