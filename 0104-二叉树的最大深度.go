/**
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
 * id: 8
 *
 * 给定一个二叉树，找出其最大深度。
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 返回它的最大深度 3 。
 *
 */

package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Printf("%d", myAtoi("+ 002147483646"))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	return dfs(root, 0, 0)
}

func dfs(root *TreeNode, curDeep int, maxDeep int) int {
	if root == nil {
		return maxDeep
	}

	curDeep++

	if maxDeep < curDeep {
		maxDeep = curDeep
	}
	//fmt.Printf("%d, %d\n", curDeep, maxDeep)
	maxDeep = dfs(root.Left, curDeep, maxDeep)

	if maxDeep < curDeep {
		maxDeep = curDeep
	}
	//fmt.Printf("%d, %d\n", curDeep, maxDeep)
	maxDeep = dfs(root.Right, curDeep, maxDeep)

	return maxDeep
}