/**
 * url: https://leetcode-cn.com/problems/single-number/
 * id: 136
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 *
 * 说明：
 *
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 * 示例 1:
 * 输入: [2,2,1]
 * 输出: 1
 *
 * 示例 2:
 * 输入: [4,1,2,1,2]
 * 输出: 4
 */

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1}
	fmt.Printf("%d\n", singleNumber(nums))
}

func singleNumber(nums []int) int {
	//方法1
	num := 0
	len := len(nums)
	for i := 0; i < len; i++ {
		num = num ^ nums[i]
	}

	//方法2
	num = 0
	for _, value := range nums {
		num = num ^ value
	}

	return num
}

/**
 * 思路：位运算之异或（相同数运算结果为0, 一个数与0运算结果为这个数本身）
 *
 * 执行用时 : 24 ms, 在Single Number的Go提交中击败了54.76% 的用户
 * 内存消耗 : 4.9 MB, 在Single Number的Go提交中击败了49.81% 的用户
 */
