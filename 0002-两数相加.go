/**
 * url: https://leetcode-cn.com/problems/two-sum/
 * id: 1
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 *
 * 示例:
 *
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 */

package main

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num, flag int
	var head, point *ListNode

	for l1 != nil || l2 != nil || flag != 0  {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		num = n1 + n2 + flag

		var node ListNode
		if num < 10 {
			node.Val = num
			flag = 0
		} else {
			node.Val = num % 10
			flag = 1
		}

		if head == nil {
			head = &node
			point = head
		} else {
			point.Next = &node
			point = point.Next
		}
	}

	return head
}

/**
 * 思路： 进位法
 */