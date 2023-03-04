package main

import "log"

/*
*
https://leetcode.cn/problems/set-matrix-zeroes/description/
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
*/
func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	log.Println("矩阵置零:", matrix)
}

// setZeroes 数组标记法 O(mn) O(m+n)
// 从前向后遍历，不会有遗漏，使用2个数组分别标记行和列，然后重新遍历置零
func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}

// setZeroes1 变量标记法 O(mn) O(1)
func setZeroes1(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row0, col0 := false, false
	// 遍历第一行（其实是遍历第一行的每一列）
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}

	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}

	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}

// setZeroes2 变量标记法 O(mn) O(1)
func setZeroes2(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		
		if col0 {
			matrix[i][0] = 0
		}
	}
}
