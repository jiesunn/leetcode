/**
 * url: https://leetcode-cn.com/problems/palindrome-linked-list/
 * id: 234
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 * 输入: 1->2->2->1
 * 输出: true
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	p := ListNode{1, nil}
	q := ListNode{2, &p}
	head := ListNode{1, &q}
	if isPalindrome(&head) {
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	//方法1
	//newHead := copyList(head)

	//方法2
	newHead := middle(head)

	newHead = reverseList(newHead)

	p, q := head, newHead
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func copyList(head *ListNode) *ListNode {
	var copyHead *ListNode
	var q *ListNode
	p := head
	for p != nil {
		node := ListNode{p.Val, nil}
		if copyHead == nil {
			copyHead = &node
			q = copyHead
		} else {
			q.Next = &node
			q = q.Next
		}
		p = p.Next
	}
	return copyHead
}

func reverseList(head *ListNode) *ListNode {
	p, q := head, head.Next
	for q != nil {
		if p == head {
			p.Next = nil
		}
		temp := q.Next
		q.Next = p
		p = q
		q = temp
	}
	return p
}

func middle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil  {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
	}
	return slow
}

/**
 * 思路： 反转链表
 * 优化： 快慢指针找到链表中点，反转链表后半段
 */