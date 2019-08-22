package main

import (
    "fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	var nodeMap = make(map[int]int)
	nodeMap = test(root, 1, nodeMap)
	fmt.Println(nodeMap)
	maxkey := 1
	for key, value := range nodeMap {
		if value >= nodeMap[maxkey] {
			maxkey = key
		}
	}
	return maxkey
}

func test(root *TreeNode, floor int, nodeMap map[int]int) map[int]int {
	if root == nil {
		return nodeMap
	}

	_, ok := nodeMap[floor]
	fmt.Println(floor, ok)
	if ok {
		nodeMap[floor] += root.Val
	} else {
		nodeMap[floor] = root.Val
	}
	nodeMap = test(root.Left, floor+1, nodeMap)
	nodeMap = test(root.Right, floor+1, nodeMap)
	return nodeMap
}