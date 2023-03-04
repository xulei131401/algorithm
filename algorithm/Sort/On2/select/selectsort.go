package main

import (
	"log"
)

/*
*
选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理是：
第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，然后再从剩余的未排序元素中寻找到最小（大）元素，
然后放到已排序的序列的末尾。以此类推，直到全部待排序的数据元素的个数为零。选择排序是不稳定的排序方法

时间复杂度：O(n^2)
空间复杂度：O(1)
不稳定

步骤：
	1.最外层for表示趟数
	2.内层for用来获取剩余区间内的最小值交换到对应位置
	3.一句话总结：每一趟都选取一个最大值或者最小值到前边，N趟以后自然有序

*/
func main() {
	nums := []int{3, 8, 4, 7, 6, 5, 1, 2}
	log.Println("选择排序:", selectSort(nums))
}

func selectSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// 假设i位置存放剩余i+1~length-1区间的最小值，依次遍历
	for i := 0; i < length-1; i++ {
		// minIdx 临时变量保存上一个最小值
		minIdx := i
		for j := i + 1; j <= length-1; j++ {
			// 找最小数或者最大数并保存索引
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		// 交换元素，此处表现出不稳定性，因为是跳跃性的交换元素
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}
