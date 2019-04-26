/**
 * url: https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
 * id: 240
 *
 * 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
 * 每行的元素从左到右升序排列。
 * 每列的元素从上到下升序排列。
 *
 * 示例:
 *
 * 现有矩阵 matrix 如下：
 *
 * [
 *   [1,   4,  7, 11, 15],
 *   [2,   5,  8, 12, 19],
 *   [3,   6,  9, 16, 22],
 *   [10, 13, 14, 17, 24],
 *   [18, 21, 23, 26, 30]
 * ]
 *
 * 给定 target = 5，返回 true。
 *
 * 给定 target = 20，返回 false。
 */

package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if searchMatrix(matrix, 0) {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) != 0 {
		row := 0
		col := len(matrix[0]) - 1
		for {
			if row >= len(matrix) || col < 0 {
				break
			}
			if matrix[row][col] == target {
				return true
			}
			if matrix[row][col] < target {
				row++
			} else {
				col--
			}
		}
	}

	return false
}

/**
 * 思路：若matrix[R][C]>target, 则matrix[r][c](r>=R&&c>=C)>target, 反之亦然
 *
 * 执行用时 : 40 ms, 在Search a 2D Matrix II的Go提交中击败了96.02% 的用户
 * 内存消耗 : 6.3 MB, 在Search a 2D Matrix II的Go提交中击败了46.03% 的用户
 */
