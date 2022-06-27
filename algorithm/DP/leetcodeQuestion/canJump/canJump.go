package main

import "log"

/**
题目：https://leetcode-cn.com/problems/jump-game/

	跳跃游戏
	给定一个非负整数数组nums ，你最初位于数组的 第一个下标 。

	数组中的每个元素代表你在该位置可以跳跃的最大长度。

	判断你是否能够到达最后一个下标。


	示例1：

	输入：nums = [2,3,1,1,4]
	输出：true
	解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
	示例2：

	输入：nums = [3,2,1,0,4]
	输出：false
	解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


	提示：

	1 <= nums.length <= 3 * 104
	0 <= nums[i] <= 105


*/

func main() {
	nums := []int{3, 2, 1, 0, 4}
	log.Println("跳跃游戏-贪心:", canJump(nums))
	log.Println("跳跃游戏-动态规划:", canJUmp2(nums))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	var res int
	for i := 0; i < len(nums); i++ {
		// res 为右边界，计算最远之内还能不能到大最后
		if i <= res {
			res = max(res, i+nums[i])
			if res >= len(nums)-1 {
				return true
			}
		} else {
			return false
		}
	}

	return false
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func canJUmp2(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(nums)-1]
}
