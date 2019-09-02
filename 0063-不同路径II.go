/**
 * url: https://leetcode-cn.com/problems/unique-paths/submissions/
 */

package main

import (
	"fmt"
	"strconv"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	f := false
	var arr = make(map[string]int)
	for i := 0; i < m; i++ {
		key := strconv.Itoa(i) + "_0"
		if f || obstacleGrid[i][0] == 1 {
			arr[key] = 0
			f = true
			continue
		}
		arr[key] = 1
	}
	f = false
	for j := 0; j < n; j++ {
		key := "0_" + strconv.Itoa(j)
		if f || obstacleGrid[0][j] == 1 {
			arr[key] = 0
			f = true
			continue
		}
		arr[key] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			key := strconv.Itoa(i) + "_" + strconv.Itoa(j)
			if obstacleGrid[i][j] == 1 {
				arr[key] = 0
				continue
			}
			key1 := strconv.Itoa(i-1) + "_" + strconv.Itoa(j)
			key2 := strconv.Itoa(i) + "_" + strconv.Itoa(j-1)
			arr[key] = arr[key1] + arr[key2]
		}
	}
	key := strconv.Itoa(m-1) + "_" + strconv.Itoa(n-1)
	fmt.Println(arr, key)
	return arr[key]
}