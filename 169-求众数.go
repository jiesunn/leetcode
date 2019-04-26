/**
 * url: https://leetcode-cn.com/problems/majority-element/description/
 * id: 169
 *
 * 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在众数。
 *
 * 示例 1:
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 */

package main

import (
	"fmt"
)

func main() {
	nums := [3]int{3, 2, 3}
	fmt.Printf("%d\n", majorityElement(nums[:]))
}

func majorityElement(nums []int) int {
	num := 0
	count := 0
	len := len(nums)
	for i := 0; i < len; i++ {
		if count == 0 {
			num = nums[i]
		}
		if num == nums[i] {
			count++
		} else {
			count--
		}
	}

	return num
}

/**
 * 思路：不同元素相消，最后剩下的一定是众数
 *
 * 执行用时 : 44 ms, 在Majority Element的Go提交中击败了45.43% 的用户
 * 内存消耗 : 6 MB, 在Majority Element的Go提交中击败了32.90% 的用户
 */
