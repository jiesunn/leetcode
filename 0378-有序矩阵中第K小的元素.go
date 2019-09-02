/**
 * url: https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/submissions/
 * id: 378
 *
 * 给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第k小的元素。
 * 请注意，它是排序后的第k小元素，而不是第k个元素。
 *
 * 示例：
 * matrix = [
 *    [ 1,  5,  9],
 *    [10, 11, 13],
 *    [12, 13, 15]
 * ],
 * k = 8,
 *
 * 返回 13。
 *
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := [][]int{{1, 3, 5}, {6, 7, 12}, {11, 14, 14}}
	//arr := [][]int{{1, 2}, {1, 3}}
	arr := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	k := 5
	v := kthSmallest(arr, k)
	fmt.Println(v)
}

func kthSmallest(matrix [][]int, k int) int {
	size := len(matrix)
	if size < 1 {
		return 0
	}

	tarRow := k / size
	tarCol := k%size - 1
	if k%size == 0 {
		tarRow = k/size - 1
		tarCol = size - 1
	}
	fmt.Println(tarRow, tarCol)

	var arr []int
	row, col := 0, 0

	// 目标及其右上部
	i, j := tarRow, tarCol
	for i >= 0 && j < size {
		fmt.Println(i, j, matrix[i][j])
		arr = append(arr, matrix[i][j])
		row = i
		col = j
		i--
		j++
	}

	// 目标的左下部
	i, j = tarRow+1, tarCol-1
	for i < size && j >= 0 {
		fmt.Println(i, j, matrix[i][j])
		arr = append(arr, matrix[i][j])
		i++
		j--
	}

	sort.Ints(arr)
	i, j = row, col
	index := 0
	for i < size && j >= 0 {
		matrix[i][j] = arr[index]
		fmt.Println(i, j, arr[index])
		if i == tarRow && j == tarCol {
			if tarCol == size-1 && tarRow < size-1 {
				if matrix[i][j] > matrix[i+1][0] {
					return matrix[i+1][0]
				}
			}
			if tarCol == 0 && tarRow > 0 {
				if matrix[i][j] < matrix[i-1][size-1] {
					return matrix[i-1][size-1]
				}
			}
			return arr[index]
		}
		index++
		i++
		j--
	}

	return 0
}
