/**
 * url: https://leetcode-cn.com/problems/unique-paths/submissions/
 */

package main

import (
	"strconv"
)

func uniquePaths(m int, n int) int {
	var arr = make(map[string]int)
	for i := 0; i < m; i++ {
		key := strconv.Itoa(i) + "_0"
		arr[key] = 1
	}
	for j := 0; j < n; j++ {
		key := "0_" + strconv.Itoa(j)
		arr[key] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			key := strconv.Itoa(i) + "_" + strconv.Itoa(j)
			key1 := strconv.Itoa(i-1) + "_" + strconv.Itoa(j)
			key2 := strconv.Itoa(i) + "_" + strconv.Itoa(j-1)
			arr[key] = arr[key1] + arr[key2]
		}
	}
	key := strconv.Itoa(m) + "_" + strconv.Itoa(n)
	return arr[key]
}