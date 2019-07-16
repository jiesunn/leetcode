/**
 * url: https://leetcode-cn.com/problems/3sum/
 * id: 15
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 * 满足要求的三元组集合为：
 * [
 *   [-1, 0, 1],
 *   [-1, -1, 2]
 * ]
 *
 */

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//nums := []int{0,0,0,-13,5,13,12,-2,-11,-1,12,-3,0,-3,-7,-7,-5,-3,-15,-2,14,14,13,6,-11,-11,5,-15,-14,5,-5,-2,0,3,-8,-10,-7,11,-5,-10,-5,-7,-6,2,5,3,2,7,7,3,-10,-2,2,-12,-11,-1,14,10,-9,-15,-8,-7,-9,7,3,-2,5,11,-13,-15,8,-3,-7,-12,7,5,-2,-6,-3,-10,4,2,-5,14,-3,-1,-10,-3,-14,-4,-3,-7,-4,3,8,14,9,-2,10,11,-10,-4,-15,-9,-1,-1,3,4,1,8,1}
	nums := []int{-2,-1,0,1,2,3,-4,4,-5,5}
	arr := threeSum(nums)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]) ; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Printf("\n")
	}
}

func threeSum(nums []int) [][]int {
	var arr [][]int
	m := make(map[string]int)
	size := len(nums)
	if size >= 3 {
		sort.Ints(nums)
		for i := 1; i < size - 1 ; i++ {
			left, right := 0, size - 1
			for left < i && right > i {
				temp := nums[left] + nums[i] + nums[right]
				//fmt.Printf("%d, %d, %d\n", nums[left], nums[i], nums[right])
				if temp <= 0 {
					if temp == 0 {
						key := strconv.Itoa(nums[left]) + "_" + strconv.Itoa(nums[i]) + "_" + strconv.Itoa(nums[right])
						//fmt.Printf("%s\n", key)
						_, ok := m[key]
						if ok == false {
							m[key] = 1
							arr = append(arr, []int{nums[left], nums[i], nums[right]})
						}
					}
					left++
				} else {
					right--
				}
			}
		}
	}
	return arr
}

/**
 * 思路: 排序，以0为界限，固定两端各一个值，根据不同的条件遍历其间的元素
 */