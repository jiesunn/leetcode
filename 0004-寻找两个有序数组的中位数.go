/**
 * url: https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
 * id: 4
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 * 则中位数是 2.0
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1,2}
	nums2 := []int{3,4}
	fmt.Printf("%f", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 保证数组1一定最短
	arr1 := nums1
	arr2 := nums2
	if len(nums1) > len(nums2) {
		arr1 = nums2
		arr2 = nums1
	}

	// Ci 为第i个数组的割,比如C1为2时表示第1个数组只有2个元素。LMaxi为第i个数组割后的左元素。RMini为第i个数组割后的右元素。
	// 我们目前是虚拟加了'#'所以数组1是2*n长度
	n, m := len(arr1), len(arr2)
	LMax1, LMax2, RMin1, RMin2, c1, c2, lo, hi := 0, 0, 0, 0, 0, 0, 0, 2 * n

	// 二分查找满足条件的C1
	for lo <= hi {
		c1 = (lo + hi) / 2
		c2 = m + n - c1

		// 边界条件的取值
		if c1 == 0 {
			LMax1 = math.MinInt64
		} else {
			LMax1 = arr1[(c1 - 1) / 2]
		}
		if c1 == 2 * n {
			RMin1 = math.MaxInt64
		} else {
			RMin1 = arr1[c1 / 2]
		}
		if c2 == 0 {
			LMax2 = math.MinInt64
		} else {
			LMax2 = arr2[(c2 - 1) / 2]
		}
		if c2 == 2 * m {
			RMin2 = math.MaxInt64
		} else {
			RMin2 = arr2[c2 / 2]
		}

		// 移位
		if LMax1 > RMin2 {
			hi = c1 - 1
		} else if LMax2 > RMin1 {
			lo = c1 + 1
		} else {
			break
		}
	}

	return (math.Max(float64(LMax1), float64(LMax2)) + math.Min(float64(RMin1), float64(RMin2))) / 2.0
}

/**
思路： TOP K
参考： https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/4-xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-shu
*/