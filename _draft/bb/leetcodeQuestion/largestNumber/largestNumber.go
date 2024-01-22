package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

/*
*
最大数
题目：https://leetcode.cn/problems/largest-number/solutions/?orderBy=hot&languageTags=golang

给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 109
*/
func main() {
	nums := []int{3, 30, 34, 5, 9}
	log.Println("最大数:", largestNumber(nums))
}

// largestNumber 排序法 O(nlognlogm) O(logn)
func largestNumber(nums []int) string {
	// 最直观的做法就是
	// 先把数组元素都处理成字符串
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}

	// 排序,相邻两个值组成的字符串看哪个比较大然后交换位置
	sort.Slice(ss, func(i, j int) bool {
		//log.Println(ss, i, j, "ss[i]=", ss[i], ";ss[j]=", ss[j])
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})

	//log.Println(ss)
	// 拼接
	str := strings.Join(ss, "")
	// 处理首个字符是0的特殊情况
	if str[0] == '0' {
		return "0"
	}

	return str
}
