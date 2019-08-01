/**
 * url: https://leetcode-cn.com/problems/4sum/
 * id: 18
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
 * 使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 * 注意：答案中不可以包含重复的四元组。
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 * 满足要求的四元组集合为：
 * [
 *   [-1,  0, 0, 1],
 *   [-2, -1, 1, 2],
 *   [-2,  0, 0, 2]
 * ]
 *
 */

package main

import (
	"sort"
	"strconv"
)

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	var keyMap = make(map[string]int)
	size := len(nums)
	if size >= 4 {
		sort.Ints(nums)
		for i := 1; i < size-2; i++ {
			for j := i + 1; j < size-1; j++ {
				t1 := nums[i] + nums[j]
				left, right := 0, size-1
				for left < i && right > j && left >= 0 && right < size {
					t2 := nums[left] + nums[right]
					if t1+t2 <= target {
						if t1+t2 == target {
							key := strconv.Itoa(nums[left]) + "_" + strconv.Itoa(nums[i]) + "_" + strconv.Itoa(nums[j]) + "_" + strconv.Itoa(nums[right])
							_, ok := keyMap[key]
							if !ok {
								keyMap[key] = 1
								res = append(res, []int{nums[left], nums[i], nums[j], nums[right]})
							}
						}
						left++
					} else {
						right--
					}
				}
			}
		}
	}
	return res
}