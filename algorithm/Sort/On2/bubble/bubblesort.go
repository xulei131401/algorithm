package main

import (
	"log"
)

/*
*
冒泡排序：冒泡排序（Bubble Sort），是一种计算机科学领域的较简单的排序算法。
它重复地走访过要排序的元素列，依次比较两个相邻的元素，如果顺序（如从大到小、首字母从Z到A）错误就把他们交换过来。走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素列已经排序完成。
这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端（升序或降序排列），就如同碳酸饮料中二氧化碳的气泡最终会上浮到顶端一样，故名“冒泡排序”。

时间复杂度：O(N)~O(N2)
稳定性排序

步骤：
	1.最外层for表示趟数
	2.第二层for从第一个数开始一次比较，交换顺序
	3.优化点1：这一趟若没有交换元素的动作，认为已经有序，无需进行后续遍历；优化点2：第二趟的比较次数可以直接从0~最后一次交换的元素区间内遍历即可
*/
func main() {
	nums := []int{11, 8, 2, 5, 7, 10, 3, 6}
	nums1 := []int{11, 8, 2, 5, 7, 10, 3, 6}
	nums2 := []int{11, 8, 2, 5, 7, 10, 3, 6}
	log.Println("冒泡排序:", bubbleSort(nums))
	log.Println("冒泡排序-优化1:", bubbleSort2(nums1))
	log.Println("冒泡排序-优化2:", bubbleSort3(nums2))
}

func bubbleSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// start表示几趟
	// 每一趟排序都能确定一个数字的最终位置
	start := 1
	for start <= length {
		// j=0表示从第一个数开始处理
		// j必须小于length-1，并且随着start增大而减小
		for j := 0; j < length-start; j++ {
			// 比较j,j+1大小
			if nums[j+1] < nums[j] {
				// 稳定，因为只会交换相邻元素
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

		start++
	}

	return nums
}

// bubbleSort2 优化点：若当前趟没有发生数据交换，认为已全部有序
func bubbleSort2(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// 发现某一趟排序之后没有发生过交换元素的行为，认为已经有序，提前中断
	isSwap := true
	start := 1
	for start <= length {
		if !isSwap {
			break
		}

		for j := 0; j < length-start; j++ {
			// 比较j,j+1大小
			if nums[j+1] < nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSwap = true
			}
		}

		start++
	}

	return nums
}

// bubbleSort3 优化点：每一趟只需要处理0~最后交换区间内的元素即可
// 也就是说每一趟需要交换的次数可以不同
func bubbleSort3(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	// 更优化的写法：
	isSwap := true
	// 最后一个没有经过交换的元素的下标
	indexOfLastUnsortedElement := length - 1
	// 上次发生交换的位置
	swappedIndex := -1

	for isSwap {
		// 本轮无交换表示有序，直接退出
		isSwap = false
		// 更加优化的写法是：每次经过一趟排序以后，indexOfLastUnsortedElement~length-1位置的元素是有序的，只需要遍历0~indexOfLastUnsortedElement的元素交换即可
		for j := 0; j < indexOfLastUnsortedElement; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				// 说明本趟排序发生了交换，则还需要下一趟排序
				isSwap = true
				// 标记当前发生交换行为的元素index
				swappedIndex = j
			}
		}

		// 设置最后一个排序交换的index
		indexOfLastUnsortedElement = swappedIndex
	}

	return nums
}
