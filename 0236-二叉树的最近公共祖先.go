package main

var m = make(map[string]*TreeNode)
var pKey, qKey string

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	dfs(root, p, q, "")

	curKey := "";
	for i := 0; i< len(pKey); i++ {
		if pKey[i] == qKey[i] {
			curKey += string(pKey[i])
		}
	}
	return m[curKey]
}

func dfs(root, p, q *TreeNode, key string) {
	if root == nil {
		return
	}
	
	m[key] = root
	if root == p {
		pKey = key
	}
	if root == q {
		qKey = key
	}

	dfs(root.Left, p, q, key + "1")
	dfs(root.Right, p, q, key + "2")
}