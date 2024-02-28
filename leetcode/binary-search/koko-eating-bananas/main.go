package main

import (
	"log"
)

/*
*
https://leetcode.cn/problems/koko-eating-bananas/description/
https://leetcode.com/problems/koko-eating-bananas/description/
*/
func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	log.Println("[手写二分]在排序数组中查找元素的第一个和最后一个位置:", minEatingSpeed(piles, h))
}

func minEatingSpeed(piles []int, h int) int {
	return 0
}
