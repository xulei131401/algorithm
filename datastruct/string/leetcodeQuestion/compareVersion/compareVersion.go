package main

import (
	"log"
	"strconv"
	"strings"
)

/**
比较版本号
题目：https://leetcode.cn/problems/compare-version-numbers/solutions/?languageTags=golang

给你两个版本号 version1 和 version2 ，请你比较它们。

版本号由一个或多个修订号组成，各修订号由一个 '.' 连接。每个修订号由 多位数字 组成，可能包含 前导零 。每个版本号至少包含一个字符。修订号从左到右编号，下标从 0 开始，最左边的修订号下标为 0 ，下一个修订号下标为 1 ，以此类推。例如，2.5.33 和 0.1 都是有效的版本号。

比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较 忽略任何前导零后的整数值 。也就是说，修订号 1 和修订号 001 相等 。如果版本号没有指定某个下标处的修订号，则该修订号视为 0 。例如，版本 1.0 小于版本 1.1 ，因为它们下标为 0 的修订号相同，而下标为 1 的修订号分别为 0 和 1 ，0 < 1 。

返回规则如下：

如果 version1 > version2 返回 1，
如果 version1 < version2 返回 -1，
除此之外返回 0。


示例 1：

输入：version1 = "1.01", version2 = "1.001"
输出：0
解释：忽略前导零，"01" 和 "001" 都表示相同的整数 "1"
示例 2：

输入：version1 = "1.0", version2 = "1.0.0"
输出：0
解释：version1 没有指定下标为 2 的修订号，即视为 "0"
示例 3：

输入：version1 = "0.1", version2 = "1.1"
输出：-1
解释：version1 中下标为 0 的修订号是 "0"，version2 中下标为 0 的修订号是 "1" 。0 < 1，所以 version1 < version2


提示：

1 <= version1.length, version2.length <= 500
version1 和 version2 仅包含数字和 '.'
version1 和 version2 都是 有效版本号
version1 和 version2 的所有修订号都可以存储在 32 位整数 中

*/

func main() {
	version1 := "1.0"
	version2 := "1.1.0"

	log.Println("比较版本号(字符串分隔):", compareVersion(version1, version2))
	log.Println("比较版本号(双指针):", compareVersion2(version1, version2))
}

// compareVersion O(n+m) O(n+m)
func compareVersion(version1, version2 string) int {
	// 拆分字符串的时候使用了2个切片存放
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < len(v1) || i < len(v2); i++ {
		x, y := 0, 0 // 默认修订号是0
		if i < len(v1) {
			x, _ = strconv.Atoi(v1[i]) // 字符串转数字
		}

		if i < len(v2) {
			y, _ = strconv.Atoi(v2[i])
		}

		// 版本号大
		if x > y {
			return 1
		}

		// 版本号校
		if x < y {
			return -1
		}
		// 相等就继续遍历
	}

	// 相等
	return 0
}

// compareVersion2 O(n+m) O(1)
func compareVersion2(version1, version2 string) int {
	n, m := len(version1), len(version2)
	i, j := 0, 0
	// 只要还能遍历就继续遍历
	// Go字符串数组取索引得到的是rune类型
	for i < n || j < m {
		res1, res2 := 0, 0
		for ; i < n; i++ {
			if version1[i] == '.' {
				i++ // 跳过点号
				break
			}

			// 正序计算数字字符串的数值
			res1 = res1*10 + int(version1[i]-'0')
		}

		for ; j < m; j++ {
			if version2[j] == '.' {
				j++ // 跳过点号
				break
			}

			// 正序遍历就得这么写
			res2 = res2*10 + int(version2[j]-'0')
		}

		if res1 > res2 {
			return 1
		}

		if res1 < res2 {
			return -1
		}
	}

	return 0
}
