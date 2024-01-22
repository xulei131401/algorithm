package main

import "log"

/**
题目：https://leetcode.cn/problems/word-search/
单词搜索

给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。


示例 1：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true
示例 3：


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false

提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成

进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

*/
func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"
	log.Println("单词搜索-回溯:", exist(board, word))
}

// exist 回溯 O(MN 3^L) O(MN)
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m) // 用来标记回溯的过程中元素是否已使用过
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var backtracking func(x, y, i int) bool // i 表示匹配word的位置
	backtracking = func(x, y, i int) bool {
		if i == len(word) {
			return true
		}

		// 边界条件
		if x < 0 || x >= m || y < 0 || y >= n {
			return false
		}

		// 使用过或者不匹配当前字符直接返回false
		if used[x][y] || board[x][y] != word[i] {
			return false
		}

		used[x][y] = true
		// 下，上，右，左
		res := backtracking(x+1, y, i+1) || backtracking(x-1, y, i+1) ||
			backtracking(x, y+1, i+1) || backtracking(x, y-1, i+1)

		if res {
			return true
		} else {
			// 发现走不通的话，重置
			used[x][y] = false
			return false
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 遍历矩阵，寻找第一个匹配的字符，然后以此作为开始进行回溯
			if board[i][j] == word[0] && backtracking(i, j, 0) {
				return true
			}
		}
	}

	return false
}
