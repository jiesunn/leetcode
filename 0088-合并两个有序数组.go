/**
 * url: https://leetcode-cn.com/problems/merge-sorted-array/
 * id: 88
 *
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 *
 * 说明:
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 * 示例:
 *
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * 输出: [1,2,2,3,5,6]
 */

package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	i := m - 1
	j := n - 1
	for {
		if i < 0 || j < 0 {
			break
		}
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--
	}
	for {
		if j < 0 {
			break
		}
		nums1[p] = nums2[j]
		j--
		p--
	}
}

/**
 * 思路：从nums1的空位开始，从后往前依次填入未访问过的最大值
 *
 * 执行用时 : 4 ms, 在Merge Sorted Array的Go提交中击败了98.95% 的用户
 * 内存消耗 : 3.7 MB, 在Merge Sorted Array的Go提交中击败了45.45% 的用户
 */
