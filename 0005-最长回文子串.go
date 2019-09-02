/**
 * url: https://leetcode-cn.com/problems/longest-palindromic-substring/
 * id: 5
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1:
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 * 示例 2:
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s ", longestPalindrome(""))
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	maxLen := 1
	reStr := s[0:1]
	var dp [1000][1000]bool

	// 如果 dp[l, r] = true 那么 dp[l + 1, r - 1] 也一定为 true
	// 关键在这里：[l + 1, r - 1] 一定至少有 2 个元素才有判断的必要
	// 因为如果 [l + 1, r - 1] 只有一个元素，不用判断，一定是回文串
	// 如果 [l + 1, r - 1] 表示的区间为空，不用判断，也一定是回文串
	// [l + 1, r - 1] 一定至少有 2 个元素 等价于 l + 1 < r - 1，即 r - l >  2

	for r := 0; r < len(s); r++ {
		for l := 0; l < r; l++ {
			// 区间应该慢慢放大
			// 状态转移方程：如果头尾字符相等并且中间也是回文
			// 在头尾字符相等的前提下，如果收缩以后不构成区间（最多只有 1 个元素），直接返回 True 即可
			// 否则要继续看收缩以后的区间的回文性
			// 重点理解 or 的短路性质在这里的作用
			// fmt.Printf("%d, %s\n", maxLen, s[l:r+1])
			if s[l] == s[r] && (r - l <= 2 || dp[l + 1][r - 1]) {
				dp[l][r] = true
				curLen := r - l + 1
				if curLen > maxLen {
					maxLen = curLen
					reStr = s[l:r+1]
				}
			}
		}
	}

	return reStr
}

/**
 * 思路： dp[l, r] = (s[l] == s[r] and (l - r >= -2 or dp[l + 1, r - 1]))
 */