/**
 * url: https://leetcode-cn.com/problems/minimum-path-sum/
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := minPathSum([][]int{{1,2}, {1,1}})
	fmt.Println(v)
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var arr = make(map[string]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			key := strconv.Itoa(i) + "_" + strconv.Itoa(j)
			key1 := strconv.Itoa(i-1) + "_" + strconv.Itoa(j)
			key2 := strconv.Itoa(i) + "_" + strconv.Itoa(j-1)
			_, ok1 := arr[key1]
			_, ok2 := arr[key2]
			if i == 0 && j == 0 {
				arr[key] = grid[i][j]
				continue
			}
			if i == 0 || (ok2 && arr[key1] >= arr[key2]) {
				arr[key] = arr[key2] + grid[i][j]
				continue
			}
			if j == 0 || (ok1 && arr[key1] < arr[key2]) {
				arr[key] = arr[key1] + grid[i][j]
				continue
			}
		}
	}
	fmt.Println(arr)
	key := strconv.Itoa(m-1) + "_" + strconv.Itoa(n-1)
	return arr[key]
}