/**
 * url: https://leetcode-cn.com/problems/remove-invalid-parentheses/
 * id: 301
 *
 * 删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
 * 说明: 输入可能包含了除 ( 和 ) 以外的字符。
 *
 * 示例 1:
 *
 * 输入: "()())()"
 * 输出: ["()()()", "(())()"]
 *
 * 示例 2:
 *
 * 输入: "(a)())()"
 * 输出: ["(a)()()", "(a())()"]
 *
 * 示例 3:
 *
 * 输入: ")("
 * 输出: [""]
 *
 */

package main

import (
	"fmt"
)

func main() {
	for _,value := range removeInvalidParentheses(")(f") {
		fmt.Printf("'%s'\n", value)
	}
}

func removeInvalidParentheses(s string) []string {
	var arr []string

	//记录要删掉的左括号或右括号数量
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "(" {
			left++
		}
		if s[i:i+1] == ")" {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}

	return dfs(arr, s, 0, left, right)
}

/**
 * arr []string 结果集
 * s string 	当前字符串
 * index int 	要删除字符的下标
 * left int 	要删除的左括号的数目
 * right int 	要删除的右括号类型
 */
func dfs(arr []string, s string, index int, left int, right int) []string {
	if left == 0 && right == 0 {
		if check(s) {
			arr = append(arr, s)
		}
		return arr
	}
	for i:= index; i < len(s); i++ {
		if i - 1 >= index && s[i] == s[i - 1] {
			continue
		}
		if left > 0 && s[i:i+1] == "(" {
			arr = dfs(arr, s[:i] + s[i+1:], i, left - 1, right)
		}
		if right > 0 && s[i:i+1] == ")" {
			arr = dfs(arr, s[:i] + s[i+1:], i, left, right - 1)
		}
	}
	return arr
}

// 检验合法性
func check(s string) bool {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "(" {
			cnt++
		}
		if s[i:i+1] == ")" {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}

/**
 思路： 找出需要删除的括号数，然后dfs
 */
